use anyhow::{ensure, Context};
use tokio::sync::Notify;
use tokio::try_join;

mod common;
use common::{init, with_nats};

#[tokio::test(flavor = "multi_thread")]
async fn rust() -> anyhow::Result<()> {
    init().await;

    with_nats(|_, nats_client| async {
        let client = wrpc_transport_nats::Client::new(nats_client, "test-prefix".to_string());

        let shutdown = Notify::new();
        let started = Notify::new();

        try_join!(
            async {
                mod bindings {
                    wit_bindgen_wrpc::generate!("server");
                }

                #[derive(Clone)]
                struct Handler;
                use bindings::exports::wrpc::keyvalue;
                type Result<T, E = keyvalue::store::Error> = core::result::Result<T, E>;

                impl<Ctx: Send> keyvalue::store::Handler<Ctx> for Handler {
                    async fn delete(
                        &self,
                        _cx: Ctx,
                        bucket: String,
                        key: String,
                    ) -> anyhow::Result<Result<()>> {
                        assert_eq!(bucket, "bucket");
                        assert_eq!(key, "key");
                        Ok(Ok(()))
                    }

                    async fn exists(
                        &self,
                        _cx: Ctx,
                        bucket: String,
                        key: String,
                    ) -> anyhow::Result<Result<bool>> {
                        assert_eq!(bucket, "bucket");
                        assert_eq!(key, "key");
                        Ok(Ok(true))
                    }

                    async fn get(
                        &self,
                        _cx: Ctx,
                        bucket: String,
                        key: String,
                    ) -> anyhow::Result<Result<Option<Vec<u8>>>> {
                        assert_eq!(bucket, "bucket");
                        assert_eq!(key, "key");
                        Ok(Ok(Some(vec![0x42, 0xff])))
                    }

                    async fn set(
                        &self,
                        _cx: Ctx,
                        bucket: String,
                        key: String,
                        value: Vec<u8>,
                    ) -> anyhow::Result<Result<()>> {
                        assert_eq!(bucket, "bucket");
                        assert_eq!(key, "key");
                        assert_eq!(value, b"test");
                        Ok(Ok(()))
                    }

                    async fn list_keys(
                        &self,
                        _cx: Ctx,
                        bucket: String,
                        cursor: Option<u64>,
                    ) -> anyhow::Result<Result<keyvalue::store::KeyResponse>> {
                        assert_eq!(bucket, "bucket");
                        assert_eq!(cursor, Some(42));
                        Ok(Ok(keyvalue::store::KeyResponse {
                            cursor: None,
                            keys: vec!["key".to_string()],
                        }))
                    }
                }

                impl<Ctx: Send> keyvalue::atomics::Handler<Ctx> for Handler {
                    async fn increment(
                        &self,
                        _cx: Ctx,
                        bucket: String,
                        key: String,
                        delta: u64,
                    ) -> anyhow::Result<Result<u64, keyvalue::store::Error>> {
                        assert_eq!(bucket, "bucket");
                        assert_eq!(key, "key");
                        assert_eq!(delta, 42);
                        Ok(Ok(4242))
                    }
                }

                let fut = bindings::serve(&client, Handler, shutdown.notified());

                started.notify_one();

                fut.await.context("failed to serve world")
            },
            async {
                mod bindings {
                    wit_bindgen_wrpc::generate!({
                        world: "client",
                        additional_derives: [Eq, PartialEq],
                    });
                }
                use bindings::wrpc::keyvalue;

                started.notified().await;

                try_join!(
                    async {
                        keyvalue::store::delete(&client, "bucket", "key")
                            .await
                            .context("failed to call `delete`")?
                            .context("`delete` call failed")
                    },
                    async {
                        let v = keyvalue::store::exists(&client, "bucket", "key")
                            .await
                            .context("failed to call `exists`")?
                            .context("`exists` call failed")?;
                        ensure!(v, "`exists` should have returned `true`");
                        Ok(())
                    },
                    async {
                        let v = keyvalue::store::get(&client, "bucket", "key")
                            .await
                            .context("failed to call `get`")?
                            .context("`get` call failed")?;
                        ensure!(
                            v.as_deref() == Some(&[0x42, 0xff]),
                            "`get` should have returned `Some([0x42, 0xff])`, got `{v:?}`"
                        );
                        Ok(())
                    },
                    async {
                        keyvalue::store::set(&client, "bucket", "key", b"test")
                            .await
                            .context("failed to call `set`")?
                            .context("`set` call failed")
                    },
                    async {
                        let v = keyvalue::store::list_keys(&client, "bucket", Some(42))
                            .await
                            .context("failed to call `list-keys`")?
                            .context("`list-keys` call failed")?;
                        ensure!(
                            v == keyvalue::store::KeyResponse {
                                cursor: None,
                                keys: vec!["key".to_string()]
                            },
                            r#"`list-keys` should have returned `{{None, ["key"]}}`, got `{v:?}`"#
                        );
                        Ok(())
                    },
                )?;
                shutdown.notify_one();
                Ok(())
            }
        )?;
        Ok(())
    })
    .await
}
