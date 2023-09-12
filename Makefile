run-server:
	go run server/server.go

run-client:
	go run client/client.go

proto:
	protoc --proto_path=calculator --go_out=calculator --go_opt=paths=source_relative \
	--go-grpc_out=calculator --go-grpc_opt=paths=source_relative \
	calculator/*.proto