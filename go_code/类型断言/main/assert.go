/*
 * @Author: your name
 * @Date: 2021-11-23 17:05:18
 * @LastEditTime: 2021-11-23 17:47:34
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /go_project/src/go_code/类型断言/main/assert.go
 */

package main

import (
	"fmt"
)

type Point struct {
	x int
	y int
}

// 不能将空接口赋给一个结构体，需要类型断言
func assert_demo() {
	var a interface{}
	var b = Point{3, 4}
	a = b
	// 不能将空接口赋给一个结构体
	// var d Point = a

	//类型断言，将接口转化为自定义类型使用var.(Type)
	//如果a指向的是Point类型那么就转化为Point类型，如果不是就报错
	var d Point = a.(Point)
	fmt.Println(d)

}

func assert_demo_2() {
	var a interface{}
	var b float32 = 78.2
	a = b
	c := a.(float32)
	fmt.Println(c)
	fmt.Printf("c的类型是%T，值是%v\n", c, c)

	// 抛出异常 interface conversion: interface {} is float32, not int"
	// d := a.(int)
	// fmt.Println(d)

	//如何进行进行断言的时候，进行检测机制
	d, ok := a.(int)
	if ok {
		fmt.Printf("d的类型是%T，值是%v\n", d, d)
	} else {
		fmt.Println("类型错误")
	}

	//另一种写法

	if z, ok_2 := a.(float32); ok_2 {
		fmt.Printf("z的类型是%T，值是%v\n", z, z)
	} else {
		fmt.Println("类型错误~")
	}
}

func assert_demo_3(items ...interface{}) {
	for _, item := range items {
		switch item.(type) {
		case int, int32, int64:
			fmt.Println("item 是一个 int 类型")
		case float64, float32:
			fmt.Println("item 是一个 float 类型")
		case [4]rune:
			fmt.Println("item 是一个 [4]rune 类型")
		case bool:
			fmt.Println("item 是一个 bool 类型")
		case string:
			fmt.Println("item 是一个 string 类型")
		case Point:
			fmt.Println("item 是一个 Point 类型")
		case *Point:
			fmt.Println("item 是一个 *Point 类型")
		default:
			fmt.Println("item 是其他类型")
		}
	}

}

func main() {
	assert_demo()
	assert_demo_2()
	a := [4]rune{'a', 'w', 'x', 'z'}
	b := a[:]
	p := Point{3, 6}

	assert_demo_3(4, 3.2, true, "abcd", a, b, p, &p)
}
