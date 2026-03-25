package icmp

import (
	"fmt"
	"net"
	"study/ip"
	"syscall"
)

func SendEcho(address string, data []byte) error {
	var dst [4]byte
	copy(dst[:], net.ParseIP(address).To4())

	ih := &ip.Header{
		VersionIHL:         (4 << 4) | (5 & 0x0F),
		Tos:                0,
		TotalLength:        29,
		Identification:     0,
		FlagsFragment:      0x4000, // df = 1
		TTL:                64,
		Protocol:           1, // icmp
		SourceAddress:      [4]byte{127, 0, 0, 1},
		DestinationAddress: dst,
	}
	buf := ih.BuildHeader()

	e := &Echo{
		Identifier: 0,
		Sequence:   0,
		Data:       data,
	}
	buf = append(buf, e.BuildEchoPacket()...)

	s, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, syscall.IPPROTO_RAW)
	if err != nil {
		return fmt.Errorf("socket errror: %s", err.Error())
	}
	defer syscall.Close(s)

	err = syscall.SetsockoptInt(s, syscall.IPPROTO_IP, syscall.IP_HDRINCL, 1)
	if err != nil {
		return fmt.Errorf("socket option error: %s", err.Error())
	}

	err = syscall.Sendto(s, buf, 0, &syscall.SockaddrInet4{
		Port: 0,
		Addr: dst,
	})
	if err != nil {
		return fmt.Errorf("socket send error: %s", err.Error())
	}

	return nil
}
