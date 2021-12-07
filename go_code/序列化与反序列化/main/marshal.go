/*
 * @Author: your name
 * @Date: 2021-11-24 20:40:15
 * @LastEditTime: 2021-11-24 21:20:44
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /go_project/src/go_code/序列化与反序列化/main/marshal.go
 */
package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Grade float64
}

func marshal_struct() {
	//如果结构体的字段是小写，那么不会出现来该字段，因为json.Marshal不能访问
	stu := Student{"Tom", 17, 68.5}
	bytes, err := json.Marshal(stu)
	if err == nil {
		fmt.Println(string(bytes))
	}

	//反序列化
	k := Student{}
	err = json.Unmarshal(bytes, &k)
	if err == nil {
		fmt.Println(k)
	}
}

func marshal_map() {
	var a map[string]interface{} = make(map[string]interface{})
	a["name"] = "红孩儿"
	a["age"] = 8
	a["father"] = "牛魔王"
	a["mother"] = "铁扇公主"
	bytes, err := json.Marshal(a)
	if err == nil {
		fmt.Println(string(bytes))
	}

	// 反序列化不需要make,make的操作封装到了Unmarshal中
	var b map[string]interface{}
	err = json.Unmarshal(bytes, &b)
	if err == nil {
		fmt.Println(b)
	}
}

func marshal_slice() {
	var a []map[string]interface{} = make([]map[string]interface{}, 2)
	a[0] = make(map[string]interface{})
	a[1] = make(map[string]interface{})
	a[0]["name"] = "红孩儿"
	a[0]["age"] = 8
	a[0]["father"] = "牛魔王"
	a[0]["mother"] = "铁扇公主"

	a[1]["name"] = "哪吒"
	a[1]["age"] = 9
	a[1]["father"] = "哪吒他爹"
	a[1]["mother"] = "哪吒他妈"
	a[1]["brother"] = "哪吒他两个哥哥"

	bytes, err := json.Marshal(a)
	if err == nil {
		fmt.Println(string(bytes))
	}

	//反序列化 slice 不需要make
	var b []map[string]interface{}
	err = json.Unmarshal(bytes, &b)
	if err == nil {
		fmt.Println(b)
	}
}

func main() {
	marshal_struct()
	marshal_map()
	marshal_slice()
}
