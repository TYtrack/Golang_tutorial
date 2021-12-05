/*
 * @Author: your name
 * @Date: 2021-12-01 15:33:06
 * @LastEditTime: 2021-12-05 02:20:29
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /go_project/src/go_code/聊天室/server/process/smsProcess.go
 */

package process

import (
	"encoding/json"
	"fmt"
	"go_code/聊天室/common"
	"go_code/聊天室/server/model"
	"go_code/聊天室/server/utils"
	"net"
)

type SmsProcess struct {
	Conn     net.Conn
	Username string
}

func (this *SmsProcess) ReceiveMsg(msg *common.Message) (smsMessage common.SmsMessage) {
	smsMessage = common.SmsMessage{}
	bytes := []byte(msg.Data)
	json.Unmarshal(bytes, &smsMessage)
	fmt.Printf("%v: %v发送了%v\n", smsMessage.Sendtime, smsMessage.Username, smsMessage.Content)
	return

}

func (this *SmsProcess) ForwardMsg(msg *common.Message, from_username string) {

	bytes, err := json.Marshal(*msg)
	if err != nil {
		fmt.Println("序列化失败")
	}
	user_map, err := model.MyUserDao.GetAllUser()

	for _, username := range user_map {
		if username == from_username {
			continue
		}
		up, err := MyUserMgr.GetUserByUsername(username)
		if err == model.ERROR_USER_OFFLINE {
			fmt.Printf("用户%v不在线\n", username)
			model.MyUserDao.PutOfflineChat(username, string(bytes))
			continue
		}
		ut := &utils.Transfer{
			Conn: up.Conn,
		}
		err = ut.WritePkg(bytes)
		if err != nil {
			fmt.Printf("给%v发送消息失败\n", username)
		} else {
			fmt.Printf("给%v发送消息成功\n", username)
		}
	}

}

func (this *SmsProcess) ReceiveChat(msg *common.Message) (chatMessage common.ChatMessage) {
	chatMessage = common.ChatMessage{}
	bytes := []byte(msg.Data)
	json.Unmarshal(bytes, &chatMessage)
	fmt.Printf("%v: %v向%v发送了%v\n", chatMessage.Sendtime, chatMessage.FromUsername, chatMessage.ToUsername, chatMessage.Content)
	return
}

//转发私聊
func (this *SmsProcess) ForwardChat(msg *common.Message, to_username string) (err error) {
	bytes, err := json.Marshal(*msg)
	if err != nil {
		fmt.Println("序列化失败")
	}

	_, err = model.MyUserDao.GetUserByUserName(to_username)
	if err != nil {
		if err == model.ERROR_USER_NOT_EXISTS {
			fmt.Println("用户不存在")
			return

		}
	}

	up, err := MyUserMgr.GetUserByUsername(to_username)
	if err != nil {
		if err == model.ERROR_USER_OFFLINE {
			fmt.Println("用户已经下线")
			model.MyUserDao.PutOfflineChat(to_username, string(bytes))
			return

		}
	}

	ut := &utils.Transfer{
		Conn: up.Conn,
	}

	err = ut.WritePkg(bytes)
	if err != nil {
		fmt.Printf("向%v转发失败\n", to_username)

	} else {
		fmt.Printf("向%v转发成功\n", to_username)

	}
	return
}

func (this *SmsProcess) SendOfflineMsg(to_username string) (err error) {
	msgs, err := model.MyUserDao.GetOfflineMsg(to_username)
	if err != nil {
		fmt.Printf("给%v发送离线消息失败\n", to_username)
		return
	}

	up, err := MyUserMgr.GetUserByUsername(to_username)
	ut := &utils.Transfer{
		Conn: up.Conn,
	}

	for _, msg := range msgs {
		ut.WritePkg([]byte(msg))
	}
	fmt.Printf("给%v发送离线消息成功\n", to_username)
	model.MyUserDao.RemOfflineMsg(to_username)
	return
}

func (this *SmsProcess) SendOfflineChat(to_username string) (err error) {
	msgs, err := model.MyUserDao.GetOfflineChat(to_username)
	if err != nil {
		fmt.Printf("给%v发送离线私聊失败:%v\n", to_username, err)
		return
	}

	up, err := MyUserMgr.GetUserByUsername(to_username)
	ut := &utils.Transfer{
		Conn: up.Conn,
	}

	for _, msg := range msgs {
		ut.WritePkg([]byte(msg))
	}
	fmt.Printf("给%v发送私聊消息成功\n", to_username)
	model.MyUserDao.RemOfflineChat(to_username)
	return
}
