/*
 * @Author: your name
 * @Date: 2021-11-26 09:34:01
 * @LastEditTime: 2021-11-26 10:17:21
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /go_project/src/go_code/管道/goroutineAndChannel/goroutine_comm.go
 */
package main

import "fmt"

//开启一个writeDate协程，向管道intChan写入五十个数据
//开启一个readDate协程，从管道intChan读取writeDate写的数据
// 主线程要等到两个协程都完成才能退出

func writeData(intChan chan<- int) {
	for i := 1; i <= 50; i++ {
		intChan <- (i * 2)
	}
	// 一定要close，不然readData不能
	close(intChan)
}

func readData(intChan <-chan int, exitChan chan<- bool) {

	for {
		// close之后ok会返回false
		num, ok := <-intChan
		if !ok {
			break
		}
		fmt.Printf("%v. ", num)
	}
	exitChan <- true
	close(exitChan)
}

func main() {
	intChan := make(chan int, 1)
	exitChan := make(chan bool)
	go writeData(intChan)
	go readData(intChan, exitChan)

	<-exitChan

}
