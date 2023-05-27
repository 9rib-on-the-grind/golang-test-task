protoc \
	-I rpc \
	--go_out=rpc \
	--go_opt=paths=source_relative \
	--go-grpc_out=rpc \
	--go-grpc_opt=paths=source_relative \
	rpc/rpc.proto