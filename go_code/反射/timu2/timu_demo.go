/*
 * @Author: your name
 * @Date: 2021-11-29 18:51:53
 * @LastEditTime: 2021-11-29 19:09:30
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /go_project/src/go_code/反射/timu2/timu_demo.go
 */
package main

import (
	"fmt"
	"reflect"
)

type Cal struct {
	Num1 int `json:"第一个数"`
	Num2 int `json:"第二个数"`
}

func (cal Cal) GetSub(name string) {
	res := fmt.Sprintf("%v 完成了减法，结果是%v", name, cal.Num1-cal.Num2)
	fmt.Println(res)
}

func TestCal(a interface{}) {
	typ := reflect.TypeOf(a)
	val := reflect.ValueOf(a)

	num := val.NumField()
	for i := 0; i < num; i++ {
		tag := typ.Field(i).Tag.Get("json")
		zz := typ.Field(i).Name
		fmt.Printf("第%v个字段名字是%v,其值是%v, tag是%v\n", i, zz, val.Field(i), tag)
	}

	call := val.Method(0)

	arr := make([]reflect.Value, 1)
	arr[0] = reflect.ValueOf("Jerry")
	call.Call(arr)

}

func main() {
	a := Cal{8, 3}
	TestCal(a)
}
