/*
 * @Author: your name
 * @Date: 2021-11-29 18:15:48
 * @LastEditTime: 2021-11-29 18:46:39
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /go_project/src/go_code/反射/bridge_test/bridge_test.go
 */

package bridgetest_test

import (
	"fmt"
	"reflect"
	"testing"
)

//桥接模式
func TestBridge(t *testing.T) {
	f1 := func(n int, m int) {
		t.Log(n, m)
	}

	f2 := func(n int, m int, s1 string) {
		t.Log(n, m, s1)
	}

	bridge_func := func(c1 interface{}, paras ...interface{}) {
		n := len(paras)
		inValue := make([]reflect.Value, n)
		for i := 0; i < n; i++ {
			inValue[i] = reflect.ValueOf(paras[i])
		}
		function := reflect.ValueOf(c1)
		function.Call(inValue)
	}
	bridge_func(f1, 4, 7)
	bridge_func(f2, 4, 7, "hello")

}

type Student struct {
	Name string
	Age  int
}

// 更改任意结构体的字段名
func TestModifyFiled(t *testing.T) {
	var stu *Student = &Student{"Jerry", 19}
	sv := reflect.ValueOf(stu)
	sv = sv.Elem()
	sv.Field(0).SetString("Tom")
	sv.FieldByName("Age").SetInt(15)
	fmt.Println(stu)

}
