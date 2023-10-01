package entity

import (
	"testing"
)

func TestGetSerialNumber(t *testing.T) {
	c := NewCommand()
	values := c.GetInfoCMD("wmic bios get serialnumber /format:csv")
	if len(values) <= 1 {
		t.Fail()
	}
}
