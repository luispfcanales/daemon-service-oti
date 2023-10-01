package entity

import (
	"testing"
)

func TestGetIPV4(t *testing.T) {

	ip := NewIpHostSystem(NewCommand())
	ip.WorkerLoadInfo()

	want := "192.168.0.10"

	if want != string(ip.Ipv4) {
		t.Fail()
	}
}
