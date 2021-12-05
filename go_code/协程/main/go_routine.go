/*
 * @Author: your name
 * @Date: 2021-11-25 15:00:58
 * @LastEditTime: 2021-11-25 22:31:30
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /go_project/src/go_code/协程/main/go_routine.go
 */
package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

/* 协程四个特点
1、有独立栈空间
2、共享堆空间
3、调度由用户控制
4、协程是轻量级线程【编译器优化】
*/

func calSu(left int, right int) {
	count := 0
	for i := left; i <= right; i++ {
		var flag bool = false
		for j := 2; j <= int(math.Sqrt(float64(i))); j++ {
			if i%j == 0 {
				flag = true
				break
			}

		}
		if flag == false {
			count++
		}

	}
	fmt.Printf("%v \n", count)
}

func printHW() {
	for i := 0; i < 10; i++ {
		fmt.Println("hello world")
		time.Sleep(100 * time.Millisecond)
	}
}

// 使用协程来计算累和
var m1 map[int]int = make(map[int]int)
var lock sync.Mutex

//低水平代码用锁，高水平用channel
func leihe(num int) {
	res := 1
	for i := 2; i <= num; i++ {
		res += i
	}
	//加了锁还会出现资源race，除非在读的时候也加上锁
	lock.Lock()
	m1[num] = res
	lock.Unlock()
}

func main() {
	// 1、calSu（）
	// s1 := time.Now()
	// calSu(1, 5000000)
	// calSu(5000001, 10000000)
	// s3 := time.Now().Sub(s1)
	// fmt.Println(s3)

	// s1 = time.Now()
	// go calSu(1, 2500000)
	// go calSu(2500001, 5000000)
	// go calSu(5000001, 7500000)
	// calSu(7500001, 10000000)
	// s3 = time.Now().Sub(s1)
	// fmt.Println(s3)

	//2、printHW
	// go printHW()
	// for i := 0; i < 10; i++ {
	// 	fmt.Println("hello goroutine")
	// 	time.Sleep(100 * time.Millisecond)
	// }

	//3、leihe（）
	// 会出现fatal error: concurrent map writes
	// 可以使用go build -race go_rountine.go来查看是否出现资源竞争

	for i := 2; i <= 100; i++ {
		go leihe(i)
	}

	time.Sleep(time.Second)
	// 为什么这里要加锁，主线程不知道底层是否占有锁？？
	lock.Lock()
	fmt.Println(m1)
	lock.Unlock()

}
