/*
 * @Author: your name
 * @Date: 2021-12-01 21:09:03
 * @LastEditTime: 2021-12-02 14:36:13
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /go_project/src/go_code/聊天室/client/main/main.go
 */
package main

import (
	"fmt"
	"go_code/聊天室/client/process"
	"os"
)

func exit() {
	fmt.Println("----------------------退出MyRoom聊天室----------------------")
	os.Exit(0)
}

func printCai() {
	fmt.Println("----------------------欢迎进入MyRoom聊天室----------------------")
restart:
	fmt.Println("                          1、登陆")
	fmt.Println("                          2、注册")
	fmt.Println("                          3、退出")
	fmt.Println("                         请选择（1-3）：")
	var choose int
	fmt.Scanf("%d\n", &choose)
	switch choose {
	case 1:
		up := &process.UserProcess{}
		up.Login()
	case 2:
		up := &process.UserProcess{}
		err := up.Register()
		if err == nil {
			fmt.Println("                注册成功")
			printCai()
		} else {
			goto restart
		}
	case 3:
		exit()
	default:
		fmt.Println("                输入错误，请重新选择选项")
		goto restart
	}

}

func main() {
	printCai()
}
