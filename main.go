package main

import (
	"fmt"
)

type GUID {
	time_low uint32
	time_mid uint16
	time_high_and_version uint16
	clock_seq_high_and_reserved uint8
	clock_seq_low uint8
	node [6]byte
}

func main() {
	fmt.Println("hello guid world")
}