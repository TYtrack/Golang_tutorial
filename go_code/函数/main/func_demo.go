/*
 * @Author: your name
 * @Date: 2021-11-18 12:05:55
 * @LastEditTime: 2021-11-18 12:49:30
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /go_project/src/go_code/函数/main/func_demo.go
 */
package main

import "fmt"

// 对函数的返回值进行命名
// 没有返回值
func add_multi(a, b int) (sum, multi int) {
	sum = a + b
	multi = a * b
	return
}

// 可变参数必须放在参数列表的最后一个位置
// 可以通过len(args)来得到可变参数的长度
func sum_demo(args ...int) int {
	res := 0
	for i := 0; i < len(args); i++ {
		res += args[i]
	}
	return res
}

// 将函数作为变量，并作为一个参数传入另一个函数
type FuncVar func(int, int) int

func chu(a, b int) int {
	if b == 0 {
		return 0
	}
	return a / b
}

func Func_Para(funcvar FuncVar, x int, y int) int {
	return funcvar(x, y)
}

//引用传递，交换值
func swap(i, j *int) {
	*i, *j = *j, *i
}

//init函数，自动调用，在main函数之前执行，在全局变量定义之后执行
func init() {
	fmt.Println("init")
}

func main() {
	//命名返回值
	fmt.Println(add_multi(3, 7))

	//可变参数
	fmt.Println(sum_demo(5, 1, 3, 5, 7, 9))

	//函数变量
	fmt.Println(Func_Para(chu, 10, 3))

	//引用传递交换值
	var i, j = 3, 7
	swap(&i, &j)
	fmt.Printf("swap result is i:%v ,j:%v \n", i, j)

	// 匿名函数demo1，只调用一次
	func(x, y int) int {
		fmt.Println("anonymous_1 fun res is", x+y)
		return x + y
	}(3, 6)

	// 匿名函数demo2，调用多次
	anonymous_2 := func(x, y int) int {
		fmt.Println("anonymous_2 fun res is", x+y)
		return x + y
	}
	anonymous_2(5, 7)
}
