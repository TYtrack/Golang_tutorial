/*
 * @Author: your name
 * @Date: 2021-11-26 14:03:41
 * @LastEditTime: 2021-11-26 14:12:35
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /go_project/src/go_code/管道/select_demo/select_demo.go
 */

package main

import (
	"fmt"
	"strconv"
)

func main() {
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}
	strChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		strChan <- "hello " + strconv.Itoa(i)
	}

	// 传统的方法在遍历管道时，如果不close会阻塞导致deadlock
	// 问题：在实际需求中，不好确认什么时候关闭管道
	// 解决办法，使用select
for_label:
	for {
		select {
		//注意，这里如果intChan没有关闭，不会阻塞导致deadlock，而会往下个case走
		case v := <-intChan:
			fmt.Printf("读取了数据:%v\n", v)
		case v := <-strChan:
			fmt.Printf("读取了数据:%v\n", v)
		default:
			fmt.Printf("读取不到数据了\n")
			break for_label
		}
	}
}
