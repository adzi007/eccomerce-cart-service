build_proto:
	protoc --go_out=./cart_proto --go_opt=paths=source_relative \
    --go-grpc_out=./cart_proto --go-grpc_opt=paths=source_relative \
    cart.proto