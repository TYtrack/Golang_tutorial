/*
 * @Author: your name
 * @Date: 2021-11-22 12:11:03
 * @LastEditTime: 2021-11-22 13:00:23
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /go_project/src/go_code/方法/main/method.go
 */

package main

import "fmt"

type A struct {
	num  int
	Desc string
}

// 重写
// 为一个结构体实现String方法，返回值是string
// 在fmt.Println会默认输出String方法的返回值
func (a A) String() string {
	res := fmt.Sprintf("the num is %v, the des is %v", a.num, a.Desc)
	return res
}

// 结构体A的一个方法，值拷贝不会修改结构体的内容
func (a A) printDesc() {
	a.Desc = "world"
	fmt.Println(a.Desc)
}

func (a A) jisuan() (res int) {
	i := 1
	for {
		if i > 1000 {
			break
		} else {
			res += i
			i++
		}
	}
	return
}

func (a A) jisuan2(n int) (res int) {
	i := 1
	for {
		if i > n {
			break
		} else {
			res += i
			i++
		}
	}
	return
}

func (a A) jisuan3(n ...int) (res int) {
	i := 0
	for {
		if i == len(n) {
			break
		} else {
			res += n[i]
			i++
		}
	}
	return
}
func main() {
	a := A{3, "hello"}
	fmt.Println(a)
	a.printDesc()
	fmt.Println(a.Desc)
	fmt.Println("jisuan res is", a.jisuan())
	fmt.Println("jisuan2 res is", a.jisuan2(6))
	fmt.Println("jisuan3 res is", a.jisuan3(1, 4, 6, 2, 6, 8))
}
