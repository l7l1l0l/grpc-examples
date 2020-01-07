module grpc-examples/auth/client

replace grpc-examples/auth/pb => ../pb

go 1.13

require (
	github.com/gogo/protobuf v1.3.1 // indirect
	google.golang.org/grpc v1.26.0
	grpc-examples/auth/pb v0.0.0-00010101000000-000000000000
)
