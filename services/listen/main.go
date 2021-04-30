package main

import (
	"github.com/asim/go-micro/v3"
	"sxx-go-micro/Common/config"
	"sxx-go-micro/Common/service"
	proto "sxx-go-micro/proto"
	"sxx-go-micro/services/listen/handler"
)

func main() {
	service.Create(
		config.SERVICE_LISTEN,
		func(service micro.Service) {
			// 注册处理函数
			proto.RegisterDemoServiceHandler(service.Server(), new(handler.DemoServiceHandler))
		})
}