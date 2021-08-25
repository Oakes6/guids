package guids

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// GUIDV4 generates a GUID of version 4
func GUIDV4() string {
	// 	Generate 16 random bytes (=128 bits)
	bytes := make([]byte, 16)
	rand.Read(bytes)
	//  Adjust certain bits according to RFC 4122 section 4.4 as follows:
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

	var sb strings.Builder
	for i, b := range bytes {
		if i == 5 || i == 7 || i == 9 || i == 11 {
			sb.WriteString("-")
		}
		fmt.Fprintf(&sb, "%02x", b)
	}
	return sb.String()

}

func clearBit(b byte, i int) byte {
	return b &^ (1 << i)
}

func setBit(b byte, i int) byte {
	return b | (1 << i)
}
