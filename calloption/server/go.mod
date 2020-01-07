module grpc-examples/calloption/server

replace grpc-examples/calloption/pb => ../pb

go 1.13

require (
	golang.org/x/net v0.0.0-20191209160850-c0dbc17a3553
	google.golang.org/grpc v1.26.0
	grpc-examples/calloption/pb v0.0.0-00010101000000-000000000000
)
