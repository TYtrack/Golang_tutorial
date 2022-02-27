/*
 * @Author: your name
 * @Date: 2021-12-21 14:31:00
 * @LastEditTime: 2021-12-21 15:22:50
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /src/go_code/golang的坑/hole_01/main.go
 */

//关于协程的全局变量

package main

import (
	"fmt"
	"time"
)

func main() {

	for k := 0; k < 4; k++ {
		i, j := 10, 20
		if k%2 == 0 {
			i = 20
			j = 10
		} else {
			i = 10
			j = 20
		}
		fmt.Printf("@@@ %v  -> %v\n", i, j)
		go func() {
			fmt.Printf("^^^ %v  -> %v\n", i, j)
		}()
	}
	time.Sleep(time.Second * 3)
}
