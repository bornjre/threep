package threep

import (
	"encoding/binary"
	"unsafe"
)

type Packet struct {
	Header Header
	Data   []byte
}

func Encode(p Packet) ([]byte, error) {
	return attachHeader(&p.Header, p.Data), nil
}

func Decode(data []byte) (*Packet, error) {
	h, bodyraw := getHeader(data)
	return &Packet{Header: *h, Data: bodyraw}, nil
}

type Header struct {
	Version uint8
	TypeId  uint8
	Pid     uint16
}

func NewHeader(version uint8, typeId uint8, pid uint16) *Header {
	return &Header{
		Version: version,
		TypeId:  typeId,
		Pid:     pid,
	}
}

func attachHeader(header *Header, body []byte) []byte {

	b := make([]byte, 4)

	binary.BigEndian.PutUint16(b, uint16(header.Pid))
	b[2] = byte(header.TypeId)
	b[3] = byte(header.Version)

	return append(body, b...)
}

func UpdateHeader(header *Header, body []byte) []byte {
	offset := len(body) - int(unsafe.Sizeof(Header{}))
	return attachHeader(header, body[:offset])
}

func getHeader(raw []byte) (*Header, []byte) {
	h := &Header{}
	offset := len(raw) - int(unsafe.Sizeof(Header{}))
	headerbytes := raw[offset:]

	h.Pid = (binary.BigEndian.Uint16(headerbytes[:2]))
	h.TypeId = (headerbytes[2])
	h.Version = (headerbytes[3])
	return h, raw[:offset]
}
