package db

import (
	"database/sql"
	"AdCampaign/model"
	_ "github.com/lib/pq"
	"log"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("postgres", "user=postgres password=Sujichandu@12 dbname=delivery sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
}

func GetActiveCampaigns() ([]model.Campaign, error) {
	rows, err := DB.Query("SELECT id, name, image, cta, status FROM campaigns WHERE status='ACTIVE'")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var campaigns []model.Campaign
	for rows.Next() {
		var c model.Campaign
		err := rows.Scan(&c.ID, &c.Name, &c.Image, &c.CTA, &c.Status)
		if err != nil {
			return nil, err
		}
		campaigns = append(campaigns, c)
	}
	return campaigns, nil
}

func GetTargetingRules(campaignID string) ([]model.TargetingRule, error) {
	rows, err := DB.Query("SELECT campaign_id, dimension, type, value FROM targeting_rules WHERE campaign_id=$1", campaignID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rules []model.TargetingRule
	for rows.Next() {
		var r model.TargetingRule
		err := rows.Scan(&r.CampaignID, &r.Dimension, &r.Type, &r.Value)
		if err != nil {
			return nil, err
		}
		rules = append(rules, r)
	}
	return rules, nil
}