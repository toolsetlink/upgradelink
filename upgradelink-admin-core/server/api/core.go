//		Simple Admin
//
//		This is simple admin api doc
//
//		Schemes: http, https
//		Host: localhost:9100
//		BasePath: /
//		Version: 1.6.3
//		Contact: yuansu.china.work@gmail.com
//		SecurityDefinitions:
//		  Token:
//		    type: apiKey
//		    name: Authorization
//		    in: header
//		Security:
//	  	  Token:
//	    Consumes:
//		  - application/json
//
//		Produces:
//		  - application/json
//
// swagger:meta
package main

import (
	"flag"
	"fmt"

	"upgradelink-admin-core/server/api/internal/config"
	"upgradelink-admin-core/server/api/internal/handler"
	"upgradelink-admin-core/server/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/core.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c, conf.UseEnv())

	server := rest.MustNewServer(c.RestConf, rest.WithCors(c.CROSConf.Address))
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
