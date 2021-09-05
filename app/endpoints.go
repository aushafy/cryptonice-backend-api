package app

import (
	"github.com/aushafy/cryptonice-api/controllers/ping"
	"github.com/aushafy/cryptonice-api/controllers/scanner"
)

func mapEndpoints() {

	router.GET("/ping", ping.Ping)

	v1 := router.Group("/api/v1")
	{
		v1.GET("/scan/:url", scanner.RunHTTPScanner)
	}
}
