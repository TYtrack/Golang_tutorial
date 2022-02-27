/*
 * @Author: your name
 * @Date: 2021-12-09 21:41:00
 * @LastEditTime: 2021-12-09 22:01:21
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /goproject/src/go_code/web开发之模版/条件判断与range循环/if_and_range.go
 */
package main

import (
	"fmt"
	"html/template"
	"net/http"
)

//模版中使用了range 和if
func tmplDemo(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./t.tmpl")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	names := []string{"Tom", "Jerry", "Sally"}

	tmpl.Execute(w, names)
}

func main() {
	http.HandleFunc("/tyty", tmplDemo)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
