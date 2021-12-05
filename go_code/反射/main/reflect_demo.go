/*
 * @Author: your name
 * @Date: 2021-11-28 15:38:04
 * @LastEditTime: 2021-11-28 16:40:40
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /go_project/src/go_code/反射/main/reflect_demo.go
 */

package main

import (
	"fmt"
	"reflect"
)

//演示对reflect.Value interface{} 和变量之间的转换
// 并获取变量到type kind 值
//变量到inteface{}：参数传递
func reflect_test(b interface{}) {
	fmt.Println(b)

	//1、interface{}到reflect.Value：函数reflect.Valueof
	// rval不是真正的100，不是变量100，不能进行加减乘除
	rval := reflect.ValueOf(b)

	fmt.Println(rval)

	//2、接口、reflect.Value 到 reflect.Type
	rtype_1 := reflect.TypeOf(b)
	rtype_2 := reflect.TypeOf(rval)

	fmt.Println(rtype_1)
	fmt.Println(rtype_2)

	// 3、转化为整形
	val_temp := rval.Int()
	fmt.Printf("val_temp:%v, %T\n", val_temp, val_temp)

	//4\转化为接口，再转化为整型
	iv := rval.Interface()
	is, ok := iv.(int)
	if ok {
		fmt.Println("***", is)
	}

}

//对结构体的反射
type Student struct {
	Name string
	Age  int
}

func reflect_test_02(b interface{}) {
	fmt.Println("-----------------------")
	// 获取 reflect.Type
	rtype := reflect.TypeOf(b)
	fmt.Println(rtype)

	rval := reflect.ValueOf(b)
	fmt.Println(rval)

	iv := rval.Interface()
	fmt.Printf("iv : %v , %T\n", iv, iv)

	switch iv.(type) {
	case Student:
		stu := iv.(Student)
		fmt.Println("Student _name:", stu.Name)
	default:
		fmt.Println("不确定")
	}

}

func detail(b interface{}) {
	rtype := reflect.TypeOf(b)
	rval := reflect.ValueOf(b)
	// reflect.Value.Kind得到的是一个常量，Type的具体类型
	fmt.Printf("the Kind is %v\n", rval.Kind())
	fmt.Printf("the Kind is %v\n", rtype.Kind())
}

func reflect_modify_value(b interface{}) {
	// 会报错，因为b是一个指针
	// rval := reflect.ValueOf(b)
	// rval.SetInt(78)

	//使用Elem，得到指针指向的值的封装，即rval.Elem()可以看作为*rval
	rval := reflect.ValueOf(b)
	rval.Elem().SetInt(78)
	fmt.Println("after modify: ", rval.Elem())

}

func main() {
	i := 100
	// reflect_test(i)

	// var stu Student = Student{"Tom", 19}
	// reflect_test_02(stu)
	// detail(i)
	// detail(stu)
	reflect_modify_value(&i)
	fmt.Println("after modify: ", i)

}
