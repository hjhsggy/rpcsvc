package services

import (
	"context"
	"encoding/json"
	"fmt"

	pb "github.com/hjhsggy/rpcsvc/proto"
)

func (d *DemoServiceImpl) ListTag(ctx context.Context, req *pb.ListTagRequest) (*pb.ListTagResponse, error) {

	data := `[{"tag_id":1,"tag_name":"1","tag_status": 1},
	{"tag_id":2,"tag_name":"2","tag_status":1},
	{"tag_id":3,"tag_name":"3","tag_status":2}]`

	rsp := &pb.ListTagResponse{}

	err := json.Unmarshal([]byte(data), &rsp.Item)
	if err != nil {
		fmt.Printf("Error unmarsh:%v", err)
		return nil, err
	}

	fmt.Println("list tag ok")

	return rsp, nil
}
