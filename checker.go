package uom_checker

import "encoding/xml"

type UomList struct {
	UOM []Attr `xml:"UOM"`
}

type Attr struct {
	Code string `xml:"code"`
	Desc string `xml:"desc"`
}

var UL *UomList
var uomCodes map[string]struct{}

// Prefetch UOM List
func init() {
	uomCodes = make(map[string]struct{})
	xml.Unmarshal([]byte(uoms), &UL)
	for _, uom := range UL.UOM {
		uomCodes[uom.Code] = struct{}{}
	}
}

// IsUOM checks whether uomCode is an UOM Code
func IsUOM(uomCode string) bool {
	if _, ok := uomCodes[uomCode]; ok {
		return true
	}
	return false
}
