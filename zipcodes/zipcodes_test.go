package zipcodes

import (
	_ "embed"
	"testing"
)

//go:embed testdata/zipcodes.csv
var testdata []byte

func TestLoadZipcodesInnerCanary(t *testing.T) {

	res, err := loadZipcodesInner(testdata)
	if err != nil {
		t.Error(err)
	}

	canaryZip := res.GetZoneForZip("22182")

	if canaryZip == nil {
		t.Error("got nil")
	}
	if canaryZip.Zone != "7b" {
		t.Error("canary check failed")
	}
}

func TestLoadZipcodesInnerNil(t *testing.T) {

	res, err := loadZipcodesInner(testdata)
	if err != nil {
		t.Error(err)
	}

	nilZip := res.GetZoneForZip("xxxxx")
	if nilZip != nil {
		t.Error("should be nil")
	}

}
