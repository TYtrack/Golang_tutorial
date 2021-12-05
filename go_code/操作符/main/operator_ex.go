/*
 * @Author: your name
 * @Date: 2021-11-10 19:01:38
 * @LastEditTime: 2021-11-10 19:10:11
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /go_project/src/go_code/操作符/main/operator_ex.go
 */
package main

import (
	"fmt"
	"strconv"
)

func main() {
	//没有++i --i
	//只有i++,i--
	i := 5
	i++
	// 且i++只能作为一句话，不能作为表达式

	fmt.Println(i)
	// 下面注释是错误的 ，因为I++只能作为一句话
	//str_1 := strconv.FormatInt(int64(i++), 10)
	str_1 := strconv.FormatInt(int64(i), 10)
	fmt.Println(str_1)

	//go语言不支持三目操作

}
