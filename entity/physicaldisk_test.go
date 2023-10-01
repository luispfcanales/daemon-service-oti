package entity

import (
	"strconv"
	"strings"
	"testing"
)

func TestDiskType(t *testing.T) {
	c := NewCommand()

	want_disk := []string{"HDD", "SSD"}

	want := c.GetInfoPOWERSHELL("get-physicaldisk")
	gotRow := strings.Fields(want)

	for _, got := range gotRow {
		if want_disk[0] == got || want_disk[1] == got {
			return
		}
	}
	t.Fail()
}

func TestDiskSize(t *testing.T) {
	var gotValue string = ""
	c := NewCommand()

	afterValue := "Online"

	want := c.GetInfoPOWERSHELL("get-disk")
	gotRow := strings.Fields(want)

	for index, got := range gotRow {
		if got == afterValue {
			gotValue = gotRow[index+1]
			break
		}
	}

	_, err := strconv.ParseFloat(gotValue, 64)
	if err != nil {
		t.Fail()
	}
}
