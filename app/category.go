package app

import (
	"encoding/xml"
)

type Category struct {
	XMLName xml.Name `json:"-" xml:"category"`
	Base
}
