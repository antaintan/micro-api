package main

import (
	"huochain/gomicro-demo/api/config"
	"log"

	"github.com/micro/go-web"
)

func main() {
	service := web.NewService(web.Name("go.micro.api"))
	service.init()
	service.Handle("/", config.BuildRouter)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
