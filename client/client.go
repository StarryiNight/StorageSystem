package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"storageSystem/pbfiles"
)

func main() {
	conn,err:=grpc.Dial("127.0.0.1:8080",grpc.WithInsecure())
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	prodClient:= pbfiles.NewProdServiceClient(conn)
	ctx:=context.Background()
	fmt.Printf("1.添加\n2.查询\n3.修改\n4.删除\n0.退出\n")
	for true{
		var (
			choose int
			res    *pbfiles.ProdResponse
			err    error
		)
		fmt.Scan(&choose)
		switch choose {
		case 0:
			return
		case 1:
			var u pbfiles.ProdRegister
			fmt.Scan(&u.Key,&u.Value)
			res, err = prodClient.RegisterProdStock(ctx, &u)

		case 2:
			var u pbfiles.ProdRequest
			fmt.Scan(&u.Key)
			res, err = prodClient.GetProdStock(ctx,&u)
		case 3:
			var u pbfiles.ProdRegister
			fmt.Scan(&u.Key,&u.Value)
			res, err = prodClient.UpdateProdStock(ctx,&u)
		case 4:
			var u pbfiles.ProdRequest
			fmt.Scan(&u.Key)
			res, err = prodClient.DeleteProdStock(ctx,&u)
		}
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(res)
	}
}
