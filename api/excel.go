package api

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"sendpulse/api/client/datastructs"
	"sendpulse/excel"
	"strconv"
)

const (
	campaignName      = "Название рассылки"
	campaighSendDate  = "Дата отправки"
	bookNamesList     = "Книги-получатели"
	senderName        = "Имя отправителя"
	senderEmail       = "Почта отправителя"
	emailQuantity     = "Количество получателей"
	delivered         = "Доставлено"
	deliveredStatus   = "Delivered"
	openCount         = "Количество открытий"
	openCountStatus   = "Opened"
	linkCount         = "Количество переходов"
	linkCountStatus   = "Link redirected"
	openedOfDelivered = "% открытий от доставленных"
	linkOfOpened      = "% переходов от открытий"
)

type ActualStats struct {
}

func WriteCampaignsInfoToExcelFile(fileName string, campaigns *datastructs.Campaigns) error {
	table := &excel.Table{}

	for _, campaign := range *campaigns {
		table.AddRow()
		table.SetCellValue(campaignName, campaign.Name)
		table.SetCellValue(campaighSendDate, campaign.SendDate)
		table.SetCellValue(bookNamesList, campaign.Message.BookNamesList)
		table.SetCellValue(senderName, campaign.Message.SenderName)
		table.SetCellValue(senderEmail, campaign.Message.SenderEmail)
		table.SetCellValue(emailQuantity, strconv.Itoa(campaign.AllEmailQuantity))
		table.SetCellValue(delivered, strconv.Itoa(campaign.StatusMap[deliveredStatus]))
		table.SetCellValue(openCount, strconv.Itoa(campaign.StatusMap[openCountStatus]))
		table.SetCellValue(linkCount, strconv.Itoa(campaign.StatusMap[linkCountStatus]))
		table.SetCellValue(openedOfDelivered, countPercents(campaign.StatusMap[deliveredStatus], campaign.StatusMap[openCountStatus]))
		table.SetCellValue(linkOfOpened, countPercents(campaign.StatusMap[openCountStatus], campaign.StatusMap[linkCountStatus]))
		for _, stat := range campaign.Statistics {
			table.SetCellValue("status \""+stat.Explain+"\"", strconv.Itoa(stat.Count))
		}
	}
	xlsx := excelize.NewFile()
	excel.WriteTable(xlsx, "Campaigns", table)
	xlsx.DeleteSheet("Sheet1")
	err := xlsx.SaveAs(fileName)
	if err != nil {
		return err
	}
	return nil
}

func countPercents(full int, part int) string {
	var res int
	if full == 0 {
		res = 0
	} else {
		res = 100 * part / full
	}
	return strconv.Itoa(res) + "%"
}
