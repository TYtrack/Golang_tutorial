/*
 * @Author: your name
 * @Date: 2021-11-21 00:10:51
 * @LastEditTime: 2021-11-21 19:52:04
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /go_project/src/go_code/映射/main/map_demo.go
 */

package main

import "fmt"

func map_demo() {
	stus := make(map[int]map[string]string)
	stus[0] = make(map[string]string)
	stus[0]["name"] = "Tom"
	stus[0]["sex"] = "Male"

	stus[1] = make(map[string]string)
	stus[1]["name"] = "Alice"
	stus[1]["sex"] = "FeMale"

	stus[2] = make(map[string]string)
	stus[2]["name"] = "Kim"
	stus[2]["sex"] = "FeMale"

	fmt.Println(stus)
	// 删除
	delete(stus, 2)

	fmt.Println(stus)

}

//遍历map
func map_demo_02() {
	m1 := make(map[int]int)
	m1[1] = 22
	m1[2] = 33
	m1[3] = 44
	fmt.Printf("the length is %v \n", len(m1))

	for i, v := range m1 {
		fmt.Printf("the index is %v ,the value is %v\n", i, v)
	}

}

//map切片
func map_slice_demo() {
	var monsters = make([]map[string][]rune, 2)
	if monsters[0] == nil {
		monsters[0] = make(map[string][]rune)
		monsters[0]["name"] = []rune("皮卡丘")
		monsters[0]["attrs"] = []rune("电系")
	}

	if monsters[1] == nil {

		monsters[1] = make(map[string][]rune)
		monsters[1]["name"] = []rune("杰尼龟")
		monsters[1]["attrs"] = []rune("水系")
	}

	fmt.Println(monsters)
	// 这是错误的
	// monsters[2] = make(map[string][]rune)
	// monsters[2]["name"] = []rune("妙蛙种子")
	// monsters[2]["attrs"] = []rune("草系")
	// fmt.Println(monsters)

	// 要用append函数向切片加入，一定要返回
	temp := make(map[string][]rune)
	monsters = append(monsters, temp)
	monsters[2]["name"] = []rune("妙蛙种子")
	monsters[2]["attrs"] = []rune("草系")
	fmt.Println(monsters)

}

// 是否引用类型
func map_demo_04() {
	var monsters = make([]map[string]string, 1)
	if monsters[0] == nil {
		monsters[0] = make(map[string]string)
		monsters[0]["name"] = "皮卡丘"
		monsters[0]["attrs"] = "电系"
	}

	if monsters[1] == nil {
		monsters[1] = make(map[string]string)
		monsters[1]["name"] = "杰尼龟"
		monsters[1]["attrs"] = "水系"
	}

}

func main() {
	// map_demo()
	// map_demo_02()
	// map_slice_demo()
	map_demo_04()
}
