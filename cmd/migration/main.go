package main

import (
	"context"
	"flag"
	"github.com/go-nunu/nunu-layout-advanced/cmd/migration/wire"
	"github.com/go-nunu/nunu-layout-advanced/pkg/config"
	"github.com/go-nunu/nunu-layout-advanced/pkg/log"
)

func main() {
	var envConf = flag.String("conf", "config/local.yml", "config path, eg: -conf ./config/local.yml")
	flag.Parse()
	conf := config.NewConfig(*envConf)

	logger := log.NewLog(conf)

	app, cleanup, err := wire.NewWire(conf, logger)
	defer cleanup()
	if err != nil {
		panic(err)
	}
	if err = app.Run(context.Background()); err != nil {
		panic(err)
	}
}
