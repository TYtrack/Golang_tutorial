rm -rf *.pem
###
 # @Author: TYtrack
 # @Description: ...
 # @Date: 2022-01-19 00:24:02
 # @LastEditors: TYtrack
 # @LastEditTime: 2022-01-22 18:13:26
 # @FilePath: /grpc_demo/cert/gen.sh
### 

#1、生成CA私钥以及CA证书
openssl req -x509 -newkey rsa:4096 -days 365 -nodes -keyout ca-key.pem -out ca-cert.pem -subj "/C=CN/ST=HN/L=ChangDe/O=HUST/OU=CNCC/CN=*.edu.cn/emailAddress=z@163.com"
# 打印CA证书
openssl x509 -in ca-cert.pem -noout -text

#2、 生成server私钥并生成一个CSR certificate signing request
openssl req -newkey rsa:4096  -nodes  -keyout server-key.pem -out server-req.pem -subj "/C=CN/ST=HB/L=WuHan/O=Home/OU=Bl/CN=*.hust.edu.cn/emailAddress=zt163.com"


#3、通过CA生成服务器的证书
openssl x509 -req -in server-req.pem -CA ca-cert.pem -days 69 -CAkey ca-key.pem -CAcreateserial -out server-cert.pem -extfile server-ext.cnf
openssl x509 -in server-cert.pem -noout -text


#4、 生成client私钥并生成一个CSR certificate signing request
openssl req -newkey rsa:4096  -nodes  -keyout client-key.pem -out client-req.pem -subj "/C=US/ST=NY/L=NY/O=ZZZ/OU=BlZZZ/CN=*.mit.com/emailAddress=zt.@gmailcom"

#5 、通过CA生成客户端的证书
openssl x509 -req -in client-req.pem -CA ca-cert.pem -days 69 -CAkey ca-key.pem -CAcreateserial -out client-cert.pem -extfile client-ext.cnf
openssl x509 -in client-cert.pem -noout -text


#4、验证服务器的证书是否有效
openssl verify -CAfile ca-cert.pem server-cert.pem
openssl verify -CAfile ca-cert.pem client-cert.pem
