/*
 * @Author: your name
 * @Date: 2021-11-25 21:49:38
 * @LastEditTime: 2021-11-25 21:51:22
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /go_project/src/go_code/设置CPU/main/cpu_set.go
 */
package main

import (
	"fmt"
	"runtime"
)

func main() {
	//NumCPU返回本地机器的逻辑CPU个数。
	nums := runtime.NumCPU()

	// CPU的数量是8
	fmt.Printf("CPU的数量是%v\n", nums)

	//设置本程序可以执行的最多CPU数
	runtime.GOMAXPROCS(nums - 1)

}
