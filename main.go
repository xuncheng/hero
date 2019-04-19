package main

import (
	"fmt"
	"github.com/caarlos0/env"
	"github.com/gin-gonic/gin"
	"github.com/xuncheng/hero/api"
	"github.com/xuncheng/hero/model"
	"github.com/xuncheng/hero/mysql"
)

func main() {
	// connect mysql
	cfg := model.Config{}
	if err := env.Parse(&cfg); err != nil {
		panic(err)
	}

	mysqlDSN := fmt.Sprintf("mysql://%s:%s@tcp(%s:%s)/%s", cfg.MysqlUser, cfg.MysqlPass, cfg.MysqlHost, cfg.MysqlPort, cfg.MysqlDB)
	if err := mysql.Connect(mysqlDSN); err != nil {
		panic(err)
	}
	defer func() {
		if err := mysql.Close(); err != nil {
			panic(err)
		}
	}()

	// create a gin router with default middleware
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	router.GET("/v1/ping", api.Ping)

	if err := router.Run(":8000"); err != nil {
		panic(err)
	}

}
