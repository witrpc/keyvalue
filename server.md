<h1><a name="server">World server</a></h1>
<ul>
<li>Exports:
<ul>
<li>interface <a href="#wrpc_keyvalue_store_0_2_0_draft"><code>wrpc:keyvalue/store@0.2.0-draft</code></a></li>
<li>interface <a href="#wrpc_keyvalue_atomics_0_2_0_draft"><code>wrpc:keyvalue/atomics@0.2.0-draft</code></a></li>
</ul>
</li>
</ul>
<h2><a name="wrpc_keyvalue_store_0_2_0_draft"></a>Export interface wrpc:keyvalue/store@0.2.0-draft</h2>
<hr />
<h3>Types</h3>
<h4><a name="error"></a><code>variant error</code></h4>
<p>The set of errors which may be raised by functions in this package</p>
<h5>Variant Cases</h5>
<ul>
<li>
<p><a name="error.no_such_store"></a><code>no-such-store</code></p>
<p>The host does not recognize the store identifier requested.
</li>
<li>
<p><a name="error.access_denied"></a><code>access-denied</code></p>
<p>The requesting component does not have access to the specified store
(which may or may not exist).
</li>
<li>
<p><a name="error.other"></a><code>other</code>: <code>string</code></p>
<p>Some implementation-specific error has occurred (e.g. I/O)
</li>
</ul>
<h4><a name="key_response"></a><code>record key-response</code></h4>
<p>A response to a <a href="#list_keys"><code>list-keys</code></a> operation.</p>
<h5>Record Fields</h5>
<ul>
<li>
<p><a name="key_response.keys"></a><code>keys</code>: list&lt;<code>string</code>&gt;</p>
<p>The list of keys returned by the query.
</li>
<li>
<p><a name="key_response.cursor"></a><code>cursor</code>: option&lt;<code>u64</code>&gt;</p>
<p>The continuation token to use to fetch the next page of keys. If this is `null`, then
there are no more keys to fetch.
</li>
</ul>
<hr />
<h3>Functions</h3>
<h4><a name="get"></a><code>get: func</code></h4>
<p>A bucket is a collection of key-value pairs. Each key-value pair is stored as a entry in the
bucket, and the bucket itself acts as a collection of all these entries.</p>
<p>It is worth noting that the exact terminology for bucket in key-value stores can very
depending on the specific implementation. For example:</p>
<ol>
<li>Amazon DynamoDB calls a collection of key-value pairs a table</li>
<li>Redis has hashes, sets, and sorted sets as different types of collections</li>
<li>Cassandra calls a collection of key-value pairs a column family</li>
<li>MongoDB calls a collection of key-value pairs a collection</li>
<li>Riak calls a collection of key-value pairs a bucket</li>
<li>Memcached calls a collection of key-value pairs a slab</li>
<li>Azure Cosmos DB calls a collection of key-value pairs a container</li>
</ol>
<p>In this interface, we use the term <code>bucket</code> to refer to a collection of key-value pairs
Get the value associated with the specified <code>key</code></p>
<p>The value is returned as an option. If the key-value pair exists in the
store, it returns <code>Ok(value)</code>. If the key does not exist in the
store, it returns <code>Ok(none)</code>.</p>
<p>If any other error occurs, it returns an <code>Err(error)</code>.</p>
<h5>Params</h5>
<ul>
<li><a name="get.bucket"></a><code>bucket</code>: <code>string</code></li>
<li><a name="get.key"></a><code>key</code>: <code>string</code></li>
</ul>
<h5>Return values</h5>
<ul>
<li><a name="get.0"></a> result&lt;option&lt;list&lt;<code>u8</code>&gt;&gt;, <a href="#error"><a href="#error"><code>error</code></a></a>&gt;</li>
</ul>
<h4><a name="set"></a><code>set: func</code></h4>
<p>Set the value associated with the key in the store. If the key already
exists in the store, it overwrites the value.</p>
<p>If the key does not exist in the store, it creates a new key-value pair.</p>
<p>If any other error occurs, it returns an <code>Err(error)</code>.</p>
<h5>Params</h5>
<ul>
<li><a name="set.bucket"></a><code>bucket</code>: <code>string</code></li>
<li><a name="set.key"></a><code>key</code>: <code>string</code></li>
<li><a name="set.value"></a><code>value</code>: list&lt;<code>u8</code>&gt;</li>
</ul>
<h5>Return values</h5>
<ul>
<li><a name="set.0"></a> result&lt;_, <a href="#error"><a href="#error"><code>error</code></a></a>&gt;</li>
</ul>
<h4><a name="delete"></a><code>delete: func</code></h4>
<p>Delete the key-value pair associated with the key in the store.</p>
<p>If the key does not exist in the store, it does nothing.</p>
<p>If any other error occurs, it returns an <code>Err(error)</code>.</p>
<h5>Params</h5>
<ul>
<li><a name="delete.bucket"></a><code>bucket</code>: <code>string</code></li>
<li><a name="delete.key"></a><code>key</code>: <code>string</code></li>
</ul>
<h5>Return values</h5>
<ul>
<li><a name="delete.0"></a> result&lt;_, <a href="#error"><a href="#error"><code>error</code></a></a>&gt;</li>
</ul>
<h4><a name="exists"></a><code>exists: func</code></h4>
<p>Check if the key exists in the store.</p>
<p>If the key exists in the store, it returns <code>Ok(true)</code>. If the key does
not exist in the store, it returns <code>Ok(false)</code>.</p>
<p>If any other error occurs, it returns an <code>Err(error)</code>.</p>
<h5>Params</h5>
<ul>
<li><a name="exists.bucket"></a><code>bucket</code>: <code>string</code></li>
<li><a name="exists.key"></a><code>key</code>: <code>string</code></li>
</ul>
<h5>Return values</h5>
<ul>
<li><a name="exists.0"></a> result&lt;<code>bool</code>, <a href="#error"><a href="#error"><code>error</code></a></a>&gt;</li>
</ul>
<h4><a name="list_keys"></a><code>list-keys: func</code></h4>
<p>Get all the keys in the store with an optional cursor (for use in pagination). It
returns a list of keys. Please note that for most KeyValue implementations, this is a
can be a very expensive operation and so it should be used judiciously. Implementations
can return any number of keys in a single response, but they should never attempt to
send more data than is reasonable (i.e. on a small edge device, this may only be a few
KB, while on a large machine this could be several MB). Any response should also return
a cursor that can be used to fetch the next page of keys. See the <a href="#key_response"><code>key-response</code></a> record
for more information.</p>
<p>Note that the keys are not guaranteed to be returned in any particular order.</p>
<p>If the store is empty, it returns an empty list.</p>
<p>MAY show an out-of-date list of keys if there are concurrent writes to the store.</p>
<p>If any error occurs, it returns an <code>Err(error)</code>.</p>
<h5>Params</h5>
<ul>
<li><a name="list_keys.bucket"></a><code>bucket</code>: <code>string</code></li>
<li><a name="list_keys.cursor"></a><code>cursor</code>: option&lt;<code>u64</code>&gt;</li>
</ul>
<h5>Return values</h5>
<ul>
<li><a name="list_keys.0"></a> result&lt;<a href="#key_response"><a href="#key_response"><code>key-response</code></a></a>, <a href="#error"><a href="#error"><code>error</code></a></a>&gt;</li>
</ul>
<h2><a name="wrpc_keyvalue_atomics_0_2_0_draft"></a>Export interface wrpc:keyvalue/atomics@0.2.0-draft</h2>
<hr />
<h3>Types</h3>
<h4><a name="error"></a><code>type error</code></h4>
<p><a href="#error"><a href="#error"><code>error</code></a></a></p>
<p>
----
<h3>Functions</h3>
<h4><a name="increment"></a><code>increment: func</code></h4>
<p>Atomically increment the value associated with the key in the store by the given delta. It
returns the new value.</p>
<p>If the key does not exist in the store, it creates a new key-value pair with the value set
to the given delta.</p>
<p>If any other error occurs, it returns an <code>Err(error)</code>.</p>
<h5>Params</h5>
<ul>
<li><a name="increment.bucket"></a><code>bucket</code>: <code>string</code></li>
<li><a name="increment.key"></a><code>key</code>: <code>string</code></li>
<li><a name="increment.delta"></a><code>delta</code>: <code>u64</code></li>
</ul>
<h5>Return values</h5>
<ul>
<li><a name="increment.0"></a> result&lt;<code>u64</code>, <a href="#error"><a href="#error"><code>error</code></a></a>&gt;</li>
</ul>
