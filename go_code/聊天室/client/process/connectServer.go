/*
 * @Author: your name
 * @Date: 2021-12-01 21:10:31
 * @LastEditTime: 2021-12-04 21:32:53
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /go_project/src/go_code/聊天室/client/process/connectServer.go
 */
package process

import (
	"bufio"
	"encoding/json"
	"fmt"
	"go_code/聊天室/client/model"
	"go_code/聊天室/client/utils"
	"go_code/聊天室/common"
	"net"
	"os"
)

func ShowLoginSuccessMenu(username string, conn net.Conn) {
	fmt.Printf("----------------------%v成功登陆MyRoom聊天室----------------------\n", username)
restart:
	fmt.Println("                          1、显示在线用户列表")
	fmt.Println("                          2、进入群聊")
	fmt.Println("                          3、进入私聊")
	fmt.Println("                          4、消息列表")
	fmt.Println("                          5、退出系统")
	fmt.Println("                         请选择（1-4）：")
	var choose int
	fmt.Scanf("%d\n", &choose)
	switch choose {
	case 1:
		fmt.Println("                          显示在线用户列表")
		model.MyUserOnlineMap.PrintUserOnline()

	case 2:
		fmt.Println("                          请输入想发送的内容")
		reader := bufio.NewReader(os.Stdin)
		bytes, _, _ := reader.ReadLine()

		smsProcess := &SmsProcess{
			Conn: conn,
		}
		smsProcess.SendMsg(string(bytes), username)
	case 3:
		fmt.Println("                           输入消息接受方")
		reader := bufio.NewReader(os.Stdin)
		bytes, _, _ := reader.ReadLine()
		to_username := string(bytes)

		fmt.Println("                          请输入想发送的内容")
		bytes, _, _ = reader.ReadLine()
		smsProcess := &SmsProcess{
			Conn: conn,
		}
		smsProcess.SendChat(string(bytes), username, to_username)

	case 5:
		fmt.Println("                          成功退出系统")
		os.Exit(0)
	default:
		fmt.Println("                      输入错误，请重新选择选项")
		goto restart

	}

}

func serverProcessConn(conn net.Conn) (err error) {
	tf := &utils.Transfer{
		Conn: conn,
	}

	for {

		msg, err := tf.ReadPkg()
		if err != nil {
			return err
		}
		switch msg.Type {
		case common.OneUserStatusMesType:
			userState := common.OneUserStatusMes{}
			err = json.Unmarshal([]byte(msg.Data), &userState)
			if err != nil {
				fmt.Println("序列化失败")
				continue
			}
			model.MyUserOnlineMap.UpdateUserOnline(userState.Username, userState.OnlineStatus)

		case common.SmsMessageType:
			smsMessage := common.SmsMessage{}
			err = json.Unmarshal([]byte(msg.Data), &smsMessage)
			if err != nil {
				fmt.Println("反序列化失败")
			}
			utils.MsgWirteToLog(smsMessage)
			fmt.Printf("%v: %v的群聊消息：%v\n", smsMessage.Sendtime, smsMessage.Username, smsMessage.Content)

		case common.ChatMessageType:
			chatMessage := common.ChatMessage{}
			err = json.Unmarshal([]byte(msg.Data), &chatMessage)
			if err != nil {
				fmt.Println("反序列化失败")
			}
			utils.ChatWirteToLog(chatMessage)
			fmt.Printf("%v: %v->%v的群聊消息：%v\n", chatMessage.Sendtime, chatMessage.FromUsername, chatMessage.ToUsername, chatMessage.Content)

		default:
			fmt.Println("**处理其他东西", msg)
		}

	}
}
