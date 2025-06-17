package model

type Campaign struct {
	ID     string
	Name   string
	Image  string
	CTA    string
	Status string
}

type TargetingRule struct {
	CampaignID string
	Dimension  string // "country", "os", or "app"
	Type       string // "include" or "exclude"
	Value      string // actual value like "India" or "Android"
}

/*var Campaigns = []Campaign{
	{ID: "spotify",Name:"Spotify - Music for everyone",Image:"https://somelink",CTA:"Download",Status:"ACTIVE"},
	{ID: "duolingo",Name:"Duolingo: Best way to learn",Image:"https://somelink2",CTA:"Install",Status:"ACTIVE"},
	{ID: "subwaysurfer",Name:"Subway Surfer",Image:"https://somelink3",CTA:"Play",Status:"ACTIVE"},
}

var Targeting_rules = []TargetingRule{
	{CampaignID:"spotify",Dimension:"country",Type:"include",Value:"us"},
	{CampaignID:"spotify",Dimension:"country",Type:"include",Value:"canada"},
	{CampaignID:"duolingo",Dimension:"os",Type:"include",Value:"android"},
	{CampaignID:"duolingo",Dimension:"os",Type:"include",Value:"ios"},
	{CampaignID:"duolingo",Dimension:"country",Type:"exclude",Value:"us"},
	{CampaignID:"subwaysurfer",Dimension:"os",Type:"include",Value:"android"},
	{CampaignID:"subwaysurfer",Dimension:"app",Type:"include",Value:"com.gametion.ludokinggame"},
}*/