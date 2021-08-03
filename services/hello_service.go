package services

import (
	"context"

	pb "github.com/hjhsggy/rpcsvc/proto"
)

func (d *DemoServiceImpl) SayHello(ctx context.Context, req *pb.SayHelloRequest) (*pb.SayHelloResponse, error) {

	rsp := &pb.SayHelloResponse{}
	rsp.Msg = "hello"

	return rsp, nil
}
