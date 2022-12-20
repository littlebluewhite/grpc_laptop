gen:
	protoc --proto_path=proto proto/*.proto --go_out=. --go-grpc_out=require_unimplemented_servers=false:.
clean:
	rm pb/message/*.go
server:
	go run cmd/server/main.go -port 8080
client:
	go run cmd/client/main.go -address 0.0.0.0:8080
test:
	go test -cover -race ./...