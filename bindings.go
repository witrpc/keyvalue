//go:generate wit-bindgen-wrpc go --world imports --out-dir imports --package github.com/wrpc/keyvalue/imports wit
//go:generate wit-bindgen-wrpc go --world client --out-dir client --package github.com/wrpc/keyvalue/client wit
//go:generate wit-bindgen-wrpc go --world server --out-dir server --package github.com/wrpc/keyvalue/server wit

package keyvalue
