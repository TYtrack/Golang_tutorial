/*
 * @Author: your name
 * @Date: 2021-11-26 10:43:53
 * @LastEditTime: 2021-11-26 10:58:28
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /go_project/src/go_code/管道/timu_1/leihe_channel.go
 */
package main

import "fmt"

/*
	题目：
		一个协程写数据1——2000
		另外开启八个协程取出数据，求累和，然后加入到resChan中
*/

func writeData(numChan chan<- int) {
	for i := 1; i <= 2000; i++ {
		numChan <- i
	}
	close(numChan)
}

type Lei struct {
	index int
	value int
}

func processData(numChan <-chan int, resChan chan<- Lei) {
	for {
		num, ok := <-numChan
		if !ok {
			break
		}
		res := 0
		for i := 1; i <= num; i++ {
			res += i
		}
		resChan <- Lei{num, res}
	}
	close(resChan)

}

func main() {
	numChan := make(chan int)
	resChan := make(chan Lei)
	go writeData(numChan)
	for i := 1; i <= 8; i++ {
		go processData(numChan, resChan)
	}
	for v := range resChan {
		fmt.Printf("res[%v]=%v , ", v.index, v.value)
	}

}
