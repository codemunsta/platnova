package datastruct

type ProductData struct {
	Product        string `json:"product"`
	Balance        string `json:"balance"`
	MoneyOut       string `json:"money_out"`
	MoneyIn        string `json:"money_in"`
	ClosingBalance string `json:"closing_balance"`
}

type TransactionsData struct {
	Date        string `json:"date"`
	Description string `json:"description"`
	MoneyOut    string `json:"money_out"`
	MoneyIn     string `json:"money_in"`
	Balance     string `json:"balance"`
}
