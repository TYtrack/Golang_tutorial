/*
 * @Author: your name
 * @Date: 2021-11-19 09:25:10
 * @LastEditTime: 2021-11-22 10:24:15
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /go_project/src/go_code/数组/main/array.go
 */

package main

import (
	"fmt"
	"math/rand"
	"time"
)

//数组的五种定义方式
func array_defination() {
	var a [4]int = [4]int{1, 2, 4, 6}
	var b = [4]int{4, 6, 4, 2}
	var c = [...]int{6, 3, 2, 6, 2, 2, 1}
	var d = [...]int{0: 3, 1: 89, 2: 22}
	e := [4]int{2, 2, 5, 2}

	for _, val := range a {
		fmt.Printf("%v ", val)
	}
	fmt.Println()
	for _, val := range b {
		fmt.Printf("%v ", val)
	}
	fmt.Println()
	for _, val := range c {
		fmt.Printf("%v ", val)
	}
	fmt.Println()
	for _, val := range d {
		fmt.Printf("%v ", val)
	}
	fmt.Println()
	for _, val := range e {
		fmt.Printf("%v ", val)
	}
	fmt.Println()
}

func two_d_array() {
	var a [5][5]int
	a[2][3] = 1
	a[1][2] = 1
	a[3][4] = 1
	fmt.Println(a)

}

// timu1
func timu_1() {
	var bytes [26]byte
	for i, _ := range bytes {
		bytes[i] = byte('A' + i)
		fmt.Printf("%c ", bytes[i])
	}
	fmt.Println()
}

//timu2
func timu_2() {
	var nums = [...]int{63, 475, -423, 557, 33, 643, 23}
	maxn := nums[0]
	for _, v := range nums {
		if maxn < v {
			maxn = v
		}
	}
	fmt.Println("the max value is ", maxn)
}

func timu_5() {
	var nums [5]int
	rand.Seed(time.Now().Unix())
	for i, _ := range nums {
		nums[i] = rand.Intn(100)
	}
	fmt.Printf("%v ", nums)
	fmt.Println("\n开始输出")
	for i, _ := range nums {
		fmt.Printf("%d ", nums[4-i])
	}
}

func main() {
	// array_defination()
	timu_1()
	timu_2()
	timu_5()
	two_d_array()
}
