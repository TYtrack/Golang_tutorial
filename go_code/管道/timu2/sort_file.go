/*
 * @Author: your name
 * @Date: 2021-11-26 10:58:52
 * @LastEditTime: 2021-11-26 11:42:51
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /go_project/src/go_code/管道/timu2/sort_file.go
 */

/*
题目要求：开启10个进程，各打开一个文件，随机生成1000个数，
并写入文件中，写完之后，开启另一个协程来完成排序，将排序结果重新写入另一个文件中
*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func wirteFile(filename string, fileChan chan<- string) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("文件打开错误")
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	rand.Seed(time.Now().Unix())
	for i := 1; i <= 1000; i++ {
		temp := rand.Intn(10000)
		writer.WriteString(strconv.Itoa(temp) + " \n")
	}
	writer.Flush()
	fileChan <- filename
}

func ReadFile(fileChan <-chan string, exitChan chan<- bool) {
	filename := <-fileChan
	file, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println("文件打开错误")
	}
	defer file.Close()

	nums := make([]int, 1000)
	reader := bufio.NewReader(file)
	index := 0
	for {
		numS, err := reader.ReadString('\n')

		if err == io.EOF {
			break
		}
		numS = strings.Trim(numS, " \n")
		num, _ := strconv.Atoi(numS)
		nums[index] = num
		index++
	}
	fmt.Println("nmsl", nums)
	for i := 999; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}

	file2, err := os.OpenFile("sort_"+filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("文件打开错误")
	}

	defer file2.Close()
	writer2 := bufio.NewWriter(file2)
	for i := 0; i < 1000; i++ {
		writer2.WriteString(strconv.Itoa(nums[i]) + "\n")
	}
	writer2.Flush()
	exitChan <- true
}

func main() {
	fileChan := make(chan string, 10)
	exitChan := make(chan bool, 10)
	for i := 0; i < 10; i++ {
		filename := "file_" + strconv.Itoa(i) + ".txt"
		go wirteFile(filename, fileChan)
		go ReadFile(fileChan, exitChan)
	}
	for i := 0; i < 10; i++ {
		<-exitChan
	}
	fmt.Println("over")

}
