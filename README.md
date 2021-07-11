### Command to compile all proto files inside a folder
protoc --proto_path=proto proto/*.proto --go_out=pb


### Had to execute the follow commanda to install a dev tool to compile protoc
## sudo apt install golang-goprotobuf-dev


Had to execute this follow command to make PATH understand the plugin

export PATH="$PATH:$(go env GOPATH)/bin"


## Executing command to generate the gRPC stubs
protoc --proto_path=proto proto/*.proto --go_out=pb --go-grpc_out=pb

##### Using *evans* to generate a pretty gRPC client
https://github.com/ktr0731/evans


## Installing evans
Copy the binary file to /usr/bin/evans (create the folder as sudo if not exists), then add to PATH typing **export PATH=$PATH:/usr/bin/evans**


#### Command to connect on a gRPC by evans and test a call request
evans -r repl --host localhost --port 50001





## Adding aliases in bash to help with commands

```
Type nano ~/.bash_aliases
```
then add the follow aliases

### Compiling protobuffers
alias grpc-compile='protoc --proto_path=proto proto/*.proto --go_out=pb --go-grpc_out=pb'
