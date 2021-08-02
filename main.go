package main

import (
	"net"

	"github.com/nekomeowww/vig/src/config"
	"github.com/nekomeowww/vig/src/logger"
	"github.com/nekomeowww/vig/src/router"
)

func main() {
	logger.Init()
	config.Init()
	r := router.InitRouter()
	logger.Infof("vig started on http://%s:%s", config.Conf.IP, config.Conf.Port)
	err := r.Run(net.JoinHostPort(config.Conf.IP, config.Conf.Port))
	if err != nil {
		logger.Fatalf("failed to bind to port %s", config.Conf.Port)
	}
}
