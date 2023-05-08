package entity

import (
	"log"
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
	c := NewCommand()
	want := c.GetInfoPOWERSHELL("get-disk")
	gotRow := strings.Fields(want)
	sizeGot := len(gotRow) - 3
	log.Println(gotRow[sizeGot])
}
