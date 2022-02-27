/*
 * @Author: your name
 * @Date: 2021-12-10 10:06:00
 * @LastEditTime: 2021-12-10 12:22:39
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /goproject/src/go_code/Gin/中间件/middle_func.go
 */

package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// ！！！
// 当在中间件或handler中启动新的goroutine时，
// 不能使用原始的上下文（c *gin.Context），
// 必须使用其只读副本（c.Copy()）。

//定义一个中间件
func login(c *gin.Context) {
	fmt.Println("进入一个登陆中间件")
	c.JSON(http.StatusOK, gin.H{
		"msg": "login middle layer",
	})
	// 在中间件设置值
	c.Set("login", "32")
	c.Next()

}

//定义一个计时中间件
func countTime(c *gin.Context) {
	fmt.Println("进入一个计时中间件")
	startTime := time.Now()
	c.Next() //调用后续的处理函数
	//c.Abort() //阻止调用后续的处理函数 之前的中间件以及本中间件函数会继续完成，即会打印“处理时间”，但是不会进入到其他函数了
	cost := time.Since(startTime)
	fmt.Println("处理时间：", cost)

}

//以闭包的方式去写登陆认证的，可以在闭包中进行连接数据库，或者一些其他操作
func authMiddleware(doCheck bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		if doCheck {
			//if 登陆成功
			// c.Next()
			// else
			// c.Abort()
		} else {
			//
		}
	}
}

func main() {
	// 1.创建路由
	r := gin.Default() //默认包含有Logger(), Recovery()两个中间件
	//r:= gin.New()

	//！！！全局设置中间件函数
	r.Use(countTime)

	// 2.绑定 !!!!这里设置了中间件
	r.GET("/index", login, func(c *gin.Context) {
		// 从中间件获取值
		va, _ := c.Get("login")
		fmt.Println("zzzz   ", va)
		c.JSON(http.StatusOK, gin.H{"msg": "/index"})
	})

	r.GET("/ty", login, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "/index"})
	})

	r.GET("/tyty", authMiddleware(true), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "/index"})
	})

	//为路由组注册中间件，也可以是用Use，下面两种方法二选一
	//第一种
	homeGroup := r.Group("/home", authMiddleware(true))
	//第二种
	// homeGroup := r.Group("/home")
	// homeGroup.Use(authMiddleware(true))
	{
		homeGroup.GET("/info", login)
	}
	r.Run(":9999")
}
