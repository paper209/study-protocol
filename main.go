package main

import (
	"study/icmp"
)

func main() {
	/*err := udp.Send("127.0.0.1", 8888, []byte("udptest1234"))
	if err != nil {
		panic(err)
	}*/
	err := icmp.SendEcho("127.0.0.1", []byte{})
	if err != nil {
		panic(err)
	}
}
