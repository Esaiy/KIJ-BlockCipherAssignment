package main

import "fmt"

func Aes_decrypt_scratch_shiftrows(ready_to_change_2_Dim []byte) [][]byte {
	// https://stackoverflow.com/questions/29442710/how-to-shift-byte-array-with-golang}
	// ubah jadi 2x2 array lagi
	var secondstep_array = make([][]byte, 4)
	for i := 0; i < 4; i++ {
		secondstep_array[i] = make([]byte, 4)
	}
	count_convert := 0
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			secondstep_array[i][j] = ready_to_change_2_Dim[count_convert]
			count_convert++
		}
	}

	// fungsi shift rows
	var new_array = make([][]byte, 4)
	for i := 0; i < 4; i++ {
		new_array[i] = make([]byte, 4)
	}
	// emergency loop stuck
	for i := 0; i < 4; i++ {
		new_array[0][i] = secondstep_array[0][i]
	}
	new_array[1][0] = secondstep_array[1][1]
	new_array[1][1] = secondstep_array[1][2]
	new_array[1][2] = secondstep_array[1][3]
	new_array[1][3] = secondstep_array[1][0]

	new_array[2][0] = secondstep_array[2][2]
	new_array[2][1] = secondstep_array[2][3]
	new_array[2][2] = secondstep_array[2][0]
	new_array[2][3] = secondstep_array[2][1]

	new_array[3][0] = secondstep_array[3][3]
	new_array[3][1] = secondstep_array[3][0]
	new_array[3][2] = secondstep_array[3][1]
	new_array[3][3] = secondstep_array[3][2]
	// shifting x kali = shift pada baris ke x
	// for i < 4 {
	// 	for j < 4 {
	// shifting 0x
	// if i == 0 {
	// 	new_array[i][j] = secondstep_array[i][j]
	// }
	// shifting 1x
	// if i == 1 {
	// 	// moving [1,0] ->[1,3]
	// 	if j == 0 {
	// 		new_array[i][3] = secondstep_array[i][j]
	// 	} else {
	// 		new_array[i][j-1] = secondstep_array[i][j]
	// 	}
	// }
	// shifting 2x
	// if i == 2 {
	// 	// moving [2,0] -> [2,2]; [2,1] -> [2,3]
	// 	if j == 0 {
	// 		new_array[i][2] = secondstep_array[i][j]
	// 	} else if j == 1 {
	// 		new_array[i][3] = secondstep_array[i][j]
	// 	} else {
	// 		new_array[i][j-2] = secondstep_array[i][j]
	// 	}
	// }
	// shifting 3x
	// if i == 3 {
	// 	// moving [3,0] -> [3,1]; [3,1] -> [3,2]; [3,2] -> [3,3];
	// 	if j == 0 {
	// 		new_array[i][1] = secondstep_array[i][j]
	// 	} else if j == 1 {
	// 		new_array[i][2] = secondstep_array[i][j]
	// 	} else if j == 2 {
	// 		new_array[i][3] = secondstep_array[i][j]
	// 	} else {
	// 		new_array[i][j-3] = secondstep_array[i][j]
	// 	}
	// }
	// 		j++

	// 	}
	// 	i++
	// }
	return new_array
}

func main() {
	var dummy_matrix = []byte{
		0xd4, 0xe0, 0xb8, 0x1e,
		0x27, 0xbf, 0xb4, 0x41,
		0x11, 0x98, 0x5d, 0x52,
		0xae, 0xf1, 0xe5, 0x30,
	}

	data := Aes_decrypt_scratch_shiftrows(dummy_matrix)
	fmt.Println(data)

}
