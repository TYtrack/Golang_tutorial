/*
 * @Author: TYtrack
 * @Description: ...
 * @Date: 2022-01-11 23:28:08
 * @LastEditors: TYtrack
 * @LastEditTime: 2022-01-13 15:16:48
 * @FilePath: /grpc_demo/service/laptop_server_test.go
 */

package service

import (
	"context"
	"pcbook/pb"
	"pcbook/sample"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestCreateLaptop(t *testing.T) {
	t.Parallel()
	laptop_with_noid := sample.NewLaptop()
	laptop_with_noid.Id = ""

	laptop_invalid_id := sample.NewLaptop()
	laptop_invalid_id.Id = "invalid_id"

	laptop_duplicate := sample.NewLaptop()
	store_duplicateID := NewInMemoryLaptopStore()
	store_duplicateID.Save(laptop_duplicate)

	testCase := []struct {
		name   string
		laptop *pb.Laptop
		store  LaptopStore
		code   codes.Code
	}{
		{
			name:   "success with id",
			laptop: sample.NewLaptop(),
			store:  NewInMemoryLaptopStore(),
			code:   codes.OK,
		}, {
			name:   "success no id",
			laptop: laptop_with_noid,
			store:  NewInMemoryLaptopStore(),
			code:   codes.OK,
		}, {
			name:   "success invalid id",
			laptop: laptop_invalid_id,
			store:  NewInMemoryLaptopStore(),
			code:   codes.InvalidArgument,
		}, {
			name:   "success duplicate id",
			laptop: laptop_duplicate,
			store:  store_duplicateID,
			code:   codes.AlreadyExists,
		},
	}

	for i := range testCase {
		tc := testCase[i]

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			req := &pb.CreateLaptopRequest{
				Laptop: tc.laptop,
			}
			laptop_server := NewLaotopServer(tc.store, nil, nil)
			response, err := laptop_server.CreateLaptop(
				context.Background(),
				req,
			)
			if tc.code == codes.OK {
				require.NoError(t, err)

				require.NotNil(t, response)
				require.NotEmpty(t, response.Id)

				if len(tc.laptop.Id) > 0 {
					require.Equal(t, response.Id, tc.laptop.Id)
				}

			} else {
				require.Error(t, err)
				require.Nil(t, response)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, tc.code, st.Code())
			}

		})

	}

}
