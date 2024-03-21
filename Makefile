gen:
	protoc --proto_path=proto proto/*.proto  --go_out=:pb --go-grpc_out=:pb

clean:
	del .\pb\*.go

testall:
	go test -cover ./...

testone:
	go test ./serializer/file_test.go

run:
	go run main.go