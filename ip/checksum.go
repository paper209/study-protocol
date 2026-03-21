package ip

import (
	"encoding/binary"
	"study/checksum"
)

func (h *Header) getChecksum() uint16 {
	buf := make([]byte, 20)
	buf[0] = h.VersionIHL
	buf[1] = h.Tos
	binary.BigEndian.PutUint16(buf[2:], h.TotalLength)
	binary.BigEndian.PutUint16(buf[4:], h.Identification)
	binary.BigEndian.PutUint16(buf[6:], h.FlagsFragment)
	buf[8] = h.TTL
	buf[9] = h.Protocol
	copy(buf[12:16], h.SourceAddress[:])
	copy(buf[16:20], h.DestinationAddress[:])

	return checksum.Checksum(buf)
}
