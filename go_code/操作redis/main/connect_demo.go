/*
 * @Author: your name
 * @Date: 2021-11-30 13:22:43
 * @LastEditTime: 2021-11-30 13:49:19
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /go_project/src/go_code/操作redis/main/connect_demo.go
 */

package main

import (
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
)

func operator_expire() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("连接错误:", err)
		return
	}
	defer conn.Close()

	filed1, val1 := "Name", "Tom"
	_, err = conn.Do("Setex", filed1, 4, val1)
	// 或者先set 再使用_, err = conn.Do("expire", filed1, 4)

	if err != nil {
		fmt.Println("执行错误:", err)
	}

	res, err := redis.String(conn.Do("Get", filed1))
	if err != nil {
		fmt.Println("执行错误1:", err)
	}
	fmt.Printf("the res is %v ,the type is %T \n", res, res)
	time.Sleep(time.Second * 5)

	res, err = redis.String(conn.Do("Get", filed1))
	if err != nil {
		fmt.Println("执行错误2:", err)
	}
	fmt.Printf("the res is %v ,the type is %T \n", res, res)

}

func operator_string() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("连接错误:", err)
		return
	}
	defer conn.Close()

	filed1, val1 := "Name", "Tom"
	_, err = conn.Do("Set", filed1, val1)
	if err != nil {
		fmt.Println("执行错误1:", err)
	}

	res, err := redis.String(conn.Do("Get", filed1))
	if err != nil {
		fmt.Println("执行错误2:", err)
	}
	fmt.Printf("the res is %v ,the type is %T \n", res, res)
}

func operator_hash() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("连接错误:", err)
		return
	}
	defer conn.Close()

	key, field1, field2 := "stu", "name", "age"

	_, err = conn.Do("hset", key, field1, "Tom", field2, 19)
	if err != nil {
		fmt.Println("载入错误:", err)
		return
	}

	res_1, err := redis.String(conn.Do("hget", key, field1))
	if err != nil {
		fmt.Println("载入错误:", err)
		return
	}
	fmt.Printf("the res_1 is %v ,the type is %T \n", res_1, res_1)

	res_2, err := redis.Int(conn.Do("hget", key, field2))
	if err != nil {
		fmt.Println("载入错误:", err)
		return
	}
	fmt.Printf("the res_2 is %v ,the type is %T \n", res_2, res_2)

	//注意这里是Strings，上面都是String
	res_3, err := redis.Strings(conn.Do("hmget", key, field1, field2))
	if err != nil {
		fmt.Println("载入错误:", err)
		return
	}
	for i := 0; i < len(res_3); i++ {
		fmt.Printf("the res_3 is %v ,the type is %T \n", res_3[i], res_3[i])
	}

}

func main() {
	// operator_string()

	// operator_hash()
	operator_expire()
}
