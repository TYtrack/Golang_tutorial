/*
 * @Author: your name
 * @Date: 2021-12-01 15:33:30
 * @LastEditTime: 2021-12-05 02:10:24
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /go_project/src/go_code/聊天室/server/main/processor.go
 */
package main

import (
	"fmt"
	"go_code/聊天室/common"
	"go_code/聊天室/server/process"
	"go_code/聊天室/server/utils"
	"io"
	"net"
)

type Processor struct {
	Conn     net.Conn
	Username string
}

// 根据服务器接受的消息种类不同，选择调用不同的函数
func (this *Processor) serveProcessMsg(msg *common.Message) (err error) {
	switch (*msg).Type {
	case common.LoginMesType:
		userProcess := &(process.UserProcess{
			Conn: this.Conn,
		})
		username, err := userProcess.ProcessLogin(msg)
		this.Username = username

		smsProcess := &(process.SmsProcess{
			Conn: this.Conn,
		})
		smsProcess.SendOfflineMsg(username)
		smsProcess.SendOfflineChat(username)

		return err

	case common.RegisterMesType:
		userProcess := &(process.UserProcess{
			Conn: this.Conn,
		})
		return userProcess.ProcessRegister(msg)
	case common.SmsMessageType:
		smsProcess := &(process.SmsProcess{
			Conn: this.Conn,
		})
		smsMessage := smsProcess.ReceiveMsg(msg)
		smsProcess.ForwardMsg(msg, smsMessage.Username)
	case common.ChatMessageType:
		chatProcess := &(process.SmsProcess{
			Conn: this.Conn,
		})
		chatMessag := chatProcess.ReceiveChat(msg)
		chatProcess.ForwardChat(msg, chatMessag.ToUsername)

	default:
		fmt.Println("默认消息")
	}
	return

}

func (this *Processor) process_2() (err error) {
	for {
		tf := &utils.Transfer{
			Conn: this.Conn,
		}
		msg, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Printf("%v客户端关闭连接:%v", this.Username, err)
				//开始注销登陆
				userProcess := &(process.UserProcess{
					Conn: this.Conn,
				})
				userProcess.ProcessLogout(this.Username)
				return err
			} else {
				fmt.Println("服务器接受登陆结构体发送失败:", err)
			}
		}
		err = this.serveProcessMsg(&msg)
		if err != nil {
			return err
		}

		fmt.Printf("结构体内容为%v\n", msg)
	}
	return
}
