package main

import (
	"google.golang.org/grpc"
	"net"
	"storageSystem/dao"
	"storageSystem/pbfiles"
	"storageSystem/services"
)

func main() {
	dao.Init()
	rpcServer:=grpc.NewServer()
	pbfiles.RegisterProdServiceServer(rpcServer,new(services.ProdService))
	lis,_:=net.Listen("tcp", ":8080")

	rpcServer.Serve(lis)
}