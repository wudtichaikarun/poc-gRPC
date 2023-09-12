## Install protobuf

```bash
brew install protobuf
go get google.golang.org/grpc
go get github.com/golang/protobuf/protoc-gen-go

```

## Define the Protocol Buffers (protobuf) schema in a file named calculator.proto:

```proto
syntax = "proto3";

package calculator;

service Calculator {
  rpc Add(AddRequest) returns (AddResponse);
}

message AddRequest {
  int32 num1 = 1;
  int32 num2 = 2;
}

message AddResponse {
  int32 result = 1;
}
```

## Generate Go code from the protobuf file

```bash
$ protoc --go_out=plugins=grpc:. calculator.proto
```
