gen:
	protoc --proto_path=proto proto/*.proto --go-grpc_out=pb

clean:
	del pb/*.go

run:
	go run main.go