package main

import (
	"flag"
	"fmt"

	{{.ImportPackages}}
)

var configFile = flag.String("f", "etc/{{.ServiceName}}.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	{{if .UseGin}}server := rest.MustNewGinServer(c.RestConf){{else}}server := rest.MustNewServer(c.RestConf){{end}}
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
