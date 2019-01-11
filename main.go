package main

import (
	//"fmt"
	"fmt"
	"sendpulse/api"
	"sendpulse/api/client"
)

const(
	campaignsFileName = "CampaignsReport.xlsx"
)

func main(){


	client := client.New(
		"c4013a11dba026e8f052739376a27045",
		"6e1fcffe7b01c3836e4f791471274620",
		nil)
	books, err := client.GetAddressBooks()
	if err !=nil{
		panic(err)
	}
	for _, book := range *books{
		fmt.Println("book: " + book.Name)
	}

	campaigns, err := client.GetCampaigns()
	for _, campaign := range *campaigns{
		client.GetCampaignInfo(campaign)
	}
	if err != nil{
		panic(err)
	}
	api.WriteCampaignsInfoToExcelFile(campaignsFileName,campaigns)

}
