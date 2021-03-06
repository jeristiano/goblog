package main

import (
	"net/http"

	"goblog/app/http/middlewares"
	"goblog/config"
	c "goblog/pkg/config"

	"goblog/bootstrap"
)

func init() {
	// 初始化配置信息
	config.Initialize()
}

func main() {
	// 初始化 SQL
	bootstrap.SetupDB()

	// 初始化路由绑定
	router := bootstrap.SetupRoute()

	http.ListenAndServe(":"+c.GetString("app.port"), middlewares.RemoveTrailingSlash(router))
}
