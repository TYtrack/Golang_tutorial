/*
 * @Author: your name
 * @Date: 2021-11-22 18:06:54
 * @LastEditTime: 2021-11-22 18:56:15
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /go_project/src/go_code/工厂模式/main/factory_demo.go
 */
package main

import (
	"fmt"
	"go_code/factory_method/model"
)

func main() {
	// 如果 model 包的 结构体变量首字母大写，引入后，直接使用, 没有问题
	stu := model.Student{"Tom", 87.4}
	fmt.Println(stu)

	// 如果 model 包的 结构体变量首字母小写，引入后，不能直接使用, 可以工厂模式解决
	tea := model.NewTeacher("Jerry", 26)
	fmt.Println(*tea)

}
