//	upgrade
//
//	Description: upgrade service
//
//	Schemes: http, https
//	Host: localhost:8181
//	BasePath: /
//	Version: 0.0.1
//	SecurityDefinitions:
//	  Token:
//	    type: apiKey
//	    name: Authorization
//	    in: header
//	Security:
//	  Token:
//	Consumes:
//	  - application/json
//
//	Produces:
//	  - application/json
//
// swagger:meta
package main

import (
	"flag"
	"fmt"

	"upgradelink-admin-upgrade/server/internal/config"
	"upgradelink-admin-upgrade/server/internal/handler"
	"upgradelink-admin-upgrade/server/internal/svc"
	"upgradelink-admin-upgrade/server/internal/utils/osx"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/upgrade.yaml", "the config file")

func main() {
	flag.Parse()

	fmt.Println(osx.Env())

	var c config.Config
	conf.MustLoad(*configFile, &c, conf.UseEnv())

	server := rest.MustNewServer(c.RestConf, rest.WithCors(c.CROSConf.Address))
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
