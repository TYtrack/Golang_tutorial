/*
 * @Author: your name
 * @Date: 2021-12-01 15:32:36
 * @LastEditTime: 2021-12-05 02:10:48
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /go_project/src/go_code/聊天室/server/process/process.go
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

type UserProcess struct {
	Conn     net.Conn
	Username string
}

func (this *UserProcess) ProcessLogin(msg *common.Message) (username string, err error) {
	login_msg := common.LoginMes{}
	err = json.Unmarshal([]byte((*msg).Data), &login_msg)
	if err != nil {
		fmt.Println("反序列化失败")
		return
	}
	this.Username = login_msg.User.Username
	username = login_msg.User.Username

	retMes := common.Message{}
	retMes.Type = common.LoginRetMesType
	loginRetMsg := common.LoginRetMes{}

	_, err = model.MyUserDao.Login(login_msg.User.Username, login_msg.User.Password)
	if err != nil {
		if err == model.ERROR_USER_PWD {
			loginRetMsg.Code = 300
			loginRetMsg.ErrorMes = "密码不对或账号不对"
			fmt.Println("密码不对或账号不对")
		} else if err == model.ERROR_USER_NOT_EXISTS {
			loginRetMsg.Code = 500
			loginRetMsg.ErrorMes = "账号未注册"
			fmt.Println("账号未注册")
		} else {
			loginRetMsg.Code = 400
			loginRetMsg.ErrorMes = "服务器内部错误"
			fmt.Println("服务器内部错误")

		}
		//500：未注册

	} else {
		loginRetMsg.Code = 200
		loginRetMsg.ErrorMes = "账号登陆成功"

		MyUserMgr.AddOnlineUser(this)

		onlineUser := MyUserMgr.GetAllUser()
		var users []string = make([]string, 0)
		for k := range onlineUser {
			users = append(users, k)
		}

		loginRetMsg.Users = users

		MyUserMgr.NotifyOtherUserOnline(this.Username)

	}

	fmt.Printf("当前在线用户：%v\n", loginRetMsg)
	data, err := json.Marshal(loginRetMsg)
	if err != nil {
		fmt.Println("序列化loginRetMes结构体失败:", err)
		return
	}
	retMes.Data = string(data)

	data_Msg, err := json.Marshal(retMes)
	if err != nil {
		fmt.Println("序列化Message结构体失败:", err)
		return
	}
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data_Msg)

	return
}

func (this *UserProcess) ProcessRegister(msg *common.Message) (err error) {

	register_unmarshal_msg := (*msg).Data
	fmt.Println("userprocess 87", register_unmarshal_msg)

	user, err := model.MyUserDao.Register(register_unmarshal_msg)

	retMes := common.Message{}
	retMes.Type = common.RegisterRetMesType

	registerRetMsg := common.RegisterRetMes{}
	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			registerRetMsg.Code = 300
			registerRetMsg.ErrorMes = "用户名已存在"
			fmt.Println("用户名已存在")
		} else {
			registerRetMsg.Code = 400
			registerRetMsg.ErrorMes = "服务器内部错误"
			fmt.Println("服务器内部错误")
		}
	} else {
		registerRetMsg.Code = 200
		registerRetMsg.ErrorMes = "账号注册成功"
		fmt.Println(user, "账号注册成功")
	}

	data, err := json.Marshal(registerRetMsg)
	if err != nil {
		fmt.Println("序列化registerRetMsg结构体失败:", err)
		return err
	}
	retMes.Data = string(data)

	data_Msg, err := json.Marshal(retMes)
	if err != nil {
		fmt.Println("序列化Message结构体失败:", err)
		return err
	}
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data_Msg)

	return err
}

func (this *UserProcess) ProcessLogout(username string) (err error) {
	MyUserMgr.DeleteOnlineUser(username)
	MyUserMgr.NotifyOtherUserOffline(username)
	return
}
