gen:
	protoc --proto_path=proto proto/*.proto --go_out=.
clean:
	rm pb/message/*.go
run:
	go run main.go