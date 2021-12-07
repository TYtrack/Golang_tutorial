/*
 * @Author: your name
 * @Date: 2021-11-26 14:18:57
 * @LastEditTime: 2021-11-26 14:25:40
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /go_project/src/go_code/管道/panic_recover/panic_recover.go
 */

package main

import (
	"fmt"
	"time"
)

// 一般来说，如果一个协程出现了panic，那么如果不进行处理的话，会导致整个程序崩溃
// 这时候，就可以使用recover来捕获Panic，这样主线程不会崩溃，继续运行

func sayHello() {
	for i := 0; i < 10; i++ {
		fmt.Println("hello world_  ", i)
		time.Sleep(time.Second)
	}
}

func test() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Printf("test err is %v\n", err)
		}
	}()

	var mymap map[int]string
	mymap[0] = "zzz"

}
func main() {
	go sayHello()
	go test()
	time.Sleep(time.Second * 20)

}
