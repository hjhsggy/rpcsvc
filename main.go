package main

import (
	"flag"
	"net"
	"net/http"

	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/hjhsggy/rpcsvc/proto"
	svc "github.com/hjhsggy/rpcsvc/services"
)

func main() {

	var port string

	flag.StringVar(&port, "port", "9000", "启动端口号")
	flag.Parse()

	addr, err := runTcpServer(port)
	if err != nil {
		panic(err)
	}

	m := cmux.New(addr)

	// 通过 content-type 区分http1.1 跟 http2.0
	grpcL := m.MatchWithWriters(cmux.HTTP2MatchHeaderFieldPrefixSendSettings("content-type", "application/grpc"))
	httpL := m.Match(cmux.HTTP1Fast())

	grpcSvr := runGrpcServer()
	httpSvr := runHttpServer(port)

	go grpcSvr.Serve(grpcL)
	go httpSvr.Serve(httpL)

	err = m.Serve()
	if err != nil {
		panic(err)
	}

}

func runTcpServer(port string) (net.Listener, error) {
	return net.Listen("tcp", ":"+port)
}

func runGrpcServer() *grpc.Server {

	s := grpc.NewServer()

	pb.RegisterDemoServer(s, &svc.DemoServiceImpl{})
	reflection.Register(s)

	return s
}

func runHttpServer(port string) *http.Server {

	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`pong`))
	})

	return &http.Server{
		Addr:    ":" + port,
		Handler: serveMux,
	}

}
