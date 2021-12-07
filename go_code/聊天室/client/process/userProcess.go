/*
 * @Author: your name
 * @Date: 2021-12-01 21:00:58
 * @LastEditTime: 2021-12-04 19:51:28
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /go_project/src/go_code/聊天室/client/process/userProcess.go
 */

package process

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"go_code/聊天室/client/model"
	"go_code/聊天室/client/utils"
	"go_code/聊天室/common"
	"net"
)

type UserProcess struct {
	Conn net.Conn
}

func (this *UserProcess) Login() (err error) {
	username, password := this.LoginMenu()

	// 连接服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8989")
	this.Conn = conn
	if err != nil {
		fmt.Println("服务器连接错误:", err)
		return errors.New("服务器连接错误")
	}

	// defer conn.Close()

	// 发送消息
	var mes common.Message
	mes.Type = common.LoginMesType

	var loginMes common.LoginMes
	loginMes.User.Username = username
	loginMes.User.Password = password

	data, err := json.Marshal(loginMes)
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
	// loginMes是要发送的东西
	// 需要先发loginMes 的长度([]byte类型)，防止粘包、拆包，
	// 需要使用encoding/binary包来将int转化为byte[]

	tf := &utils.Transfer{
		Conn: conn,
	}

	err = tf.WritePkg(mes_data)
	if err != nil {
		return
	}

	msg, err := tf.ReadPkg()
	if err != nil {
		return
	}
	//反序列化msg.Data
	var loginRetMsg common.LoginRetMes
	err = json.Unmarshal([]byte(msg.Data), &loginRetMsg)
	if err != nil {
		return
	}

	if loginRetMsg.Code == 200 {
		//在这里开启一个协程，保持和服务器通信

		for _, userN := range loginRetMsg.Users {
			model.MyUserOnlineMap.UpdateUserOnline(userN, common.UserOnlineStatus)
		}
		go serverProcessConn(conn)

		for {
			ShowLoginSuccessMenu(username, conn)
		}
	}

	return nil
}

func (this *UserProcess) LoginMenu() (username string, password string) {
	fmt.Println("----------------------欢迎登陆MyRoom聊天室----------------------")
	fmt.Println("                         请输入账号")
	var user, pwd string
	fmt.Scanln(&user)
	fmt.Println("                         请输入密码")
	fmt.Scanln(&pwd)

	return user, pwd
}

func (this *UserProcess) Register() (err error) {
	username, password := this.RegisterMenu()

	// 连接服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8989")
	this.Conn = conn
	if err != nil {
		fmt.Println("服务器连接错误:", err)
		return errors.New("服务器连接错误")
	}

	// defer conn.Close()

	// 发送消息
	var mes common.Message
	mes.Type = common.RegisterMesType

	var registerMes common.RegisterMes
	registerMes.User.Username = username
	registerMes.User.Password = password

	data, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("序列化registerMes结构体失败:", err)
		return errors.New("序列化registerMes结构体失败")
	}
	mes.Data = string(data)

	mes_data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("序列化Message结构体失败:", err)
		return errors.New("序列化Message结构体失败")
	}

	var lenPkg uint32 = uint32(len(mes_data))
	var bytes [4]byte
	binary.BigEndian.PutUint32(bytes[:], lenPkg)

	n, err := conn.Write(bytes[:])
	if n != 4 || err != nil {
		fmt.Println("结构体长度发送失败:", err)
		return errors.New("结构体长度发送失败")
	}

	// 发送数据包
	n, err = conn.Write(mes_data)
	if err != nil {
		fmt.Println("注册结构体发送失败:", err)
		return errors.New("注册结构体长度发送失败")
	}

	tf := &utils.Transfer{
		Conn: conn,
	}
	msg, err := tf.ReadPkg()
	if err != nil {
		return
	}
	//反序列化msg.Data
	var registerRetMes common.RegisterRetMes
	err = json.Unmarshal([]byte(msg.Data), &registerRetMes)
	if err != nil {
		return
	}
	if registerRetMes.Code == 200 {
		//在这里开启一个协程，保持和服务器通信
		fmt.Println("**注册成功")
	} else if registerRetMes.Code == 300 {
		fmt.Println("**注册失败，用户名存在")
		err = errors.New("注册失败，用户名存在")
	} else {
		fmt.Println("**注册失败，服务器内部存在错误")
		err = errors.New("注册失败，服务器内部存在错误")
	}

	return
}

func (this *UserProcess) RegisterMenu() (username string, password string) {
	fmt.Println("----------------------欢迎注册MyRoom聊天室----------------------")
	fmt.Println("                       请输入新账号用户名")
	var user, pwd string
	fmt.Scanln(&user)
	fmt.Println("                        请输入新账户密码")
	fmt.Scanln(&pwd)

	return user, pwd
}
