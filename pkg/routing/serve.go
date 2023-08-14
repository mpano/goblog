package routing

import (
	"fmt"
	"goblog/pkg/config"
	"log"
)

func Serve() {
	r := GetRouter()
	configs := config.Get()
	err := r.Run(fmt.Sprintf("%s:%s", configs.Server.Host, configs.Server.Port))
	if err != nil {
		log.Fatal("there is error in router")
		return
	}
}
