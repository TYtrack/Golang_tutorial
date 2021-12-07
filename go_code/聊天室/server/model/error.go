/*
 * @Author: your name
 * @Date: 2021-12-02 01:24:38
 * @LastEditTime: 2021-12-05 01:36:25
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /go_project/src/go_code/聊天室/server/model/errors.go
 */

package model

import "errors"

var (
	ERROR_USER_NOT_EXISTS = errors.New("用户不存在")
	ERROR_USER_EXISTS     = errors.New("用户已存在")
	ERROR_USER_PWD        = errors.New("用户密码错误")
	ERROR_USER_OFFLINE    = errors.New("用户已下线")

	ERROR_REDIS  = errors.New("Redis错误")
	ERROR_SERVER = errors.New("服务器错误")
)
