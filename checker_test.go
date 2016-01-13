package uom_checker_test

import (
	// "fmt"
	"github.com/keymastervn/uom_checker"
	"testing"
)

func TestIsUOM(t *testing.T) {
	var uomID string

	// fmt.Printf("%#v", uom_checker.UL)

	uomID = "BX"
	if !uom_checker.IsUOM(uomID) {
		t.Error("OMG- FAILED")
	}
	uomID = "CA"
	if !uom_checker.IsUOM(uomID) {
		t.Error("OMG- FAILED")
	}
	uomID = "CS"
	if !uom_checker.IsUOM(uomID) {
		t.Error("OMG- FAILED")
	}
	uomID = "B8"
	if !uom_checker.IsUOM(uomID) {
		t.Error("OMG- FAILED")
	}
	uomID = "CALIFORNIA"
	if uom_checker.IsUOM(uomID) {
		t.Error("OMG- FAILED")
	}
}
