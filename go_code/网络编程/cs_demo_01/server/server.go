/*
 * @Author: your name
 * @Date: 2021-11-29 22:41:47
 * @LastEditTime: 2021-11-29 23:57:27
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /go_project/src/go_code/网络编程/cs_demo_01/server/server.go
 */

package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	fmt.Println("进入协程")

	for {
		var data []byte = make([]byte, 512)
		n, err := conn.Read(data)
		if err != nil {
			fmt.Println("发生错误:", err)
			break
		}
		if n == 0 {
			fmt.Println("数据长度为0")
			break
		}
		fmt.Printf("数据长度%v：%v", n, string(data[:n]))
	}

}

func main() {
	fmt.Println("开始监听")
	listener, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("监听失败：", err)
		return
	}
	//  Close关闭该接口，并使任何阻塞的Accept操作都会不再阻塞并返回错误。
	defer listener.Close()

	for {
		conn, err2 := listener.Accept()
		if err2 != nil {
			fmt.Println("接受连接失败：", err)
			continue
		}
		fmt.Println("连接成功：", conn)
		fmt.Println("客户端IP：", conn.RemoteAddr())

		go process(conn)

	}

}
