package services

import (
	"context"
	"fmt"
	"os"
	"time"

	pb "github.com/hjhsggy/rpcsvc/proto"

	"reflect"
	"testing"

	"google.golang.org/grpc"
)

func TestMain(m *testing.M) {
	fmt.Println("begin")

	opts := []grpc.ServerOption{
		grpc.ConnectionTimeout(5000 * time.Millisecond),
	}

	grpc.NewServer(opts...)

	m.Run()
	fmt.Println("end")
	os.Exit(0)
}

func TestDemoServiceImpl_ListTag(t *testing.T) {
	type args struct {
		ctx context.Context
		req *pb.ListTagRequest
	}
	tests := []struct {
		name    string
		d       *DemoServiceImpl
		args    args
		want    *pb.ListTagResponse
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "list",
			args: args{
				ctx: context.Background(),
				req: &pb.ListTagRequest{},
			},
			want: &pb.ListTagResponse{
				Item: []*pb.ListTagResponse_Data{
					{TagId: 1, TagName: "1", TagStatus: 1},
					{TagId: 2, TagName: "2", TagStatus: 1},
					{TagId: 3, TagName: "3", TagStatus: 2},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.d.ListTag(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("DemoServiceImpl.ListTag() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DemoServiceImpl.ListTag() = %v, want %v", got, tt.want)
			}
		})
	}
}
