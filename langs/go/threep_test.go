package threep

import (
	"testing"
)

func TestHeader(t *testing.T) {

	mybody := []byte("hellomessage")

	connid := uint16(12)
	cmdtype := uint8(93)

	h := NewHeader(0, cmdtype, connid)

	packet := attachHeader(h, mybody)
	header, _ := getHeader(packet)

	t.Logf("%d", packet)
	t.Logf("%+v", header)

	if header.Pid != h.Pid {
		t.Error("worong pid")
	}

	if header.TypeId != h.TypeId {
		t.Error("worong msgtype")
	}
	if header.Version != h.Version {
		t.Error("worong version")
	}
}
