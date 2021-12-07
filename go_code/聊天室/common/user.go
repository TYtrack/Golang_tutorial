/*
 * @Author: your name
 * @Date: 2021-12-02 01:22:08
 * @LastEditTime: 2021-12-03 14:06:05
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /go_project/src/go_code/聊天室/server/model/user.go
 */

package common

const (
	UserOnlineStatus = iota
	UserOfflineStatus
)

type User struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	Userid       int    `json:"userid"`
	OnlineStatus int    `json:"onlinestatus"`
}
