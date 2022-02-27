/*
 * @Author: your name
 * @Date: 2021-12-08 00:35:45
 * @LastEditTime: 2021-12-08 21:54:48
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /goproject/src/go_code/MapReduce/Map/map_struct.go
 */
package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Map_t struct {
}

// 单位是M
func Partitioning(inputFile string, size int, ch chan []byte, partition func()) {

	file, err := os.Open(inputFile)

	finfo, err := file.Stat()

	if err != nil {
		fmt.Println("get file info failed:", file, size)
	}

	fmt.Println(finfo, size)

	//每次最多拷贝1m
	bufsize := 1024 * 1024 * 64
	if size < bufsize {
		bufsize = size
	}
	buf := make([]byte, bufsize)
	num := (int(finfo.Size()) + size - 1) / size
	fmt.Println(num, len(buf))
	for i := 0; i < num; i++ {
		copylen := 0
		newfilename := finfo.Name() + strconv.Itoa(i)
		newfile, err1 := os.Create(newfilename)
		if err1 != nil {
			fmt.Println("failed to create file", newfilename)
		} else {
			fmt.Println("create file:", newfilename)
		}
		for copylen < size {
			n, err2 := file.Read(buf)
			if err2 != nil && err2 != io.EOF {
				fmt.Println(err2, "failed to read from:", file)
				break
			}
			if n <= 0 {
				break
			}
			//fmt.Println(n, len(buf))
			//写文件
			w_buf := buf[:n]
			newfile.Write(w_buf)
			copylen += n
		}
	}
	return
}

func countWord(word string) (res int) {
	dir := "./Life/"
	count := 0

	namefileInfo, err := ioutil.ReadDir(dir)

	if err != nil {
		fmt.Println(err)
		return
	}
	for _, info := range namefileInfo {
		filetemp, _ := os.OpenFile(dir+info.Name(), os.O_RDONLY, 0777)
		reader := bufio.NewReader(filetemp)
		defer filetemp.Close()

		for {
			zz, err := reader.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					break
				} else {
					fmt.Println("uiiu", err)
				}
			}

			if len(zz) == 0 {
				continue
			}
			akk := strings.Split(zz, " ")
			for _, str := range akk {

				str = strings.Replace(str, "\n", "", -1)
				str = strings.Replace(str, ":", "", -1)
				str = strings.Replace(str, ",", "", -1)
				str = strings.Replace(str, ".", "", -1)
				str = strings.Replace(str, "'", "", -1)
				str = strings.Replace(str, "\"", "", -1)
				str = strings.Replace(str, "(", "", -1)
				str = strings.Replace(str, ")", "", -1)
				str = strings.Replace(str, "!", "", -1)
				str = strings.Replace(str, "?", "", -1)

				str = strings.ToLower(str)
				if str == word {
					count++
				}
			}
		}
	}
	return count
}

func map2(inputFile string) (err error) {
	defer func() {
		err := recover() //recover能够捕获异常
		if err != nil {
			fmt.Println("err is ", err)
		}
	}()

	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Errorf("zzz%v\n", err)
		return
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	res := make(map[string]int)
	for {
		str, err2 := reader.ReadString('\n')
		if err2 != nil {
			if err2 == io.EOF {
				break
			}
			fmt.Errorf("%v\n", err2)
			return err2
		}

		akk := strings.Split(str, " ")
		for _, str = range akk {

			str = strings.Replace(str, "\n", "", -1)
			str = strings.Replace(str, ":", "", -1)
			str = strings.Replace(str, ",", "", -1)
			str = strings.Replace(str, ".", "", -1)
			str = strings.Replace(str, "'", "", -1)
			str = strings.Replace(str, "\"", "", -1)
			str = strings.Replace(str, "(", "", -1)
			str = strings.Replace(str, ")", "", -1)
			str = strings.Replace(str, "!", "", -1)
			str = strings.Replace(str, "?", "", -1)

			str = strings.ToLower(str)
			res[str]++
		}

	}
	keys := make([]string, 0)
	for index, _ := range res {
		keys = append(keys, index)
	}
	sort.Strings(keys)

	l1 := strings.LastIndex(inputFile, "/")

	output_file := "Res/" + inputFile[l1+1:]
	fmt.Println(output_file)
	file2, err := os.OpenFile(output_file, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0777)

	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	defer file2.Close()
	fmt.Println("yuy22u")

	writer := bufio.NewWriter(file2)

	indexLetter := make(map[rune]int)
	before := ' '
	nums_byte := 0

	for _, k := range keys {

		stt := fmt.Sprintf("%v : %v\n", k, res[k])

		if len(k) > 0 && rune(k[0]) != before {
			indexLetter[rune(k[0])] = nums_byte
			before = rune(k[0])
		}

		num_byte, _ := writer.WriteString(stt)
		nums_byte += num_byte

	}
	writer.Flush()

	index_filename := "Index/" + inputFile[l1+1:]

	indexFile, err := os.OpenFile(index_filename, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Println("tt", err)
		return
	}
	defer indexFile.Close()

	writer2 := bufio.NewWriter(indexFile)

	for k, v := range indexLetter {
		zz := fmt.Sprintf("%c : %v\n", k, v)
		writer2.WriteString(zz)
	}
	writer2.Flush()

	fmt.Println(indexLetter)

	return

}

func reduce_2(startChar rune) (err error) {
	dir := "./Index/"
	dir2 := "./Res/"
	namefileInfo, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
		return
	}
	prefixMap := make(map[string]int)
	for _, info := range namefileInfo {
		// fmt.Println(info.Name())
		filetemp, _ := os.OpenFile(dir+info.Name(), os.O_RDONLY, 0777)
		reader := bufio.NewReader(filetemp)
		defer filetemp.Close()
		for {
			zz, err := reader.ReadString('\n')
			if err == io.EOF {
				break
			}
			var cc rune
			var startIndex int
			fmt.Sscanf(zz, "%c : %v\n", &cc, &startIndex)
			if cc != startChar {
				continue
			} else {
				fmt.Println("find successful")
				filetemp2, _ := os.OpenFile(dir2+info.Name(), os.O_RDONLY, 0777)
				defer filetemp2.Close()
				filetemp2.Seek(int64(startIndex), io.SeekStart)
				reader2 := bufio.NewReader(filetemp2)
				for {

					temp_str, err := reader2.ReadString('\n')
					if err == io.EOF {
						break
					}
					var word string
					var freq int
					fmt.Sscanf(temp_str, "%v : %v\n", &word, &freq)
					if len(temp_str) < 0 || temp_str[0] != byte(startChar) {
						break
					}
					prefixMap[word] += freq
				}

			}
			//reader.

		}
	}
	fmt.Println(prefixMap)
	return
}

func main() {
	dir := "./Life"

	namefileInfo, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, info := range namefileInfo {
		fmt.Println(info.Name())
		map2("./Life/" + info.Name())
	}
	map2("zz.txt")
	//Partitioning("golang.pdf", 1024*1024*16, nil, nil)

	reduce_2('y')
	fmt.Println("youth", countWord("youth"))
	fmt.Println("yelling", countWord("yelling"))
}
