/*
 * @Author: your name
 * @Date: 2021-12-07 12:56:29
 * @LastEditTime: 2021-12-07 13:00:42
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /goproject/src/go_code/反射/Type_API/type_api.go
 */

package main

import (
	"fmt"
	"reflect"
)

func add(a, int, b int, c int) (d int, num int) {
	return a + b + c, 3
}

//关于函数的Type的API
func function_Type_Api() {
	tp := reflect.TypeOf(add)
	fmt.Println("函数的参数数量:", tp.NumIn())
	fmt.Println("函数的返回值数量:", tp.NumOut())

	fmt.Println("函数的第二个参数的Type:", tp.In(1))
	fmt.Println("函数的第二个返回值数量:", tp.Out(1))

}

func main() {
	function_Type_Api()
}
