/*
 * @Author: your name
 * @Date: 2021-11-25 23:39:05
 * @LastEditTime: 2021-11-26 09:33:07
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /go_project/src/go_code/管道/main/chanbel_demo.go
 */

package main

import "fmt"

//管道读取整数
func channel_demo() {
	var ch chan int = make(chan int, 4)
	ch <- 1
	ch <- 3
	ch <- 7
	fmt.Printf("the len is %v,the cap is %v\n", len(ch), cap(ch))

	num1 := <-ch
	num2 := <-ch
	num3 := <-ch
	// 在没有协程的情况下，多一个会报错
	// fatal error: all goroutines are asleep - deadlock!
	// num4 := <-ch
	fmt.Printf("%v, %v, %v\n", num1, num2, num3)

}

//管道读取map
func channel_map() {
	var ch chan map[string]string = make(chan map[string]string, 3)
	m1 := make(map[string]string, 3)
	m1["地点"] = "北京"
	m1["经典"] = "天安门"
	ch <- m1
	m2 := <-ch
	fmt.Println(m2)

}

type Stu struct {
	name  string
	grade float64
}

//向管道传入任意类型
func channel_interface() {
	var ch chan interface{} = make(chan interface{}, 4)
	s1 := Stu{"Tim", 78.5}

	ch <- s1
	ch <- 88
	ch <- 78.8
	s2 := <-ch
	if s3, ok := s2.(Stu); ok {
		fmt.Printf("the name is %v\n", s3.name)
	}
}

//使用关键字close关闭管道写
func channel_close() {
	var ch chan interface{} = make(chan interface{}, 4)
	s1 := Stu{"Tim", 78.5}

	ch <- s1
	ch <- 88
	ch <- 78.8

	s2 := <-ch
	// 关闭管道之后，不能写数据，但是可以读数据
	close(ch)
	if s3, ok := s2.(Stu); ok {
		fmt.Printf("the name is %v\n", s3.name)
	}
	s4 := <-ch
	fmt.Printf("the s4 is %v\n", s4)

}

//管道的遍历
func channel_range() {
	var ch chan interface{} = make(chan interface{}, 4)
	s1 := Stu{"Tim", 78.5}

	ch <- s1
	ch <- 88
	ch <- 78.8

	//使用range遍历前一定要close管道，管道的range遍历只有一个参数(没有index)
	close(ch)
	for v := range ch {
		fmt.Printf("the value is %v\n", v)
	}

}

func main() {
	// channel_demo()
	// channel_map()
	// channel_interface()
	// channel_close()
	channel_range()
}
