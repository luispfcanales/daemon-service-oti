package entity

import (
	"strings"
	"testing"
)

func TestFn_MustExecuteCommand(t *testing.T) {
	s := NewCommand().MustExecCmd("ipconfig | findstr /i ipv4")

	want := "192.168.0.10"

	v := strings.Split(string(s), ":")
	got := strings.Trim(v[1], "\n")
	got = strings.TrimSpace(got)

	if want != got {
		t.Fail()
	}
}
