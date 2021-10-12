package main

import (
	"fmt"
	"kij-block-cipher/decrypt"
)

var key = [16]byte{
	0x2b, 0x28, 0xab, 0x09,
	0x7e, 0xae, 0xf7, 0xcf,
	0x15, 0xd2, 0x15, 0x4f,
	0x16, 0xa6, 0x88, 0x3c,
}

func main() {
	something := decrypt.ExportedKey
	for i := 0; i < 4; i++ {
		for j := 0; j < 40; j++ {
			fmt.Printf("%x ", something[i][j])
		}
		fmt.Println()
	}
}
