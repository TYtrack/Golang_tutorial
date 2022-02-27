/*
 * @Author: TYtrack
 * @Description: ...
 * @Date: 2022-02-07 21:13:15
 * @LastEditors: TYtrack
 * @LastEditTime: 2022-02-07 21:24:59
 * @FilePath: /go_code/交替打印字符数字/turn_print.go
 */

package main

import (
	"fmt"
	"sync"
)

func main() {
	wait_group := sync.WaitGroup{}
	number, alpha := make(chan bool), make(chan bool)
	wait_group.Add(1)
	go func() {
		i := 1
		for {
			select {
			case <-number:
				fmt.Print(i)
				i++
				fmt.Print(i)
				i++

				alpha <- true

			}

		}

	}()

	go func(wait *sync.WaitGroup) {
		i := 'A'
		for {
			select {
			case <-alpha:
				if i >= 'Z' {
					wait.Done()
					return
				}
				fmt.Print(string(i))
				i++
				fmt.Print(string(i))
				i++
				number <- true
			}

		}

	}(&wait_group)
	number <- true

	wait_group.Wait()
}
