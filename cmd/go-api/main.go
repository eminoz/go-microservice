package main

import (
	"github.com/eminoz/go-api/pkg/config"
	"github.com/eminoz/go-api/pkg/db"
	"github.com/eminoz/go-api/pkg/router"
)

func main() {

	config.SetupConfig()
	db.SetDatabase()
	f := router.SetUp()
	f.Listen(":" + config.GetConfig().Port)
}
