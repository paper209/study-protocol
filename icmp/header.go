package icmp

import (
	"encoding/binary"
	"study/checksum"
)

// total bytes = 4
type Header struct {
	Type     uint8  // 1
	Code     uint8  // 1
	Checksum uint16 // 2
}

// total bytes = 4
type Echo struct {
	Identifier uint16
	Sequence   uint16
	Data       []byte
}

func (h *Header) BuildHeader(buf []byte) {
	buf[0] = h.Type
	buf[1] = h.Code
	binary.BigEndian.PutUint16(buf[2:4], 0)
}

func (e *Echo) BuildEchoPacket() []byte {
	buf := make([]byte, 8+len(e.Data))
	// build header
	h := &Header{
		Type:     8,
		Code:     0,
		Checksum: 0,
	}
	h.BuildHeader(buf)

	binary.BigEndian.PutUint16(buf[4:6], e.Identifier)
	binary.BigEndian.PutUint16(buf[6:8], e.Sequence)

	// data
	copy(buf[8:], e.Data)

	// checksum
	binary.BigEndian.PutUint16(buf[2:4], checksum.Checksum(buf))

	return buf
}
