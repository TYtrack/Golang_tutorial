/*
 * @Author: your name
 * @Date: 2021-12-07 15:14:53
 * @LastEditTime: 2021-12-13 13:18:28
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /goproject/src/go_code/系统自带rpc/simple_demo/client/client_demo.go
 */

package main

import (
	"fmt"
	"net/rpc"
)

type Params struct {
	X, Y int
}

func main() {
	client, _ := rpc.Dial("tcp", "localhost:1234")
	var reply string
	client.Call("HelloService.Hello", "tyty", &reply)
	fmt.Println(reply)

	p := Params{19, 5}
	var ret int
	client.Call("MatttttService.Chenfa", p, &ret)
	fmt.Println(ret)

	client.Call("MatttttService.Quyu", p, &ret)
	fmt.Println(ret)
}
