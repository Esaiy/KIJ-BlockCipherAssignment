package main

import (
	"fmt"
	"kij-block-cipher/decrypt"
)

func main() {
	var dummy_matrix = [][]byte{
		{0x32, 0x88, 0x31, 0xe0},
		{0x43, 0x5a, 0x31, 0x37},
		{0xf6, 0x30, 0x98, 0x07},
		{0xa8, 0x8d, 0xa2, 0x34},
	}

	data := decrypt.AddRoundKey(dummy_matrix, 0)
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			fmt.Printf("%x ", data[i][j])
		}
		fmt.Printf("\n")
	}
}
