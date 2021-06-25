# To run the plugin:
1. build it: 
```
pushd pkg/plugin
go build
```
2. Export the addr:
```
export GF_PLUGIN_GRPC_ADDRESS_PRINT=127.0.0.1:<port>
``` 
3. `./plugin --standalone=true`

# To run the client:
1. build it:
```
pushd pkg/client
go build
```
2. Export the addr (same as plugin):
```
export PLUGIN_ADDR=127.0.0.1:<port>
```
3. `./client your-message`

Message will be printed to the plugin's stdout 
