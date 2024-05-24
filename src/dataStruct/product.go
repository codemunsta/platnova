package datastruct

type ProductData struct {
	Product        string  `json:"product"`
	Balance        float64 `json:"balance"`
	MoneyOut       float64 `json:"money_out"`
	MoneyIn        float64 `json:"money_in"`
	ClosingBalance float64 `json:"closing_balance"`
}

type TransactionsData struct {
	Date        string `json:"date"`
	Description string `json:"description"`
	MoneyOut    string `json:"money_out"`
	MoneyIn     string `json:"money_in"`
	Balance     string `json:"balance"`
}
