package main

import (
	"log"
	"net"

	"github.com/nekomeowww/vig/config"
	"github.com/nekomeowww/vig/router"
)

func main() {
	config.Init()
	r := router.InitRouter()
	err := r.Run(net.JoinHostPort(config.Conf.IP, config.Conf.Port))
	if err != nil {
		log.Fatalf("Failed to bind to port %s", config.Conf.Port)
	}
}
