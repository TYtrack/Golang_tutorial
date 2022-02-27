/*
 * @Author: your name
 * @Date: 2021-12-09 10:26:43
 * @LastEditTime: 2021-12-09 23:23:05
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /goproject/src/go_code/web开发之模版/template_demo/t_demo.go
 */

package main

import (
	"fmt"
	"html/template"
	"net/http"
)

//普通字符串
func Hello(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("hello.tmpl")
	if err != nil {
		return
	}
	tmpl.Execute(w, "zzz")
}

//传递结构体，可以使用同样的方法传递map(这时候key不需要首字母大写)
type Stu struct {
	Name string
	//这里是小写，最后template渲染不出来
	gender string
	Age    int
}

func t_Struct(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("t_struct.tmpl")
	if err != nil {
		return
	}
	stu := Stu{
		"Tom", "男", 78,
	}
	tmpl.Execute(w, stu)

}

//自定义函数，将execute的输入进行转化再进行渲染，
// 在Parse之前调用Funcs添加自定义的kua函数
func myself_define(w http.ResponseWriter, r *http.Request) {
	k := func(info string) (res string, err error) {
		return info + " good ", err
	}

	tmpl := template.New("define_self_function.tmpl").Funcs(template.FuncMap{"kua": k})
	tmpl, _ = tmpl.ParseFiles("define_self_function.tmpl")

	user := Stu{
		Name:   "小王子",
		gender: "男",
		Age:    18,
	}
	tmpl.Execute(w, user.Name)
}

func main() {
	http.HandleFunc("/", Hello)
	http.HandleFunc("/a", t_Struct)
	http.HandleFunc("/b", myself_define)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
