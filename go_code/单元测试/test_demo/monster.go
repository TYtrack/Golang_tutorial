/*
 * @Author: your name
 * @Date: 2021-11-25 13:27:51
 * @LastEditTime: 2021-11-25 14:03:01
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /go_project/src/go_code/单元测试/test_demo/monster.go
 */

package test_demo

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Monster struct {
	Name  string
	Age   int
	Skill string
}

func (m *Monster) String() string {
	res := fmt.Sprintf("结构体：name is %v,age is %v,skill is %v", m.Name, m.Age, m.Skill)
	return res
}

func (m *Monster) Store() error {
	bytes, err := json.Marshal(m)
	if err == nil {
		fmt.Println("序列化结果:   ", string(bytes))
	}
	file, err := os.OpenFile("monstre.json", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(string(bytes) + "\n")
	writer.Flush()
	if err != nil {
		return err
	}
	return nil
}

func (m *Monster) ReStore() error {

	bytes, err := ioutil.ReadFile("monstre.json")
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, m)
	if err != nil {
		return err
	}
	fmt.Println(m)
	return nil

}
