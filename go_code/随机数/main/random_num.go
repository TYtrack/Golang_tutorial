/*
 * @Author: your name
 * @Date: 2021-11-17 21:05:12
 * @LastEditTime: 2021-11-17 21:14:07
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /go_project/src/go_code/随机数/main/random_num.go
 */
package main

import (
	"fmt"
	"math/rand"
	"time"
)

//随机选一个数，从1到99隨即選擇數字，直到99break出来
// 使用rand.Seed设置随机种子
// 使用rand.Intn选择从[0,n)的数字
func lucky_99() {
	rand.Seed(time.Now().Unix())
	for {
		temp := rand.Intn(100)
		fmt.Println(temp)
		if temp == 99 {
			break
		}
	}

}

func main() {
	lucky_99()
}
