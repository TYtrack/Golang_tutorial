/*
 * @Author: TYtrack
 * @Description: ...
 * @Date: 2022-01-07 00:13:59
 * @LastEditors: TYtrack
 * @LastEditTime: 2022-01-07 14:51:37
 * @FilePath: /grpc_demo/serializer/file.go
 */

package serializer

import (
	"fmt"
	"io/ioutil"

	"github.com/golang/protobuf/proto"
)

func WriteProtobufToJSONFile(message proto.Message, filename string) (err error) {
	data, err := ProtobufToJSON(message)
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ioutil.WriteFile(filename, []byte(data), 0644)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func WriteProtobufToBinaryFile(message proto.Message, filename string) (err error) {
	data, err := proto.Marshal(message)
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func ReadProtobufFromBinaryFile(filename string, message proto.Message) (err error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = proto.Unmarshal(data, message)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
