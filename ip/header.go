package ip

import "encoding/binary"

// total bytes = 20
type Header struct {
	VersionIHL         uint8
	Tos                uint8
	TotalLength        uint16
	Identification     uint16
	FlagsFragment      uint16
	TTL                uint8
	Protocol           uint8
	Checksum           uint16
	SourceAddress      [4]byte
	DestinationAddress [4]byte
}

func (h *Header) BuildHeader() []byte {
	buf := make([]byte, 20)
	buf[0] = h.VersionIHL
	buf[1] = h.Tos
	binary.BigEndian.PutUint16(buf[2:], h.TotalLength)
	binary.BigEndian.PutUint16(buf[4:], h.Identification)
	binary.BigEndian.PutUint16(buf[6:], h.FlagsFragment)
	buf[8] = h.TTL
	buf[9] = h.Protocol
	binary.BigEndian.PutUint16(buf[10:], h.getChecksum())
	copy(buf[12:16], h.SourceAddress[:])
	copy(buf[16:20], h.DestinationAddress[:])

	return buf
}
