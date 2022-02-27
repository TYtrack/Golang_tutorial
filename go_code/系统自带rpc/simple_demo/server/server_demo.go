/*
 * @Author: your name
 * @Date: 2021-12-07 15:15:01
 * @LastEditTime: 2021-12-13 13:20:37
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /goproject/src/go_code/系统自带rpc/simple_demo/server/server_demo.go
 */

package main

import (
	"net"
	"net/rpc"
)

// 结构体字段首字母要大写，可以别人调用

// 函数名必须首字母大写

// 函数第一参数是接收参数，第二个参数是返回给客户端的参数，必须是指针类型

// 函数还必须有一个返回值error

//实现两个服务，共三个功能

type HelloService struct {
}

func (this *HelloService) Hello(request string, response *string) (err error) {
	*response = "hello" + request
	return
}

type MathService struct {
}
type Params struct {
	X, Y int
}

func (this *MathService) Chenfa(p Params, ret *int) (err error) {
	*ret = p.X * p.Y
	return
}

func (this *MathService) Quyu(p Params, ret *int) (err error) {
	*ret = p.X % p.Y
	return
}

func main() {
	rpc.RegisterName("HelloService", new(HelloService))
	rpc.RegisterName("MatttttService", new(MathService))

	listener, _ := net.Listen("tcp", ":1234")

	conn, _ := listener.Accept()

	rpc.ServeConn(conn)
}
