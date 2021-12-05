/*
 * @Author: your name
 * @Date: 2021-11-18 17:23:14
 * @LastEditTime: 2021-12-04 20:00:20
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /go_project/src/go_code/时间日期/main/date.go
 */
package main

import (
	"fmt"
	"strconv"
	"time"
)

func time_consumer(k func()) {
	now := time.Now()
	k()
	dur := time.Now().Sub(now)
	fmt.Println(dur.String())

}

func pp() {
	res := ""
	for i := 0; i < 100000; i++ {

		res += strconv.Itoa(i)

	}
}

//时间格式化
func timeFormat() {
	logTime := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println("时间是", logTime)
}

func main() {
	timeFormat()
	time_consumer(pp)

}
