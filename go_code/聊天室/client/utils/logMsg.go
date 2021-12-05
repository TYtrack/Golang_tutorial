/*
 * @Author: your name
 * @Date: 2021-12-04 19:55:55
 * @LastEditTime: 2021-12-04 21:21:09
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /聊天室/client/utils/logMsg.go
 */

package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"go_code/聊天室/common"
	"os"
)

func MsgWirteToLog(sms common.SmsMessage) {
	log_path := ""
	smsHisName := log_path + "sms_" + sms.Username + "_log.his"

	file, err := os.OpenFile(smsHisName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)

	if err != nil {
		fmt.Printf("文件打开有错误,err is %v", err)
		return
	}
	defer file.Close()

	bytes, err := json.Marshal(sms)
	if err != nil {
		fmt.Printf("SMS序列话错误,err is %v", err)
		return
	} else {
		writer := bufio.NewWriter(file)
		writer.Write(bytes)
		writer.WriteString("\n")

		// 因为这是带缓冲的，所以要写入磁盘中要flush，将缓存的数据写入到磁盘中
		// 否则文件中会丢失数据
		writer.Flush()
	}

}

func ChatWirteToLog(sms common.ChatMessage) {
	log_path := ""
	chatHisName := log_path + "chat_" + sms.ToUsername + "_log.his"

	file, err := os.OpenFile(chatHisName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)

	if err != nil {
		fmt.Printf("文件打开有错误,err is %v", err)
		return
	}
	defer file.Close()

	bytes, err := json.Marshal(sms)
	if err != nil {
		fmt.Printf("SMS序列话错误,err is %v", err)
		return
	} else {
		writer := bufio.NewWriter(file)
		writer.Write(bytes)
		writer.WriteString("\n")

		// 因为这是带缓冲的，所以要写入磁盘中要flush，将缓存的数据写入到磁盘中
		// 否则文件中会丢失数据
		writer.Flush()
	}

}
