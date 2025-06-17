package main

import (
	"AdCampaign/util"
	"github.com/gin-gonic/gin"
)

func main(){

	r := gin.Default()
	r.GET("/v1/delivery",util.DeliveryService)
	r.Run("localhost:8080")
}