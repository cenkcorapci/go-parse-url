api/api.pb.go: api/api.proto
		protoc --go_out=plugins=grpc:api/. api/api.proto
