package api

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"sendpulse/api/client/responsestructs"
	"sendpulse/excel"
	"strconv"
)

func WriteCampaignsInfoToExcelFile(fileName string, campaigns *responsestructs.Campaigns) error {
	table := &excel.Table{}

	for _, campaign := range *campaigns{
		table.AddRow()
		table.SetCellValue("Id",strconv.Itoa(campaign.Id))
		table.SetCellValue("Name",campaign.Name)
		table.SetCellValue("Sender Name",campaign.Message.SenderName)
		table.SetCellValue("Sender Email",campaign.Message.SenderEmail)
		table.SetCellValue("Message Subject",campaign.Message.Subject)
		for _, stat := range campaign.Statistics{
			table.SetCellValue("status \"" +stat.Explain + "\"",strconv.Itoa(stat.Count))
		}
	}
	xlsx := excelize.NewFile()
	excel.WriteTable(xlsx,"Campaigns",table)
	err := xlsx.SaveAs(fileName)
	if err != nil{
		return err
	}
	return nil
}
