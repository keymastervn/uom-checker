package uom_checker

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"runtime"
)

type UomList struct {
	UOM []Attr `xml:"UOM"`
}

type Attr struct {
	Code string `xml:"code"`
	Desc string `xml:"desc"`
}

var UL *UomList

// Prefetch UOM List
func init() {
	_, filename, _, _ := runtime.Caller(0)
	xmlFile, err := os.Open(path.Join(path.Dir(filename), "uom.xml"))
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer xmlFile.Close()

	b, _ := ioutil.ReadAll(xmlFile)

	xml.Unmarshal(b, &UL)
}

// IsUOM checks whether uomID is an UOM Code
func IsUOM(uomID string) bool {
	for _, uom := range UL.UOM {
		if uomID == uom.Code {
			return true
		}
	}
	return false
}

// main
func main() {
	// Do nothing
}
