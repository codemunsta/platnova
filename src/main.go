package main

import (
	"github.com/jung-kurt/gofpdf"
)

func main() {
	pdf := gofpdf.New(gofpdf.OrientationPortrait, "mm", "A4", "")
	margin := 10.0
	pdf.SetMargins(margin, margin, margin+10)
	pdf.AddPage()
}