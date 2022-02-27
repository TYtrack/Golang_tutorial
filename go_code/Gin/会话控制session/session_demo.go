/*
 * @Author: your name
 * @Date: 2021-12-12 13:24:36
 * @LastEditTime: 2021-12-12 13:37:00
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /goproject/src/go_code/Gin/会话控制session/session_demo.go
 */

package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

// 初始化一个cookie存储对象
// something-very-secret应该是一个你自己的密匙，只要不被别人知道就行
var store = sessions.NewCookieStore([]byte("something-very-secret"))

func main() {
	http.HandleFunc("/save", SaveSession)
	http.HandleFunc("/get", GetSession)
	http.HandleFunc("/delete", DeleteSession)
	err := http.ListenAndServe(":8989", nil)
	if err != nil {
		fmt.Println("HTTP server failed,err:", err)
		return
	}
}

// 保存session
func SaveSession(w http.ResponseWriter, r *http.Request) {
	// Get a session. We're ignoring the error resulted from decoding an
	// existing session: Get() always returns a session, even if empty.

	//　获取一个session对象，session-name是session的名字
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 在session中存储值
	session.Values["foo"] = "bar"
	session.Values[42] = 43
	// 保存更改
	session.Save(r, w)
}

func GetSession(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	foo := session.Values["foo"]
	fmt.Println(foo)

	foo = session.Values[42]
	fmt.Println(foo)
}

func DeleteSession(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")
	// 删除
	// 将session的最大存储时间设置为小于零的数即为删除
	session.Options.MaxAge = -1
	session.Save(r, w)
}
