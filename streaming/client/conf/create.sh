openssl ecparam -genkey -name secp384r1 -out client.key
openssl req -new -key client.key -out client.csr
openssl x509 -req -sha256 -CA ../../conf/ca.pem -CAkey ../../conf/ca.key -CAcreateserial -days 3650 -in client.csr -out client.pem

