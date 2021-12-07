/*
 * @Author: your name
 * @Date: 2021-12-03 11:29:59
 * @LastEditTime: 2021-12-04 19:43:08
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /聊天室/client/model/userOnlineList.go
 */

package model

import (
	"fmt"
	"go_code/聊天室/common"
)

var MyUserOnlineMap UserOnlineMap

type UserOnlineMap struct {
	UserMap map[string]int
}

func init() {
	MyUserOnlineMap.UserMap = make(map[string]int)
}

func (this *UserOnlineMap) UpdateUserOnline(username string, userStatus int) {
	if userStatus == common.UserOnlineStatus {
		this.UserMap[username] = userStatus
		fmt.Printf("用户上线： %v \n", username)
	} else if userStatus == common.UserOfflineStatus {
		delete(this.UserMap, username)
		fmt.Printf("用户离线： %v \n", username)
	}
}

func (this *UserOnlineMap) PrintUserOnline() {
	fmt.Printf("当前用户有：")
	for username, _ := range MyUserOnlineMap.UserMap {
		fmt.Printf("%v ,", username)
	}
	fmt.Println()
}
