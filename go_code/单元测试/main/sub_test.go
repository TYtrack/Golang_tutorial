/*
 * @Author: your name
 * @Date: 2021-11-25 13:09:00
 * @LastEditTime: 2021-11-25 14:21:35
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /go_project/src/go_code/单元测试/test_case/sub_test.go
 */
package main

import "testing"

// go test 在正确时不会输出Logf，但是如果加上-v会有Logf

// 测试用例文件必须以_test.go结尾，一个测试文件可以有很多函数
// go test -v 能够自动执行测试用例文件下的TestXxx的函数，第一个X必须是大写
// 行参必须是 *testing.T

// 测试单个文件：命令参数需要加上测试用例文件名以及待测试的文件  go test -v cal_test.go cal.go
// 测试单个方法：go test -v -run=TestAdd
func TestSub(t *testing.T) {
	res := sub(9, 3)
	if res != 6 {
		t.Fatalf("sub执行错误，结果应该是10，而不是%v", res)

	}
	t.Logf("sub执行正确")
}
