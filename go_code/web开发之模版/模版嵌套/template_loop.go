/*
 * @Author: your name
 * @Date: 2021-12-09 21:37:27
 * @LastEditTime: 2021-12-09 21:38:57
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /goproject/src/go_code/web开发之模版/模版嵌套/template_loop.go
 */
package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type Stu struct {
	Name string
	//这里是小写，最后template渲染不出来
	gender string
	Age    int
}

//模版嵌套，在解析模板时，被嵌套的模板一定要在后面解析
func tmplDemo(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./t.tmpl", "./ul.tmpl")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	user := Stu{
		Name:   "小王子",
		gender: "男",
		Age:    18,
	}
	tmpl.Execute(w, user)
}

func main() {
	http.HandleFunc("/", tmplDemo)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
