/*
 * @Author: your name
 * @Date: 2021-11-22 18:06:06
 * @LastEditTime: 2021-11-22 18:45:55
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /go_project/src/go_code/工厂模式/model/student.go
 */
package model

type Student struct {
	Name  string
	Score float64
}

// 因为是小写，所以只能在model包中使用
// 使用工厂模式来解决
type teacher struct {
	name string
	age  int
}

func NewTeacher(n string, age_ int) *teacher {
	return &teacher{
		name: n,
		age:  age_,
	}
}
