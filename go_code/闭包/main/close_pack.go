/*
 * @Author: your name
 * @Date: 2021-11-18 12:53:55
 * @LastEditTime: 2021-11-18 13:15:46
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /go_project/src/go_code/闭包/main/close_pack.go
 */
package main

import (
	"fmt"
	"strings"
)

func AddUpper() func(x int) int {
	var n int = 10
	return func(x int) int {
		n += x
		return n
	}
}

// 编写一个函数 makeSuffix 可以接收一个文件后缀名(比如.jpg)，并返回一个闭包
// 调用闭包，可以传入一个文件名，如果该文件名没有指定的后缀(比如.jpg) ,则返回 文件名.jpg , 如
func makeSuffix(suffix string) func(string) string {
	return func(s string) string {
		if strings.HasSuffix(s, suffix) {
			return strings.TrimSuffix(s, suffix)
		} else {
			return s + suffix
		}
	}
}

func main() {
	z := AddUpper()

	fmt.Println(z(6))
	fmt.Println(z(7))
	fmt.Println(z(8))
	fmt.Println(z(9))

	tt := makeSuffix(".jpg")
	fmt.Println(tt("风景画"))
	fmt.Println(tt("肖像画.jpg"))
	fmt.Println(tt("全家福.png"))

}
