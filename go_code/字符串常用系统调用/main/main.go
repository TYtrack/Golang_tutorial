/*
 * @Author: your name
 * @Date: 2021-11-18 15:58:39
 * @LastEditTime: 2021-11-18 16:12:08
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /go_project/src/go_code/字符串常用系统调用/main/main.go
 */

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	// 1、len（str）是返回字符串长度，【不支持中文】
	// 2、遍历中文需要使用[]rune(str)进行强制转化
	str_1 := "你真好"
	runes := []rune(str_1)
	fmt.Println(len(str_1), "  ", len(runes))

	// 3、进制转化，返回的是string
	zz := strconv.FormatInt(4378, 9)
	fmt.Println(zz)

	// 4、数字与字符串的转换
	kk := strconv.Itoa(777)
	pp, _ := strconv.Atoi("432")
	fmt.Println(kk, "   ", pp)

	// 5、字符串与byte[]的转换
	var bytes = []byte("nmsl")
	fmt.Println(bytes)
	str_4 := string([]byte{67, 56, 72})
	fmt.Println(str_4)

	//是否包含子串、子串数量、字符串比较（不区分大小写）、索引、最后索引、替换、切割、大小写转换、
	fmt.Println(strings.Contains("shjda", "da"))
	fmt.Println(strings.Count("sdahjda", "da"))
	fmt.Println(strings.EqualFold("shjda", "ShjDA"))
	fmt.Println(strings.Index("zzdashjda", "da"))
	fmt.Println(strings.LastIndex("zzdashjda", "da"))
	fmt.Println(strings.Replace("zzdashjda", "da", "88", -1))
	fmt.Println(strings.ToLower("zzdaFHJSa"))
	fmt.Println(strings.ToUpper("zzdaFHJSa"))

	// 去除两边/左边/右边的空格或指定字符
	fmt.Println(strings.TrimSpace("   zzdaFHJSa  "))
	fmt.Println(strings.Trim("!! zzdaFHJSa !! 9  ", "!"))
	fmt.Println(strings.TrimLeft("!! zzdaFHJSa !! 9  ", "!"))
	fmt.Println(strings.TrimRight("!! zzdaFHJSa !! 9  ", "!"))

	// 	判断是否有前缀后缀
	fmt.Println(strings.HasPrefix("tytytyakwjdbe", "tytyty"))
	fmt.Println(strings.HasSuffix("tytytyakwjdbe", "be"))

}
