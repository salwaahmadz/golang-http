package app

import (
	"encoding/xml"
)

type Product struct {
	XMLName xml.Name `xml:"product" json:"-"`
	Base
	Price    int64    `json:"price" xml:"price"`
	Year     int      `json:"-" xml:"year"`
	Category Category `json:"category" xml:"category"`
}
