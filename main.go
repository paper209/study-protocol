package main

import "study/udp"

func main() {
	err := udp.Send("127.0.0.1", 8888, []byte("udptest1234"))
	if err != nil {
		panic(err)
	}
}
