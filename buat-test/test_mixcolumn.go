package main

import "fmt"

func multiply_by_2(the_value byte) byte {
	var new_value byte
	new_value = the_value << 1
	new_value = new_value & 0xff

	if (new_value & 128) != 0 {
		new_value = new_value ^ 0x1b
	}

	return new_value
}

func multiply_by_3(the_value byte) byte {
	var the_result byte
	the_result = multiply_by_2(the_value) ^ the_value
	return the_result
}

func Aes_decrypt_scratch_mixcolumn(thirdstep_array [][]byte) [][]byte {
	// reference: https://gist.github.com/vwxyzjn/bcac5f97b5abb7708773a28b82a809b4
	// reference: https://blog.tclaverie.eu/posts/understanding-golangs-aes-implementation-t-tables/
	// multiply
	var static_matrix = [4][4]byte{
		{02, 03, 01, 01},
		{01, 02, 03, 01},
		{01, 01, 02, 03},
		{03, 01, 01, 02},
	}

	var new_array [][]byte
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			for k := 0; k < 4; k++ {
				if static_matrix[i][k] == 2 {
					new_array[i][j] = multiply_by_2(thirdstep_array[k][j])
				} else if static_matrix[i][k] == 3 {
					new_array[i][j] = multiply_by_3(thirdstep_array[k][j])
				} else {
					new_array[i][j] = thirdstep_array[k][j]
				}
			}
		}
	}
	return new_array
}

func main() {
	var dummy_matrix = [][]byte{
		{0xd4, 0xe0, 0xb8, 0x1e},
		{0xbf, 0xb4, 0x41, 0x27},
		{0x5d, 0x52, 0x11, 0x98},
		{0x30, 0xae, 0xf1, 0xe5},
	}

	data := Aes_decrypt_scratch_mixcolumn(dummy_matrix)
	fmt.Println(data)

}
