package main

import (
	"flag"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/smokecat/goweb-components/pkg/framework/go-zero/fw"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/rest/httpx"

	{{.importPackages}}
)

var configFile = flag.String("f", "etc/{{.serviceName}}.toml", "the config file")

func main() {
	flag.Parse()

	// Load the configuration file
	var c config.Config
	conf.MustLoad(*configFile, &c)

	// Set up the logger
	logc.MustSetup(c.Logger)
	logc.AddGlobalFields(
		logc.Field("service", c.Logger.ServiceName),
	)

	// Set up the validator
	xvalidator.SetDefault(nil)

	// Set response handler
	httpx.SetOkHandler(fw.OkResponseHandler)
	httpx.SetErrorHandlerCtx(fw.ErrHandler)

	// Run boot tasks before starting the server
	// logc.Must(boot.Init(context.Background(), c))

	// Start the server
	server := rest.MustNewServer(c.Rest)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Rest.Host, c.Rest.Port)
	server.Start()
}
