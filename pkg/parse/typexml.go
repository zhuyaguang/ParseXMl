package parse

import "encoding/xml"

type PatentDocumentAndRelated struct {
	XMLName xml.Name `xml:"PatentDocumentAndRelated"`
	Description Description
}

type Description struct {
	InventionTitle InventionTitle
}

type InventionTitle struct {
	content string
}