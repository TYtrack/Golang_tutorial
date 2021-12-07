/*
 * @Author: your name
 * @Date: 2021-11-17 09:50:49
 * @LastEditTime: 2021-11-17 21:17:46
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /go_project/src/go_code/分支/main/fenzhi_demo.go
 */
package main

import (
	"fmt"
	"math"
)

//使用标签来进行break
func break_demo() {

	for i := 0; i < 10; i++ {
		for j := 0; j < 7; j++ {
			if i+j == 13 {
				fmt.Printf("1\tthe result %v %v\n", i, j)
				break
			}
		}
	}
kk:
	for i := 0; i < 10; i++ {
		for j := 0; j < 7; j++ {
			if i+j == 13 {
				fmt.Printf("2\tthe result %v %v\n", i, j)
				break kk
			}
		}
	}

}

func switch_demo() {
	//每个case可以有多个表达式
	c := 6
	switch c {
	case 2, 3, 4:
		fmt.Println("ok1")
	case 6, 7, 8:
		fmt.Println("ok3")
	}

	//switch之后可以没有表达式，case之后的可以是条件判断：
	// 使用fallthrough可以穿透到下一个语句
	switch {
	case c < 8 && c > 5:
		fmt.Println("ok6")
		fallthrough
	case c < 5 && c > 1:
		fmt.Println("ok7")
	default:
		fmt.Println("ok9")
	}

	// 可以用作判断某个interface变量实际指向的类型
	var x interface{}
	var y = 10
	x = y
	switch x.(type) {
	case nil:
		fmt.Println("是 nil")
	case bool:
		fmt.Println("是 bool")
	case int:
		fmt.Println("是 int")
	case float64:
		fmt.Println("是 float64")
	}
}

func if_details() {
	var i = 3
	// 1、判断的括号可以有，可以没有，go官方说最好没有
	// 2、如果没有{}会报错，就算只有一行语句也需要花括号
	// 3、else也不能换行，要紧挨着右括号
	// 4、条件判断语句必须是条件表达式，不能像C++一样可以是赋值表达式
	if i < 4 {
		fmt.Println("ok")
	} else {
		fmt.Println("no")
	}
}

func timu1() {
	var a, b int
	fmt.Scanf("%v%v", &a, &b)
	if a+b > 50 {
		fmt.Println("hello world")
	}
}

func timu2() {
	var a, b float64
	fmt.Scanf("%f%f", &a, &b)
	if a > 10 && b < 20 {
		fmt.Println(a + b)
	}

}

// 判断是否为闰年
func runnian(i int) {
	if (i%4 == 0 && i%100 != 0) || i%400 == 0 {
		fmt.Println(i, "是闰年")
	} else {
		fmt.Println(i, "不是闰年")
	}

}

// 求二元一次的根
func gen(a, b, c float64) {
	var temp float64 = b*b - 4*a*c
	if temp == 0 {
		res := -b / float64(2*a)
		fmt.Println("有一个根:", res)
	} else if temp < 0 {
		fmt.Println("没有根")
	} else {
		res1 := (-b + math.Sqrt(temp)) / float64(2*a)
		res2 := (-b - math.Sqrt(temp)) / float64(2*a)
		fmt.Println("有两个根:", res1, " ", res2)

	}

}

func main() {
	break_demo()
	switch_demo()
	if_details()
	// timu1()
	// timu2()
	runnian(204)
	runnian(200)
	runnian(2000)
	gen(1, 25, 0)

}
