package main

import (
	"github.com/gin-gonic/gin"
	"github.com/playerpiano05/lhsqb/src/api"
)

func main() {

	testC := newConfig()
	r := gin.Default()
	r.Use()
	api.SetupRoutes(r)
	r.Run(testC.serverAddress)

}
