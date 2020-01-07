module grpc-examples/interceptor/client

replace grpc-examples/interceptor/pb => ../pb

go 1.13

require (
	github.com/gogo/protobuf v1.3.1 // indirect
	golang.org/x/net v0.0.0-20191209160850-c0dbc17a3553
	google.golang.org/grpc v1.26.0
	grpc-examples/interceptor/pb v0.0.0-00010101000000-000000000000
)
