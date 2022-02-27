/*
 * @Author: TYtrack
 * @Description: ...
 * @Date: 2022-01-07 14:27:16
 * @LastEditors: TYtrack
 * @LastEditTime: 2022-01-07 14:52:43
 * @FilePath: /grpc_demo/serializer/file_test.go
 */

package serializer

import (
	"pcbook/pb"
	"pcbook/sample"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/require"
)

func TestWriteProtobufToBinaryFile(t *testing.T) {
	t.Parallel()

	laptop := sample.NewLaptop()

	binaryFile := "../tmp/laptop.bin"
	jsonFile := "../tmp/laptop.json"

	err := WriteProtobufToBinaryFile(laptop, binaryFile)
	require.NoError(t, err)

	laptop2 := &pb.Laptop{}
	err = ReadProtobufFromBinaryFile(binaryFile, laptop2)
	require.NoError(t, err)
	require.True(t, proto.Equal(laptop, laptop2))

	err = WriteProtobufToJSONFile(laptop, jsonFile)
	require.NoError(t, err)

}
