/*
 * @Author: your name
 * @Date: 2021-11-18 22:39:33
 * @LastEditTime: 2021-11-26 14:17:03
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /go_project/src/go_code/错误处理/main/error_process.go
 */

package main

import (
	"errors"
	"fmt"
	"strings"
)

//	自定义错误
func readConf(name string) (err error) {
	if !strings.EqualFold("name", name) {
		err = errors.New("没有找到名字")
		return
	} else {
		return nil
	}
}

// 使用panic抛出异常 ，输出异常信息，然后终止程序
func process(name string) {
	err := readConf(name)
	if err != nil {
		panic(err)
	}

}

// 错误处理机制
// 在golang中抛出一个panic 的异常，然后在defer中通过recover来捕获这个异常，然后进行正常处理
// test_1无异常处理
func test_1() {
	num1, num2 := 10, 0
	num3 := num1 / num2
	fmt.Println(num3)
}

//有异常处理，使用
func test_2() {
	defer func() {
		err := recover() //recover能够捕获异常
		if err != nil {
			fmt.Println("err is ", err)
		}
	}()
	num1, num2 := 10, 0
	num3 := num1 / num2
	fmt.Println(num3)
}

func main() {
	//test_1()
	process("name")
	test_2()
	fmt.Println("hello world")
}
