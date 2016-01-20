package uom_checker_test

import (
	// "fmt"

	"testing"

	"github.com/keymastervn/uom_checker"
)

func TestIsUOM(t *testing.T) {
	var uomID string

	uomID = "BX"
	if !uom_checker.IsUOM(uomID) {
		t.Error("OMG- FAILED")
	}

	uomID = "VL"
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
