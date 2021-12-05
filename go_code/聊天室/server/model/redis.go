/*
 * @Author: your name
 * @Date: 2021-12-02 11:14:59
 * @LastEditTime: 2021-12-02 11:24:31
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /go_project/src/go_code/聊天室/server/model/redis.go
 */
package model

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

var Pool *redis.Pool

func InitPool(address string, maxIdle int, maxActive int, idleTimeout time.Duration) {
	Pool = &redis.Pool{
		//最大空闲数
		MaxIdle: maxIdle,
		//最大连接数，0为无限制
		MaxActive: maxActive,
		// 空闲连接超时时间，超过超时时间的空闲连接会被关闭。
		IdleTimeout: idleTimeout,
		//初始化代码
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", address)
		},
	}
}
