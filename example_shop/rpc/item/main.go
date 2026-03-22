package main

import (
	"log"

	item "example_shop/kitex_gen/example/shop/item/itemservice"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {
	// 使用时请传入真实 etcd 的服务地址，本例中为 127.0.0.1:2379
	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}

	itemServiceImpl := new(ItemServiceImpl)
	stockCli, err := NewStockClient()
	if err != nil {
		log.Fatal(err)
	}
	itemServiceImpl.stockCli = stockCli

	svr := item.NewServer(itemServiceImpl,
		// 指定 Registry 与服务基本信息
		server.WithRegistry(r),
		server.WithServerBasicInfo(
			&rpcinfo.EndpointBasicInfo{
				ServiceName: "example.shop.item",
			}),
	)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
