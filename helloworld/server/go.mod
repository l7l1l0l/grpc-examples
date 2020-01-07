module grpc-examples/helloworld/server

replace grpc-examples/helloworld/pb => ../pb

go 1.13

require (
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/smallnest/grpc-examples v0.0.0-20170320003128-bd4739524257
	golang.org/x/net v0.0.0-20191209160850-c0dbc17a3553
	google.golang.org/grpc v1.26.0
)
