/*
 * @Author: your name
 * @Date: 2021-11-29 23:13:07
 * @LastEditTime: 2021-11-30 00:09:35
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /go_project/src/go_code/网络编程/cs_demo_01/client/client.go
 */

package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("接受连接失败：", err)
		return
	}
	fmt.Println("连接成功：", conn)
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)

	for {
		data, err2 := reader.ReadString('\n')
		if err2 != nil {
			fmt.Println("出现错误", err)
			return
		}
		data = strings.Trim(data, " \r\n")
		if string(data) == "exit" {
			break
		}

		conn.Write([]byte(data + "\n"))
	}

}
