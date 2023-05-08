package entity

import (
	"log"
	"testing"
)

func TestGetSerialNumber(t *testing.T) {
	c := NewCommand()
	values := c.GetInfoCMD("wmic bios get serialnumber /format:csv")
	log.Println(values)
}
