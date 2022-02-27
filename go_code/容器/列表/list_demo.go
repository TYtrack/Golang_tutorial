/*
 * @Author: TYtrack
 * @Description: ...
 * @Date: 2022-02-09 21:36:12
 * @LastEditors: TYtrack
 * @LastEditTime: 2022-02-09 21:44:29
 * @FilePath: /go_code/容器/列表/list_demo.go
 */

package main

import (
	"container/list"
	"fmt"
)

// Golang的列表是通过双向链表实现的，进行愿随的插入喝删除需要同时修改前后元素的持有引用

//另外每次插入都会返回一个list.Element结构，用以指向当前插入值所在的节点，删除、移动以及指定插入都需要使用list.Element
func main() {
	tmpList := list.New()

	for i := 0; i < 10; i++ {
		tmpList.PushBack(i)
	}

	tt := tmpList.PushFront(-1)

	for l := tmpList.Front(); l != nil; l = l.Next() {
		fmt.Print(" ", l.Value)
	}

	fmt.Println()

	tmpList.Remove(tt)

	for l := tmpList.Front(); l != nil; l = l.Next() {
		fmt.Print(" ", l.Value)
	}

}
