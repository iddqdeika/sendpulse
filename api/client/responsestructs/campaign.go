package responsestructs

type Campaign struct{
	Id					int		`json:"id"`
	Name				string	`json:"name"`
	Message				Message	`json:"message"`
	Status				int		`json:"status"`
	AllEmailQuantity	int		`json:"all_email_qty"`
	TariffEmailQuantity	int		`json:"tariff_email_qty"`
	PaidEmailQuantity	int		`json:"paid_email_qty"`
	OverdraftPrice		int		`json:"overdraft_price"`
	OverdraftCurrency	string	`json:"overdraft_currency"`
	Statistics []Statictic	`json:"statistics"`
}

type Message struct{
	SenderName	string	`json:"sender_name"`
	SenderEmail	string	`json:"sender_email"`
	Subject		string	`json:"subject"`
	Body		string	`json:"body"`
	ListId		int		`json:"list_id"`
	Attachments	string	`json:"attachments"`
}


type Statictic struct {
	Code		int		`json:"code"`
	Count		int		`json:"count"`
	Explain		string	`json:"explain"`
}