/*
 * @Author: your name
 * @Date: 2021-11-10 19:13:57
 * @LastEditTime: 2021-11-10 19:22:53
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /go_project/src/go_code/键盘输入/main/keynoard_input.go
 */
package main

import "fmt"

func main() {
	// 第一种输入Scanf
	var i, j int
	fmt.Scanf("%d %d", &i, &j)
	fmt.Printf("输入的是%d,%d", i, j)

	//第二种输入Scanln
	var name string
	var age byte
	fmt.Println("请输入姓名")
	fmt.Scanln(&name)

	fmt.Println("请输入年龄")
	fmt.Scanln(&age)

	fmt.Printf("输入的是%v,%v", name, age)

	/*
		格式化输入是Scanf Printf
		直接输入输出一行是Scanln Println
	*/

}
