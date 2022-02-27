/*
 * @Author: TYtrack
 * @Description: ...
 * @Date: 2022-01-05 14:02:44
 * @LastEditors: TYtrack
 * @LastEditTime: 2022-01-05 14:12:27
 * @FilePath: /context_demo/context_demo_01.go
 */
package main

import (
	"fmt"
	"sync"
	"time"
)

var notify bool
var wg sync.WaitGroup

func f1() {
	for {
		fmt.Println("hello zzz")
		time.Sleep(time.Millisecond * 500)
		if notify {
			break
		}
	}
	wg.Done()
}

func main() {
	go f1()
	wg.Add(1)
	time.Sleep(time.Second * 5)
	notify = true
	wg.Wait()
}
