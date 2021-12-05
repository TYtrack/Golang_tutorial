/*
 * @Author: your name
 * @Date: 2021-12-01 21:09:18
 * @LastEditTime: 2021-12-04 19:33:35
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /go_project/src/go_code/聊天室/client/utils/utils.go
 */
package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go_code/聊天室/common"
	"net"
)

type Transfer struct {
	Conn net.Conn
	Buf  [8096]byte
}

func (this *Transfer) ReadPkg() (msg common.Message, err error) {

	n, err := this.Conn.Read(this.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("服务器接受登陆结构体长度失败:", err)
		return msg, err
	}

	pkgLen := binary.BigEndian.Uint32(this.Buf[:4])

	n, err = this.Conn.Read(this.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		fmt.Println("服务器接受登陆结构体发送失败:", err)
		return msg, err
	}

	err = json.Unmarshal(this.Buf[:pkgLen], &msg)
	if err != nil {
		fmt.Println("结构体序列化发送失败:", err)
		return msg, err
	}

	return msg, nil
}

func (this *Transfer) WritePkg(data []byte) (err error) {

	var pkgLen uint32 = uint32(len(data))

	binary.BigEndian.PutUint32(this.Buf[:4], pkgLen)

	n, err := this.Conn.Write(this.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("RetMsg结构体长度发送失败:", err)
		return err
	}

	n, err = this.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("RetMsg结构体发送失败:", err)
		return err
	}
	return nil
}
