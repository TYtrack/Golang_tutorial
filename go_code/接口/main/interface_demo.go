/*
 * @Author: your name
 * @Date: 2021-11-23 15:09:40
 * @LastEditTime: 2021-11-26 12:55:26
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /go_project/src/go_code/接口/main/interface_demo.go
 */

package main

import "fmt"

// 定义一个usb接口
type usb interface {
	// golang的接口中不能有任何变量
	// a int
	start()
	stop()
}

//定义了camera结构体，并实现了usb接口
type camera struct {
}

func (c camera) start() {
	fmt.Println("camera start...")
}

func (c camera) stop() {
	fmt.Println("camera stop...")
}

//定义了phone结构体，并实现了usb接口
type phone struct {
}

func (p phone) start() {
	fmt.Println("phone start...")
}

func (p phone) stop() {
	fmt.Println("phone stop...")
}

func (p phone) call() {
	fmt.Println("phone calling...")
}

//定义了computer结构体
type computer struct {
}

//定义computer的working方法，并传入接口usb
func (c *computer) working(u usb) {
	u.start()
	//类型断言
	p, ok := u.(phone)
	if ok {
		p.call()
	}
	u.stop()
}

func main() {
	var usbs [3]usb

	usbs[0] = camera{}
	usbs[1] = phone{}
	usbs[2] = phone{}

	my_computer := computer{}

	for _, v := range usbs {
		my_computer.working(v)
		fmt.Println()
	}

}
