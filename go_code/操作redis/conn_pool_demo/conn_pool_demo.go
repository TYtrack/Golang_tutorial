/*
 * @Author: your name
 * @Date: 2021-11-30 14:59:29
 * @LastEditTime: 2021-12-02 23:19:29
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /go_project/src/go_code/操作redis/conn_pool_demo/conn_pool_demo.go
 */

package main

import (
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
)

var pool *redis.Pool

func init() {
	pool = &redis.Pool{
		//最大空闲数
		MaxIdle: 8,
		//最大连接数，0为无限制
		MaxActive: 0,
		// 空闲连接超时时间，超过超时时间的空闲连接会被关闭。
		IdleTimeout: 5,
		//初始化代码
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
	}

}
func conn_demo(va string) {

	//获取一个连接，关闭连接池
	fmt.Printf("%v:   %v,%v,%v\n ", va, pool.IdleCount(), pool.ActiveCount(), pool.IdleTimeout)
}

func main() {

	fmt.Println("初始化完成")
	conn_demo("0")
	conn := pool.Get()
	z2 := pool.Get()
	t2 := pool.Get()
	k2 := pool.Get()
	z3 := pool.Get()
	t3 := pool.Get()

	conn_demo("1")
	z2.Close()
	k2.Close()
	t2.Close()
	z3.Close()
	t3.Close()

	filed1, val1 := "Name", "Tom"
	_, err := conn.Do("Set", filed1, val1)
	if err != nil {
		fmt.Println("执行错误1:", err)
	}
	conn_demo("2")

	conn.Close()
	time.Sleep(time.Second * 1)
	conn_demo("3")
	z := pool.Get()
	t := pool.Get()
	k := pool.Get()
	conn_demo("4")
	// time.Sleep(time.Second * 8)
	conn_demo("5")
	z.Close()
	t.Close()
	k.Close()
	conn_demo("6")

	pool.Close()

}
