/*
 * @Author: your name
 * @Date: 2021-11-22 14:18:38
 * @LastEditTime: 2021-11-22 14:27:42
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /go_project/src/go_code/面向对象/box/mian.go
 */

package main

import (
	"fmt"
	"strconv"
)

type box struct {
	height int
	length int
	weight int
}

func (b *box) getInfo() {
	fmt.Println("ok")
	fmt.Scanf("%v %v %v", &(b.height), &(b.length), &(b.weight))
}

func (b *box) printInfo() {
	fmt.Printf("Box info : %d %d %d\n", (*b).height, (*b).length, (*b).weight)
}

func (b *box) getTi() {
	temp := b.height * b.length * b.weight
	fmt.Println("tiji is ", strconv.FormatInt(int64(temp), 10))
}

func main() {
	var bb box
	bb.getInfo()
	bb.printInfo()
	bb.getTi()

}
