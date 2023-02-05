package main

import (
    "flag"
    "fmt"

    "middle/app/user/rpc/internal/config"
    "middle/app/user/rpc/internal/server"
    "middle/app/user/rpc/internal/svc"
    "middle/app/user/rpc/user"
    "middle/common/interceptor/rpcserver"

    "github.com/zeromicro/go-zero/core/conf"
    "github.com/zeromicro/go-zero/core/service"
    "github.com/zeromicro/go-zero/zrpc"
    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
    flag.Parse()

    var c config.Config
    conf.MustLoad(*configFile, &c)
    ctx := svc.NewServiceContext(c)

    s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
        user.RegisterUserServer(grpcServer, server.NewUserServer(ctx))

        if c.Mode == service.DevMode || c.Mode == service.TestMode {
            reflection.Register(grpcServer)
        }
    })

    //rpc log
    s.AddUnaryInterceptors(rpcserver.LoggerInterceptor)

    defer s.Stop()

    fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
    s.Start()
}
