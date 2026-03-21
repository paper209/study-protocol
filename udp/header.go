package udp

import (
	"encoding/binary"
	"fmt"
)

// total size = 8 bytes
type Header struct {
	SourcePort      uint16
	DestinationPort uint16
	Length          uint16
	Checksum        uint16
}

// total size = 12 bytes
type PseudoHeader struct {
	SourceAddress      [4]byte
	DestinationAddress [4]byte
	Zero               uint8
	Protocol           uint8
	Length             uint16
}

func (h *Header) buildPacket(ph *PseudoHeader, body []byte) []byte {
	buf := make([]byte, 8+len(body))
	binary.BigEndian.PutUint16(buf[0:], h.SourcePort)
	binary.BigEndian.PutUint16(buf[2:], h.DestinationPort)
	binary.BigEndian.PutUint16(buf[4:], h.Length)
	binary.BigEndian.PutUint16(buf[6:], h.getChecksum(ph, body))
	copy(buf[8:], body)
	return buf
}

func parseHeader(data []byte) (*Header, error) {
	if len(data) != 8 {
		return nil, fmt.Errorf("Invalid Header Length: %d bytes", len(data))
	}

	return &Header{
		SourcePort:      binary.BigEndian.Uint16(data[:2]),
		DestinationPort: binary.BigEndian.Uint16(data[2:4]),
		Length:          binary.BigEndian.Uint16(data[4:6]),
		Checksum:        binary.BigEndian.Uint16(data[6:]),
	}, nil
}
