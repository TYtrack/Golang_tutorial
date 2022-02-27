/*
 * @Author: TYtrack
 * @Description: ...
 * @Date: 2022-01-05 14:02:44
 * @LastEditors: TYtrack
 * @LastEditTime: 2022-01-05 14:45:22
 * @FilePath: /context_demo/03_demo/context_demo_03.go
 */
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func f1(ctx context.Context) {
	defer wg.Done()

zzlabel:
	for {
		fmt.Println("hello zzz")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-ctx.Done():
			break zzlabel
		default:

		}
	}

}

func main() {

	ctx, cancel := context.WithCancel(context.Background())

	go f1(ctx)
	wg.Add(1)
	time.Sleep(time.Second * 5)

	// 进行通知
	cancel()
	wg.Wait()
}
