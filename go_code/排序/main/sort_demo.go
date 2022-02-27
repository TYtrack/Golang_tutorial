/*
 * @Author: your name
 * @Date: 2021-11-21 16:51:20
 * @LastEditTime: 2021-12-07 13:43:42
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /go_project/src/go_code/排序/main/sort_demo.go
 */
package main

import (
	"fmt"
	"sort"
)

//递增、递减
func sort_int() {
	a := [10]int{2, 7, 1, 5, 74, 32, 865, 0, -3, 22}
	b := a[:]

	//递增
	fmt.Println(b)
	fmt.Println(sort.IntsAreSorted(b))
	sort.Ints(b)
	fmt.Println(b)
	fmt.Println(sort.IntsAreSorted(b))

	//递减
	greater := func(i, j int) bool {
		return b[i] > b[j]
	}
	sort.Slice(b, greater)
	fmt.Println(b)

	//按字符串长度
	s2 := []string{"666666", "1", "333", "7777777", "4444", "22", "55555"}
	sort.Slice(s2, func(i, j int) bool {
		return len(s2[i]) < len(s2[j])
	})
	fmt.Println(s2)
}

// 对自定义类型实现排序，需要实现sort.Interface接口，具体有三个方法
type IntNums []int

//需要实现三个方法，分别是Swap、Less、Len
func (x IntNums) Len() int {
	return len(x)
}
func (x IntNums) Less(i, j int) bool {
	return x[i] < x[j]
}

func (x IntNums) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}
func sort_struct() {
	var a IntNums = IntNums{5, 2, 7, 1, 4, 9, 3, 10, -1}

	sort.Sort(a)
	fmt.Println(a)

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
	fmt.Println()

}

func main() {
	sort_int()
	sort_map()
	sort_struct()
}
