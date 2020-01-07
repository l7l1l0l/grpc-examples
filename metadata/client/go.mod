module grpc-examples/metadata/client

replace grpc-examples/metadata/pb => ../pb

go 1.13

require (
	github.com/gogo/protobuf v1.3.1 // indirect
	google.golang.org/grpc v1.26.0
	grpc-examples/metadata/pb v0.0.0-00010101000000-000000000000
)
