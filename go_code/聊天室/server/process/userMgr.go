/*
 * @Author: your name
 * @Date: 2021-12-02 23:36:58
 * @LastEditTime: 2021-12-05 01:36:41
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /聊天室/server/model/userMgr.go
 */
package process

import (
	"encoding/json"
	"fmt"
	"go_code/聊天室/common"
	"go_code/聊天室/server/model"
	"go_code/聊天室/server/utils"
)

// 加上一个知识图谱
// 在服务器只有一个，全局变量
var MyUserMgr *UserMgr

type UserMgr struct {
	onlineUsers map[string]*UserProcess
}

func init() {
	MyUserMgr = &UserMgr{
		onlineUsers: make(map[string]*UserProcess, 10),
	}
}

func (this *UserMgr) AddOnlineUser(up *UserProcess) {
	this.onlineUsers[up.Username] = up
}

func (this *UserMgr) DeleteOnlineUser(username string) {
	delete(this.onlineUsers, username)
}

func (this *UserMgr) GetAllUser() map[string]*UserProcess {
	return this.onlineUsers
}

func (this *UserMgr) GetUserByUsername(username string) (up *UserProcess, err error) {
	up, ok := this.onlineUsers[username]
	if !ok {
		err = model.ERROR_USER_OFFLINE
	}
	return
}

func (this *UserMgr) NotifyOneUserOnline(bytes []byte, up *UserProcess) (err error) {

	tf := &utils.Transfer{
		Conn: up.Conn,
	}

	err = tf.WritePkg(bytes)
	if err != nil {
		fmt.Println("Message发送失败")
		return
	}
	return
}

func (this *UserMgr) PrepareNotifyOnlineBytes(username string) (bytes []byte, err error) {
	userStatus := common.OneUserStatusMes{
		Username:     username,
		OnlineStatus: common.UserOnlineStatus,
	}
	bytes, err = json.Marshal(userStatus)
	if err != nil {
		fmt.Println("OneUserStatusMes序列化失败")
		return
	}

	msg := common.Message{}
	msg.Type = common.OneUserStatusMesType
	msg.Data = string(bytes)
	bytes, err = json.Marshal(msg)
	if err != nil {
		fmt.Println("Message序列化失败")
		return
	}
	return bytes, err
}

func (this *UserMgr) NotifyOtherUserOnline(username string) (err error) {
	userMap := this.onlineUsers
	bytes, err := this.PrepareNotifyOnlineBytes(username)

	if err != nil {
		fmt.Println("Message序列化失败")
		return
	}
	for otherUser, up := range userMap {
		if otherUser == username {

			continue
		}
		this.NotifyOneUserOnline(bytes, up)
	}
	return
}

//
func (this *UserMgr) NotifyOneUserOffline(bytes []byte, up *UserProcess) (err error) {

	tf := &utils.Transfer{
		Conn: up.Conn,
	}

	err = tf.WritePkg(bytes)
	if err != nil {
		fmt.Println("Message发送失败")
		return
	}
	return
}

func (this *UserMgr) PrepareNotifyOfflineBytes(username string) (bytes []byte, err error) {
	userStatus := common.OneUserStatusMes{
		Username:     username,
		OnlineStatus: common.UserOfflineStatus,
	}
	fmt.Println("UserOfflineStatus:", common.UserOfflineStatus)
	bytes, err = json.Marshal(userStatus)
	if err != nil {
		fmt.Println("OneUserStatusMes序列化失败")
		return
	}

	msg := common.Message{}
	msg.Type = common.OneUserStatusMesType
	msg.Data = string(bytes)
	bytes, err = json.Marshal(msg)
	if err != nil {
		fmt.Println("Message序列化失败")
		return
	}
	return bytes, err
}

func (this *UserMgr) NotifyOtherUserOffline(username string) (err error) {
	userMap := this.onlineUsers
	bytes, err := this.PrepareNotifyOfflineBytes(username)

	if err != nil {
		fmt.Println("Message序列化失败")
		return
	}
	for otherUser, up := range userMap {
		if otherUser == username {
			continue
		}

		this.NotifyOneUserOffline(bytes, up)
	}
	return
}
