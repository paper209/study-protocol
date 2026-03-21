package udp

import (
	"encoding/binary"
	"study/checksum"
)

// checksum
func Checksum(h *Header, ph *PseudoHeader, body []byte) bool {
	totalSize := 20 + len(body)

	var buf []byte
	if totalSize%2 != 0 {
		buf = make([]byte, totalSize+1)
		buf[totalSize] = 0
	} else {
		buf = make([]byte, totalSize)
	}

	// pseudo header
	copy(buf[0:4], ph.SourceAddress[:])
	copy(buf[4:8], ph.DestinationAddress[:])
	buf[8] = ph.Zero
	buf[9] = ph.Protocol
	binary.BigEndian.PutUint16(buf[10:12], ph.Length)

	// header
	binary.BigEndian.PutUint16(buf[12:14], h.SourcePort)
	binary.BigEndian.PutUint16(buf[14:16], h.DestinationPort)
	binary.BigEndian.PutUint16(buf[16:18], h.Length)
	binary.BigEndian.PutUint16(buf[18:20], h.Checksum)

	copy(buf[20:], body)

	return ^checksum.Checksum(buf) == 0xFFFF
}

// get checksum
func (h *Header) getChecksum(ph *PseudoHeader, body []byte) uint16 {
	var buf []byte
	totalSize := 12 + h.Length
	if totalSize%2 != 0 {
		buf = make([]byte, totalSize+1)
		buf[totalSize] = 0 // padding
	} else {
		buf = make([]byte, totalSize)
	}

	// ipv4 pseudo header
	copy(buf[0:4], ph.SourceAddress[:])
	copy(buf[4:8], ph.DestinationAddress[:])
	buf[8] = ph.Zero
	buf[9] = ph.Protocol
	binary.BigEndian.PutUint16(buf[10:12], ph.Length)

	// header
	binary.BigEndian.PutUint16(buf[12:14], h.SourcePort)
	binary.BigEndian.PutUint16(buf[14:16], h.DestinationPort)
	binary.BigEndian.PutUint16(buf[16:18], h.Length)
	binary.BigEndian.PutUint16(buf[18:20], 0) // checksum

	copy(buf[20:], body)

	return checksum.Checksum(buf)
}
