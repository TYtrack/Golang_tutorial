/*
 * @Author: TYtrack
 * @Description: ...
 * @Date: 2022-01-05 14:02:44
 * @LastEditors: TYtrack
 * @LastEditTime: 2022-01-05 14:20:20
 * @FilePath: /context_demo/02_demo/context_demo_02.go
 */
package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var exitChan chan bool

func f1() {
	defer wg.Done()

zzlabel:
	for {
		fmt.Println("hello zzz")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-exitChan:
			break zzlabel
		default:

		}
	}

}

func main() {
	exitChan = make(chan bool)
	go f1()
	wg.Add(1)
	time.Sleep(time.Second * 5)
	exitChan <- true
	wg.Wait()
}
