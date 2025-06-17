package util

import (
	"AdCampaign/model"
	"AdCampaign/db"
	"log"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeliveryService(c *gin.Context){
	app:= c.Query("app")
	os := c.Query("os")
	country := c.Query("country")

	//fmt.Println(app,os,country)

	if app == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing app param"})
		return
	}
	if os == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing os param"})
		return
	}
	if country == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing country param"})
		return
	}

	var matched []gin.H

	
	campaigns,err := db.GetActiveCampaigns()
	if err!=nil{
		log.Fatal("error while fetching records from campaigns table",err)
	}

	for _, campaign := range campaigns {

		var rules []model.TargetingRule
		targeting_rules ,errTarget := db.GetTargetingRules(campaign.ID)
		if errTarget!=nil{
			log.Fatal("error while fetching record from targeting_rules table",errTarget)
		}
		for _,rule := range targeting_rules{
			if rule.CampaignID == campaign.ID{
			rules = append(rules,rule)
			}
		}
		//fmt.Println(rules)

		if CampaignMatches(rules, app, os, country) {
			matched = append(matched, gin.H{
				"cid": campaign.ID,
				"img": campaign.Image,
				"cta": campaign.CTA,
			})
		}
	}

	if len(matched) == 0 {
		c.Status(http.StatusNoContent)
		return
	}

	c.JSON(http.StatusOK, matched)

}

func CampaignMatches(rules []model.TargetingRule, app, os, country string) bool {
	dimensions := map[string]string{
		"app":     app,
		"os":      os,
		"country": country,
	}

	for dim, reqVal := range dimensions {

		includeList := []string{}
		excludeList := []string{}
		for _, rule := range rules {
			if rule.Dimension != dim {
				continue
			}
			if rule.Type == "include" {
				includeList = append(includeList, rule.Value)
			} else if rule.Type == "exclude" {
				excludeList = append(excludeList, rule.Value)
			}
		}

		if len(includeList) > 0 {
			matched := false
			for _, val := range includeList {
				if val == reqVal {
					matched = true
					break
				}
			}
			if !matched {
				return false
			}
		}
		for _, val := range excludeList {
			if val == reqVal {
				return false
			}
		}
	}

	return true
}