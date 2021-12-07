/*
 * @Author: your name
 * @Date: 2021-11-21 19:27:44
 * @LastEditTime: 2021-11-22 17:35:11
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /go_project/src/go_code/结构体/main/struct_demo.go
 */
package main

import (
	"encoding/json"
	"fmt"
)

type pokemon_go struct {
	name  string
	attrs string
	hd    int
}

type pokemon_go_2 struct {
	Name  string
	Attrs string
	Hd    int
}

type pokemon_go_3 struct {
	Name  string `json:"name"`
	Attrs string `json:"attrs"`
	Hd    int    `json:"hd"`
}

//序列化一个结构体
func reflect_struct_tag() {
	// 可以序列化pokemon_go_2，但是不能序列化pokemon_go，因为pokemon_go所有字段都是小写，外部json 不可以访问
	pikaqiu := pokemon_go_2{"pikaqiu", "dian22222", 3}

	s1, _ := json.Marshal(pikaqiu)

	fmt.Println(s1)
	// 需要转化为string类型
	fmt.Println(string(s1))

	// 前端代码可能不支持首字母大写的字段
	// 因此可以通过加上tag来进行反射来实现json中的字段变tag所属内容
	pi := pokemon_go_3{"jienigui", "shui22222", 8}
	s2, _ := json.Marshal(pi)
	fmt.Println(string(s2))

}

//结构体声明
func struct_use() {
	// 第一种
	var num1 pokemon_go
	num1.attrs = "电系"
	num1.name = "皮卡丘"
	num1.hd = 5
	fmt.Println(num1)

	//第二种 初始化

	num2 := pokemon_go{"水系", "杰尼龟", 3}
	fmt.Println(num2)
	//最后一个也需要逗号
	num5 := pokemon_go{
		attrs: "火系",
		name:  "小火龙",
		hd:    8,
	}
	fmt.Println(num5)

	//第三种 new
	num3 := new(pokemon_go)
	(*num3).attrs = "木系"
	num3.hd = 9
	(*num3).name = "妙蛙种子"
	fmt.Println(*num3)

	//第四种
	num4 := &pokemon_go{}
	num4.attrs = "火系"
	(*num4).hd = 91
	num4.name = "小火龙"
	fmt.Println(*num4)

}

type animal struct {
	name  string
	attrs string
	hd    int
}

type plants struct {
	name  string
	attrs string
	hds   int
}

func type_transfer() {
	num2 := pokemon_go{"水系", "杰尼龟", 3}
	fmt.Println(num2)

	// 结构体是用户单独定义的类型，和其它类型进行转换时需要有完全相同的字段(名字、个数和类 型)
	// 可以通过
	var num3 animal = animal(num2)
	fmt.Println(num3)

	//不可通过，因为字段名不一样
	// var num4 plants = plants(num2)
	// fmt.Println(num4)

}

func main() {
	// reflect_struct_tag()
	struct_use()
	// type_transfer()
}
