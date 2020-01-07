module grpc-examples/calloption/client

replace grpc-examples/calloption/pb => ../pb

go 1.13

require (
	google.golang.org/grpc v1.26.0
	grpc-examples/calloption/pb v0.0.0-00010101000000-000000000000
)
