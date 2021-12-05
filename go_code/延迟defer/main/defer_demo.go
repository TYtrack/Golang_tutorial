/*
 * @Author: your name
 * @Date: 2021-11-18 13:27:35
 * @LastEditTime: 2021-11-18 13:39:01
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /go_project/src/go_code/延迟defer/main/defer_demo.go
 */
package main

import "fmt"

//会把defer语句压入到独立的栈中，暂时不执行
// 当函数执行完别之后，从栈取出执行
// 当将语句压入栈中，相应的’值拷贝‘也将被压入栈中
func defer_demo(x, y int) {

	defer fmt.Println("defer ok1 ", x)
	defer fmt.Println("defer ok2 ", y)
	x++
	y++
	z := x + y
	fmt.Println("defer ok3 ", z)
}

func main() {
	defer_demo(3, 7)

}
