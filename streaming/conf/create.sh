openssl genrsa -out ca.key 2048
openssl req -new -x509 -days 7200 -key ca.key -out ca.pem
#Common Name:go-grpc-example
