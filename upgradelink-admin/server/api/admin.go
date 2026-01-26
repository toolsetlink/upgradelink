package main

import (
	"flag"
	"fmt"
	"upgradelink-admin/server/api/internal/config"
	"upgradelink-admin/server/api/internal/crontab"
	"upgradelink-admin/server/api/internal/handler"
	"upgradelink-admin/server/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "./etc/admin.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf, rest.WithCors(c.CROSConf.Address))
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	// 定时任务调度
	scheduler, err := crontab.InitScheduler(ctx)
	if err != nil {
		fmt.Printf("Failed to initialize scheduler: %v", err)
	}
	scheduler.Start()

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
