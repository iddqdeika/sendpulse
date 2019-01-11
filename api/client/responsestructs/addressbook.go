package responsestructs

type AddressBook struct {
	Id						int		`json:"id"`
	Name					string	`json:"name"`
	AllEmailQuantity		int		`json:"all_email_qty"`
	ActiveEmailQuantity		int		`json:"active_email_qty"`
	InactiveEmainQuantity	int		`json:"inactive_email_qty"`
	CreationDate			string	`json:"creationdate"`
	Status					int		`json:"status"`
	StatusExplain			string	`json:"status_explain"`
}

