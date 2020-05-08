# cars-api
*using go 1.14*

## Setup

To generate the GRPC Code:
1. Install `protoc`:

       brew install protoc
       
2. Install `protoc-gen-go`:

       go get -u github.com/golang/protobuf/protoc-gen-go
       
   **Note**: Be sure to install `proto-gen-go` from https://github.com/golang/protobuf. The version of `proto-gen-go` at https://github.com/protocolbuffers/protobuf-go does not include gRPC!
       
3. Download https://github.com/protocolbuffers/protobuf/releases/download/v3.11.4/protoc-3.11.4-osx-x86_64.zip and unzip

        unzip protoc-3.11.4-osx-x86_64.zip .

4.  Copy the files from `include` directory inside the `protoc` zip into the `third_party` directory.

5. Run `protoc` generation from the root of this project:

        ```protoc --proto_path=./proto --proto_path=./third_party --go_out=plugins=grpc:./proto car-service.proto```
