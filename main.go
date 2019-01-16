package main

import (
	//"fmt"
	"fmt"
	"sendpulse/api"
	"sendpulse/api/client"
	"sendpulse/api/client/config"
	"sendpulse/logger"
	"strconv"
	"sync"
)

const (
	campaignsFileName = "CampaignsReport"
	configName        = "config.cfg"
)

type Logger interface {
	Log(message string)
	Alert(message string)
}

var log Logger = &logger.DefaultLogger{}

func main() {

	cfg, err := config.GetConfig(configName)
	if err != nil {
		log.Alert(err.Error())
		log.Log("press ENTER to quit")
		fmt.Scanln()
		return
	}

	clients := make([]*client.Client, 0)

	wg := sync.WaitGroup{}

	for _, acc := range cfg.Accounts {
		client := client.New(
			acc.Name,
			acc.Id,
			acc.Secret,
			log)
		clients = append(clients, client)
		wg.Add(1)
		go func() {
			defer wg.Done()
			processClient(client)
		}()
	}

	wg.Wait()
	log.Log("Complete! Press ENTER to exit.")
	fmt.Scanln()
}

func processClient(client *client.Client) {
	log.Log("starting processing client \"" + client.Name() + "\"...")
	books, err := client.GetAddressBooks()
	if err != nil {
		log.Alert("books list getting error - " + err.Error())
	}

	campaigns, err := client.GetCampaigns()
	if err != nil {
		log.Alert("campaign list getting error - " + err.Error())
	}
	for _, campaign := range *campaigns {
		client.GetCampaignInfo(campaign)
		names := ""
		for k, v := range campaign.Message.BooksIdList {
			book := books[v]
			if book == nil {
				log.Alert("cant find book by id: \"" + v + "\"")
				break
			}
			if k != 0 {
				names += ", "
			}
			names += book.Name + "(" + strconv.Itoa(book.ActiveEmailQuantity) + ")"
		}
		campaign.Message.BookNamesList = names
	}
	if err != nil {
		panic(err)
	}
	api.WriteCampaignsInfoToExcelFile(campaignsFileName+"_"+client.Name()+".xlsx", campaigns)
	log.Log("finished processing client \"" + client.Name() + "\"...")
}
