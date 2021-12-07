/*
 * @Author: your name
 * @Date: 2021-11-25 13:34:04
 * @LastEditTime: 2021-11-25 14:02:34
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /go_project/src/go_code/单元测试/test_demo/monster_test.go
 */

package test_demo

import "testing"

func TestMarshal(t *testing.T) {
	var m *Monster = &Monster{"牛魔王", 19, "喷火"}
	err := m.Store()
	if err != nil {
		t.Fatalf("Store出现错误,%v", err)
	}
}

func TestUnmarshal(t *testing.T) {
	var b Monster
	err := b.ReStore()
	if err != nil {
		t.Fatalf("ReStore出现错误,%v", err)
	}

}
