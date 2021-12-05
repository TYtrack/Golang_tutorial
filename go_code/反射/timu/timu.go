/*
 * @Author: your name
 * @Date: 2021-11-28 16:46:15
 * @LastEditTime: 2021-11-28 19:26:29
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /go_project/src/go_code/反射/timu/timu.go
 */

package main

import (
	"fmt"
	"reflect"
)

func timu_01(b interface{}) {
	rtype := reflect.TypeOf(b)
	rvalue := reflect.ValueOf(b)
	rkind := rvalue.Kind()

	fmt.Printf("%v   ,   %v   ,   %v\n", rtype, rvalue, rkind)

	inter2 := rvalue.Interface()
	value, ok := inter2.(float64)
	if ok {
		fmt.Println("zzz  ", value)
	}

}

//遍历结构体的字段，调用结构体的方法，以及获取标签的值
func timu

func main() {
	var z float64 = 67.2
	timu_01(z)
}
