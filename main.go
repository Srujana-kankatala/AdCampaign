package main

import (
	"AdCampaign/util"
	"AdCampaign/db"
	"github.com/gin-gonic/gin"
	prometheus "github.com/zsais/go-gin-prometheus"
)

func main(){

	db.InitDB()
	r := gin.Default()

	p := prometheus.NewPrometheus("ad_campaign_service")
	p.Use(r)

	r.GET("/v1/delivery",util.DeliveryService)
	r.Run("localhost:8080")
}