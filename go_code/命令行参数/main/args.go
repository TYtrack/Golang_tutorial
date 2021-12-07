/*
 * @Author: your name
 * @Date: 2021-11-24 19:53:00
 * @LastEditTime: 2021-11-24 20:10:18
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /go_project/src/go_code/命令行参数/main/args.go
 */

package main

import (
	"flag"
	"fmt"
	"os"
)

// 缺点是要严格按照参数的顺序 且解析参数不方便
func getArgs() {
	// os.Args得到的结果是切片
	// 注意的是./args也是命令行参数
	s1 := os.Args
	fmt.Println(len(s1))
	fmt.Println(s1)
	fmt.Printf("类型是%T %T  ,   值是%v %v\n", s1[0], s1[1], s1[0], s1[1])
}

// 使用flag包来解析参数，参数的顺序可以随意
func flag_Args() {
	// 定义几个变量，来获取命令行参数

	var user string
	var password string
	var host string
	var port int
	// 参数分别是 接受用户命令行中输入的 -u 后面的参数值，第三个参数是默认值，第四个是说明
	flag.StringVar(&user, "u", "", "用户名，默认为空")
	flag.StringVar(&password, "pwd", "", "密码，默认为空")
	flag.StringVar(&host, "h", "localhost", "主机，默认为localhost")
	flag.IntVar(&port, "port", 81, "端口，默认81")

	//非常重要的，将Args的参数转化到flag中
	flag.Parse()

	fmt.Printf("user=%v,pwd=%v,host=%v,port=%v\n", user, password, host, port)
}

func main() {
	//getArgs()
	flag_Args()
}
