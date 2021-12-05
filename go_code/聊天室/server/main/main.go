/*
 * @Author: your name
 * @Date: 2021-12-01 15:32:14
 * @xLastEditTime: 2021-12-02 11:38:10
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /go_project/src/go_code/聊天室/server/main/main.go
 */

package main

import (
	"fmt"
	"go_code/聊天室/server/model"
	"net"
	"time"
)

func process_1(conn net.Conn) (err error) {
	defer conn.Close()
	pp := &Processor{
		Conn: conn,
	}
	err = pp.process_2()

	return
}

func main() {
	model.InitPool("127.0.0.1:6379", 8, 0, 300*time.Second)
	model.InitUserDao(model.Pool)

	fmt.Println("新的服务器在8989上监听")
	listener, err := net.Listen("tcp", "127.0.0.1:8989")
	defer listener.Close()

	if err != nil {
		fmt.Println("服务器监听失败")
		return
	}
	fmt.Println("服务器监听成功，等待连接")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("服务器接受连接失败")
		}
		go process_1(conn)
	}

}
