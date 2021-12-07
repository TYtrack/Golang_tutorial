/*
 * @Author: your name
 * @Date: 2021-11-30 21:16:33
 * @LastEditTime: 2021-12-04 20:11:13
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /go_project/src/go_code/聊天室/common/message.go
 */

package common

const (
	SmsMessageType  = "SmsMessage"
	ChatMessageType = "ChatMessage"

	LoginMesType    = "LoginMes"
	LoginRetMesType = "LoginRetMes"

	RegisterMesType    = "RegisterMes"
	RegisterRetMesType = "RegisterRetMes"

	ShowOnlineUserMesType = "ShowOnlineUserMes"
	OneUserStatusMesType  = "OneUserStatusMes"
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

type LoginMes struct {
	User User `json:"user"`
}

type LoginRetMes struct {
	Code     int      `json:"type"` //500：未注册  200：登陆成功  300：密码不对或账号不对
	Users    []string `json:"users"`
	ErrorMes string   `json:"errorMes"`
}

type RegisterMes struct {
	User User `json:"user"`
}

type RegisterRetMes struct {
	Code     int    `json:"type"` //300：用户名重复  200：注册成功
	ErrorMes string `json:"errorMes"`
}

type OneUserStatusMes struct {
	Username     string `json:"username"`
	OnlineStatus int    `json:"onlinestatus"` // 1表示连接
}

// 群聊消息类型
type SmsMessage struct {
	Username string `json:"username"`
	Content  string `json:"content"`
	Sendtime string `json:"sendtime"`
}

//私聊消息类型
type ChatMessage struct {
	FromUsername string `json:"from_username"`
	ToUsername   string `json:"to_username"`
	Content      string `json:"content"`
	Sendtime     string `json:"sendtime"`
}
