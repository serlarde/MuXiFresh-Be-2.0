package main

import (
	"flag"
	"fmt"

	"MuXiFresh-Be-2.0/app/modify/cmd/rpc/adjust/adjust"
	"MuXiFresh-Be-2.0/app/modify/cmd/rpc/adjust/internal/config"
	"MuXiFresh-Be-2.0/app/modify/cmd/rpc/adjust/internal/server"
	"MuXiFresh-Be-2.0/app/modify/cmd/rpc/adjust/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/adjust.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		adjust.RegisterModifyServer(grpcServer, server.NewModifyServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
