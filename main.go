package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	// generate unoptimal alg first (read/write to shared storage per GUID, )
	// fmt.Println("time stamp:")
	// fmt.Println(getTimestamp())
	fmt.Println(getGUIDV4())
}

// generate GUID Version 4
func getGUIDV4() string {
	// 	1) Generate 16 random bytes (=128 bits)
	bytes := make([]byte, 16)
	rand.Read(bytes)
	//  2) Adjust certain bits according to RFC 4122 section 4.4 as follows:
	// 		a) set the four most significant bits of the 7th byte to 0100'B, so the high nibble is "4"
	// 		b) set the two most significant bits of the 9th byte to 10'B, so the high nibble will be one of "8", "9", "A", or "B" (see Note 1).
	seventhByte := bytes[6]
	seventhByte = clearBit(seventhByte, 7)
	seventhByte = setBit(seventhByte, 6)
	seventhByte = clearBit(seventhByte, 5)
	seventhByte = clearBit(seventhByte, 4)

	ninthByte := bytes[8]
	ninthByte = setBit(ninthByte, 7)
	ninthByte = clearBit(ninthByte, 6)
	//  3) Encode the adjusted bytes as 32 hexadecimal digits
	//  4) Add four hyphen "-" characters to obtain blocks of 8, 4, 4, 4 and 12 hex digits
	var sb strings.Builder
	for i, b := range bytes {
		if i == 5 || i == 7 || i == 9 || i == 11 {
			sb.WriteString("-")
		}
		fmt.Fprintf(&sb, "%02x", b)
	}
	//  5) Output the resulting 36-character string "XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX"
	return sb.String()

}

func clearBit(b byte, i int) byte {
	return b &^ (1 << i)
}

func setBit(b byte, i int) byte {
	return b | (1 << i)
}
