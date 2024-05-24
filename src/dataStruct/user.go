package datastruct

type Data struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

type Iban struct {
	Iban string `json:"iban"`
	Bic  string `json:"bic"`
}

type PdfDocument struct {
	User         Data               `json:"user"`
	Iban         []Iban             `json:"banks"`
	Products     []ProductData      `json:"products"`
	Transactions []TransactionsData `json:"transactions"`
}