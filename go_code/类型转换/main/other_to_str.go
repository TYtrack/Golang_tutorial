/*
 * @Author: your name
 * @Date: 2021-11-09 15:59:43
 * @LastEditTime: 2021-11-18 15:40:48
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /go_project/src/go_code/类型转换/main/other_to_str.go
 */

package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println()
	//其他类型转string
	tt := fmt.Sprintf("%d", 22)
	zz := strconv.FormatInt(4378, 10)
	kk := strconv.Itoa(777)
	fmt.Println(tt, zz, kk)

	//string转其他类型，第二个参数是err
	k, ess := strconv.ParseInt("672232", 10, 64)
	pp, ess1 := strconv.Atoi("432")
	k += 1
	fmt.Println(k, ess, pp, ess1)

}
