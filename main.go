package main

import (
	"AdCampaign/util"
	"AdCampaign/db"
	"github.com/gin-gonic/gin"
)

func main(){

	db.InitDB()
	r := gin.Default()
	r.GET("/v1/delivery",util.DeliveryService)
	r.Run("localhost:8080")
}