/*
 * @Author: your name
 * @Date: 2021-11-23 15:22:18
 * @LastEditTime: 2021-11-23 15:54:02
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /go_project/src/go_code/继承/main/inherit.go
 */
package main

import "fmt"

// 1、继承是嵌套一个匿名结构，如果不是匿名就不是继承，是组合
// 2、继承可以使用嵌套的匿名结构体的所有字段，包括小写
// 3、当结构体和匿名结构体有一样的字段，采取就近原则（本身字段为默认）
// 4、当两个嵌入的匿名结构体有相同字段，则必须知名匿名结构体的名称，不然报错
// 5、如果是组合，就必须加上结构体的名称
type student struct {
	name  string
	grade float64
}

func (s student) printInfo() {
	fmt.Printf("name : %v \t grade:%v \n ", s.name, s.grade)
}

type pupil struct {
	// 匿名
	student
	money float64
}

type graduate struct {
	//匿名
	student
	balance float64
}

type high_student struct {
	//不是匿名，是组合，所以必须加上匿名结构体的变量名
	s1    student
	class string
}

//

func main() {
	pupil_1 := pupil{}
	pupil_1.money = 44
	pupil_1.name = "zzz"
	pupil_1.student.grade = 78.2
	fmt.Println(pupil_1)
	pupil_1.printInfo()

	graduate_1 := graduate{}
	graduate_1.balance = 324
	graduate_1.student.name = "kty"
	graduate_1.grade = 21.5
	fmt.Println(graduate_1)
	graduate_1.printInfo()

	high_1 := high_student{}
	high_1.s1.name = "Tome"
	high_1.class = "C412"
	// 下面会报错，因为high_student是组合不是匿名结构体，所以必须加上匿名结构体的变量名
	high_1.s1.grade = 88
	fmt.Println(high_1)
	high_1.s1.printInfo()

}
