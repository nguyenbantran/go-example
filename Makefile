proto:
	protoc --go_out=pb  --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/*.proto

clean:
	rm pb/*.go