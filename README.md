# go-api
protoc --go_out=pkg/proto/pb --go_opt=paths=source_relative \
--go-grpc_out=pkg/proto/pb --go-grpc_opt=paths=source_relative \
api.proto
