openssl ecparam -genkey -name secp384r1 -out server.key
openssl req -new -key server.key -out server.csr
openssl x509 -req -sha256 -CA ../../conf/ca.pem -CAkey ../../conf/ca.key -CAcreateserial -days 3650 -in server.csr -out server.pem
#Common Name:go-grpc-example
