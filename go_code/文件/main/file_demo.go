/*
 * @Author: your name
 * @Date: 2021-11-23 20:59:38
 * @LastEditTime: 2021-11-24 17:27:30
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /go_project/src/go_code/文件/main/file_demo.go
 */
package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func file_open() {
	file, err := os.Open("myfile.txt")
	if err != nil {
		fmt.Printf("文件打开有错误,err is %v", err)

	} else {
		file.Close()
		fmt.Println("打开成功")
	}

}

//带缓冲区的读取
func file_read() {
	file, err := os.Open("myfile.txt")
	if err != nil {
		fmt.Printf("文件打开有错误,err is %v", err)
		return
	}
	fmt.Println("打开成功")
	//否则会有内存泄漏
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		//从缓冲区读字符串
		z, err := reader.ReadString('\n')
		if err == nil {
			fmt.Print(z)
		} else if err == io.EOF {
			fmt.Println("文件读取完毕")
			break
		} else {
			break
		}
	}
}

//一次性读取文件，适合文件不太大的情况，因为没有open，所以不需要close
func read_file() {
	content, err := ioutil.ReadFile("myfile.txt")
	if err != nil {
		fmt.Printf("文件打开有错误,err is %v", err)
		return
	}
	fmt.Println(string(content))
}

// 带缓冲的写
func create_file(filename string) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)

	if err != nil {
		fmt.Printf("文件打开有错误,err is %v", err)
		return
	} else {
		writer := bufio.NewWriter(file)
		for i := 0; i < 5; i++ {
			writer.WriteString("hello\n")
		}
		// 因为这是带缓冲的，所以要写入磁盘中要flush，将缓存的数据写入到磁盘中
		// 否则文件中会丢失数据
		writer.Flush()
	}
	defer file.Close()
}

//读取file1_to_file第一种方法
func file1_to_file2_demo1() {
	file1, err1 := os.OpenFile("file1.txt", os.O_RDONLY, 0777)
	if err1 != nil {
		return
	}
	defer file1.Close()
	file2, err2 := os.OpenFile("file2.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777)
	if err2 != nil {
		return
	}
	defer file2.Close()
	reader := bufio.NewReader(file1)
	writer := bufio.NewWriter(file2)
	for {
		s1, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		writer.WriteString(s1)

	}
	writer.Flush()

}

//读取file1_to_file第二种方法
func file1_to_file2_demo2() {

	bytes_1, _ := ioutil.ReadFile("file1.txt")
	ioutil.WriteFile("file2.txt", bytes_1, 0777)

}

//判断文件是否存在
func isFile(filename string) {
	_, err := os.Stat(filename)
	if err == nil {
		fmt.Println("文件存在")
		return
	}
	if os.IsNotExist(err) {
		fmt.Println("文件不存在")
	} else {
		fmt.Println("判断失败")
	}
}

// 由io包实现的
func CopyFile(des_filename string, src_filename string) (int, error) {
	des_file, err_1 := os.OpenFile(des_filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)

	if err_1 != nil {
		return -1, err_1
	}
	defer des_file.Close()

	src_file, err_2 := os.OpenFile(src_filename, os.O_RDONLY, 0666)
	if err_2 != nil {
		return -1, err_2
	}
	defer src_file.Close()

	reader := bufio.NewReader(src_file)

	writer := bufio.NewWriter(des_file)

	nums, err_3 := io.Copy(writer, reader)
	if err_3 != nil {
		return int(nums), err_3
	}
	return int(nums), nil

}

func main() {
	// file_open()
	// file_read()
	// read_file()
	// create_file("hello.txt")
	// file1_to_file2_demo1()
	// file1_to_file2_demo2()
	// isFile("nmsl")
	// isFile("file1.txt")
	CopyFile("zy.png", "/Users/zplus/Desktop/cug_logo.png")
}
