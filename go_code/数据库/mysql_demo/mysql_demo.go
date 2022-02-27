/*
 * @Author: your name
 * @Date: 2021-12-10 12:22:58
 * @LastEditTime: 2021-12-10 16:02:58
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /goproject/src/go_code/数据库/mysql_demo/mysql_demo.go
 */
package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name   string
	Age    int
	Gender string
}

func main() {
	dsn := "root:123456@(127.0.0.1:3306)/Student?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	//db, err := gorm.Open("mysql", dsn)

	if err != nil {
		fmt.Println("db open failed : ", err)
	}
	// defer db.Close()
	db.AutoMigrate(&Student{})

	// u1 := Student{
	// 	Name:   "Tom",
	// 	Age:    23,
	// 	Gender: "男",
	// }
	// db.Create(&u1)

	// u2 := Student{
	// 	Name:   "Jerry",
	// 	Age:    38,
	// 	Gender: "女",
	// }
	// db.Create(&u2)

	//查询mysql
	var stu Student
	db.First(&stu, 1)
	fmt.Println("stu1:", stu)

	var stu2 Student
	db.First(&stu2, "age = ?", 38)
	fmt.Println("stu2:", stu2)

	//更新
	db.Model(&stu).Update("Age", 89)

	//删除
	db.Delete(&stu2)

}
