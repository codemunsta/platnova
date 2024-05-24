package utils

import (
	"encoding/json"
	"fmt"

	datastruct "github.com/codemunsta/platnova/dataStruct"
)

func MockUpData() (datastruct.PdfDocument, error) {

	pdfData := `{
		"user": {
			"name": "Sandra Saulgrieze",
			"address": "14 The Dale, Whitefield Hall, Bettystown, Meath. A92N27C"
		},
		"banks": [
			{
				"iban": "IE30REV099036022547749",
				"bic": "REVOIE23"
			},
			{
				"iban": "ITO93250041069208595",
				"bic": "REVOLT21"
			},
			{
				"iban": "IT087070024346246713",
				"bic": "RETBLT21"
			}
		],
		"products": [
			{
				"product": "Account(Current Account)",
				"balance": "$2.52",
				"money_out": "$1,944.09",
				"money_in": "$1,978.00",
				"closing_balance": "$34.43"
			}
		],
		"transactions": [
			{
				"date": "3 Feb 2023",
				"description": "Apple Pay Top-Up by *5453",
				"money_out": "",
				"money_in": "$50.00",
				"balance": "$52.52"
			},
			{
				"date": "3 Feb 2023",
				"description": "Apple Pay Tp-Up by *5453",
				"money_out": "",
				"money_in": "$100.00",
				"balance": "$152.00"
			},
			{
				"date": "3 Feb 2023",
				"description": "To LINA MILLER SAULGRIEZE",
				"money_out": "$100.00",
				"money_in": "",
				"balance": "$52.52"
			},
			{
				"date": "7 Feb 2023",
				"description": "To LINA MILLER SAULGRIEZE",
				"money_out": "$10.00",
				"money_in": "",
				"balance": "$42.52"
			}
		]
	}`
	var data datastruct.PdfDocument
	err := json.Unmarshal([]byte(pdfData), &data)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return data, err
	}
	return data, nil
}