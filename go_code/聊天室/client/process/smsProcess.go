/*
 * @Author: your name
 * @Date: 2021-12-01 21:00:49
 * @LastEditTime: 2021-12-04 21:08:59
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /go_project/src/go_code/聊天室/client/process/smsProcess.go
 */
package process

import (
	"encoding/json"
	"errors"
	"fmt"
	"go_code/聊天室/client/utils"
	"go_code/聊天室/common"
	"net"
)

type SmsProcess struct {
	Conn net.Conn
}

func (this *SmsProcess) SendMsg(content string, username string) (err error) {

	var mes common.Message
	mes.Type = common.SmsMessageType

	var smsMessage common.SmsMessage
	smsMessage.Username = username
	smsMessage.Content = content
	smsMessage.Sendtime = common.TimeFormat()

	data, err := json.Marshal(smsMessage)
	if err != nil {
		fmt.Println("序列化loginMes结构体失败:", err)
		return errors.New("序列化loginMes结构体失败")
	}
	mes.Data = string(data)

	mes_data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("序列化Message结构体失败:", err)
		return errors.New("序列化Message结构体失败")
	}
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	tf.WritePkg(mes_data)

	if err != nil {
		fmt.Println("发送结构体发送失败:", err)
		return errors.New("登陆结构体长度发送失败")
	}

	return
}

func (this *SmsProcess) SendChat(content string, from_username string, to_username string) (err error) {

	var mes common.Message
	mes.Type = common.ChatMessageType

	var chatMessage common.ChatMessage
	chatMessage.FromUsername = from_username
	chatMessage.ToUsername = to_username
	chatMessage.Content = content
	chatMessage.Sendtime = common.TimeFormat()

	data, err := json.Marshal(chatMessage)
	if err != nil {
		fmt.Println("序列化chatMessage结构体失败:", err)
		return errors.New("序列化chatMessages结构体失败")
	}
	mes.Data = string(data)

	mes_data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("序列化Message结构体失败:", err)
		return errors.New("序列化Message结构体失败")
	}
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	tf.WritePkg(mes_data)

	if err != nil {
		fmt.Println("发送结构体发送失败:", err)
		return errors.New("发送结构体发送失败")
	}

	return
}
