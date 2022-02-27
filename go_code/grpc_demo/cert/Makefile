rm -rf *.pem

openssl req -x509 -newkey rsa:4096 -days 365 -keyout ca-key.pem -out ca-cert.pem -subj "C=CN/ST=HN/L=ChangDe/O=HUST/OU=CNCC/CN=*.cnn.hust.edu.cn/emailAddress=zhengzhoazuo@163.com"

openssl x509 -in ca-cert.pem -noout -text