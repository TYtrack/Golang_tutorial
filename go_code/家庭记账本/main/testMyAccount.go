/*
 * @Author: your name
 * @Date: 2021-11-23 19:33:07
 * @LastEditTime: 2021-11-23 19:47:39
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /go_project/src/go_code/家庭记账本/main/testMyAccount.go
 */

package main

import "fmt"

func printInfo() {
	fmt.Println("打印记账")
}
func writeInput() {
	fmt.Println("记录输入")
}

func writeOutput() {
	fmt.Println("记录输出")
}

func printUI() {
	var choice int
for_label:
	for {
		fmt.Println("欢迎进入家庭记账本")
		fmt.Println("1、收支明细")
		fmt.Println("2、登记收入")
		fmt.Println("3、登记输出")
		fmt.Println("4、退出软件")
		fmt.Printf("请选择1-4: ")
	goto_label:
		fmt.Scanf("%d", &choice)
		fmt.Println()
		switch choice {
		case 1:
			printInfo()
		case 2:
			writeInput()
		case 3:
			writeOutput()
		case 4:
			break for_label
		case 5:
			fmt.Printf("请重新选择 1-4")
			goto goto_label

		}

	}

}

func main() {
	printUI()
}
