/*
 * @Author: your name
 * @Date: 2021-12-04 20:01:48
 * @LastEditTime: 2021-12-04 20:01:48
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /聊天室/common/formatTime.go
 */
package common

import (
	"time"
)

//时间格式化
func TimeFormat() (formatTime string) {
	logTime := time.Now().Format("2006-01-02 15:04:05")
	return logTime
}
