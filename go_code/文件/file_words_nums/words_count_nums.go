/*
 * @Author: your name
 * @Date: 2021-11-24 19:28:45
 * @LastEditTime: 2021-11-24 19:37:49
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /go_project/src/go_code/文件/file_words_nums/words_count_nums.go
 */

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 统计一个文件中，数字空格和字母的数量
func count(filename string) (int, int, int) {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println("file open fail")
		return -1, -1, -1
	}
	defer file.Close()
	zimu, shu, kong, ee := 0, 0, 0, 0

	reader := bufio.NewReader(file)
	for {
		s1, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("file read fail")
			return -1, -1, -1
		}
		for _, v := range s1 {
			switch {
			case v <= 'z' && v >= 'a', v <= 'Z' && v >= 'A':
				zimu++
			case v <= '9' && v >= '0':
				shu++
			case v == ' ':
				kong++
			default:
				ee++

			}

		}
	}
	fmt.Printf("字母、数字、空格各%v,%v,%v个\n", zimu, shu, kong)
	return zimu, shu, kong

}

func main() {
	count("akb.txt")
}
