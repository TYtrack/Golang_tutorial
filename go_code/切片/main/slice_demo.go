/*
 * @Author: your name
 * @Date: 2021-11-20 13:38:43
 * @LastEditTime: 2021-11-20 23:45:36
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /go_project/src/go_code/切片/main/slice_demo.go
 */

package main

import "fmt"

func slice_demo_01() {
	var a = [...]int{1, 2, 3, 4, 5}
	var b = a[1:3]
	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("切片的长度是%v\n", len(b))
	fmt.Printf("切片的容量是%v\n", cap(b))
	b[1] = 44
	fmt.Println(a)
	fmt.Println(b)

}

func slice_defination() {
	var a = [...]int{1, 2, 3, 4, 5}
	var b = a[1:3]
	fmt.Println(b)

	//	直接创建，底层数组外部不可见
	var c = make([]int, 5, 11)
	fmt.Println(c)

	// 使用原理和make类似
	var d []int = []int{1, 2, 5, 8, 77}
	fmt.Println(d)
}

func slice_append() {
	fmt.Println("数组的append")
	var a = [...]int{1, 2, 3, 4, 5}
	var b = a[1:3]
	fmt.Println(a)
	fmt.Println(b)
	b = append(b, 2, 5, 7, 4)
	fmt.Println(a)
	fmt.Println(b)
}

func timu_1(n int) []int {
	i := 2
	temp := make([]int, n+1)
	temp[0] = 1
	temp[1] = 1
	for {
		temp[i] = temp[i-1] + temp[i-2]
		if i >= n {
			break
		}
		i++
	}
	fmt.Println(temp)
	return temp

}

func main() {
	slice_demo_01()
	slice_defination()
	slice_append()
	timu_1(10)
}
