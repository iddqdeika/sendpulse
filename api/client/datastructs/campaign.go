package datastructs

import (
	"errors"
	"strconv"
)

type Campaign struct {
	Id                  int         `json:"id"`
	Name                string      `json:"name"`
	Message             Message     `json:"message"`
	Status              int         `json:"status"`
	AllEmailQuantity    int         `json:"all_email_qty"`
	TariffEmailQuantity int         `json:"tariff_email_qty"`
	PaidEmailQuantity   int         `json:"paid_email_qty"`
	OverdraftPrice      int         `json:"overdraft_price"`
	OverdraftCurrency   string      `json:"overdraft_currency"`
	Statistics          []Statictic `json:"statistics"`
	StatusMap           map[string]int
}

type Message struct {
	SenderName    string      `json:"sender_name"`
	SenderEmail   string      `json:"sender_email"`
	Subject       string      `json:"subject"`
	Body          string      `json:"body"`
	ListId        interface{} `json:"list_id"`
	BookNamesList string
	BooksIdList   []string
	Attachments   string `json:"attachments"`
}

type Statictic struct {
	Code    int    `json:"code"`
	Count   int    `json:"count"`
	Explain string `json:"explain"`
}

func (tcampaign *Campaign) ProcessBookList() error {
	sl := make([]string, 0)
	switch tcampaign.Message.ListId.(type) {
	case float64:
		i := int(tcampaign.Message.ListId.(float64))
		sl = append(sl, strconv.Itoa(i))
		tcampaign.Message.BooksIdList = sl
	case int:
		i := tcampaign.Message.ListId.(int)
		sl = append(sl, strconv.Itoa(i))
		tcampaign.Message.BooksIdList = sl
	case []interface{}:
		i := tcampaign.Message.ListId.([]interface{})
		for _, v := range i {
			switch v.(type) {
			case float64:
				temp := int(v.(float64))
				sl = append(sl, strconv.Itoa(temp))
			default:
				return errors.New("unknown listId type in campaign " + strconv.Itoa(tcampaign.Id))
			}
		}
		tcampaign.Message.BooksIdList = sl
	default:
		return errors.New("unknown listId type in campaign " + strconv.Itoa(tcampaign.Id))
	}
	return nil
}
