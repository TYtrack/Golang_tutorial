/*
 * @Author: your name
 * @Date: 2021-11-21 16:51:20
 * @LastEditTime: 2021-11-21 19:14:18
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /go_project/src/go_code/排序/main/sort_demo.go
 */
package main

import (
	"fmt"
	"sort"
)

func sort_int() {
	a := [10]int{2, 7, 1, 5, 74, 32, 865, 0, -3, 22}
	b := a[:]
	fmt.Println(b)
	sort.Ints(b)
	fmt.Println(b)
}

//map是没有进行排序的，解决办法是把键放在切片中，然后把键排序，然后再遍历
func sort_map() {
	a := make(map[int]int, 10)
	a[1] = 22
	a[2] = 433
	a[3] = 32
	a[4] = 28
	keys := make([]int, 0)
	for index, _ := range a {
		keys = append(keys, index)
	}
	sort.Ints(keys)
	for _, v := range keys {
		fmt.Printf("%v ", a[v])
	}

}

func main() {
	sort_int()
	sort_map()
}
