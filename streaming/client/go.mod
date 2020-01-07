module grpc-examples/streaming/client

replace grpc-examples/streaming/pb => ../pb

go 1.13

require (
	google.golang.org/grpc v1.26.0
	grpc-examples/streaming/pb v0.0.0-00010101000000-000000000000
)
