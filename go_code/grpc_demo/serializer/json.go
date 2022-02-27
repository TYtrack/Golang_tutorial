/*
 * @Author: TYtrack
 * @Description: ...
 * @Date: 2022-01-07 14:47:46
 * @LastEditors: TYtrack
 * @LastEditTime: 2022-01-07 14:56:30
 * @FilePath: /grpc_demo/serializer/json.go
 */
package serializer

import (
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

func ProtobufToJSON(message proto.Message) (string, error) {
	marshaler := jsonpb.Marshaler{
		EnumsAsInts:  false,
		EmitDefaults: true,
		OrigName:     true,
		Indent:       "  ",
	}
	return marshaler.MarshalToString(message)

}
