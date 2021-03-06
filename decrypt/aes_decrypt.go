package decrypt

// refference: https://www.hrpub.org/download/20171130/CSIT2-13510193.pdf
// refference: https://www.commonlounge.com/discussion/e32fdd267aaa4240a4464723bc74d0a5
var key = []byte{
	0x2b, 0x28, 0xab, 0x09,
	0x7e, 0xae, 0xf7, 0xcf,
	0x15, 0xd2, 0x15, 0x4f,
	0x16, 0xa6, 0x88, 0x3c,
}
var expandedKey = expandKey(key)
var ExportedKey = expandKey(key)

var sbox = [256]byte{
	//0     1    2      3     4    5     6     7      8    9     A      B    C     D     E     F
	0x63, 0x7c, 0x77, 0x7b, 0xf2, 0x6b, 0x6f, 0xc5, 0x30, 0x01, 0x67, 0x2b, 0xfe, 0xd7, 0xab, 0x76, //0
	0xca, 0x82, 0xc9, 0x7d, 0xfa, 0x59, 0x47, 0xf0, 0xad, 0xd4, 0xa2, 0xaf, 0x9c, 0xa4, 0x72, 0xc0, //1
	0xb7, 0xfd, 0x93, 0x26, 0x36, 0x3f, 0xf7, 0xcc, 0x34, 0xa5, 0xe5, 0xf1, 0x71, 0xd8, 0x31, 0x15, //2
	0x04, 0xc7, 0x23, 0xc3, 0x18, 0x96, 0x05, 0x9a, 0x07, 0x12, 0x80, 0xe2, 0xeb, 0x27, 0xb2, 0x75, //3
	0x09, 0x83, 0x2c, 0x1a, 0x1b, 0x6e, 0x5a, 0xa0, 0x52, 0x3b, 0xd6, 0xb3, 0x29, 0xe3, 0x2f, 0x84, //4
	0x53, 0xd1, 0x00, 0xed, 0x20, 0xfc, 0xb1, 0x5b, 0x6a, 0xcb, 0xbe, 0x39, 0x4a, 0x4c, 0x58, 0xcf, //5
	0xd0, 0xef, 0xaa, 0xfb, 0x43, 0x4d, 0x33, 0x85, 0x45, 0xf9, 0x02, 0x7f, 0x50, 0x3c, 0x9f, 0xa8, //6
	0x51, 0xa3, 0x40, 0x8f, 0x92, 0x9d, 0x38, 0xf5, 0xbc, 0xb6, 0xda, 0x21, 0x10, 0xff, 0xf3, 0xd2, //7
	0xcd, 0x0c, 0x13, 0xec, 0x5f, 0x97, 0x44, 0x17, 0xc4, 0xa7, 0x7e, 0x3d, 0x64, 0x5d, 0x19, 0x73, //8
	0x60, 0x81, 0x4f, 0xdc, 0x22, 0x2a, 0x90, 0x88, 0x46, 0xee, 0xb8, 0x14, 0xde, 0x5e, 0x0b, 0xdb, //9
	0xe0, 0x32, 0x3a, 0x0a, 0x49, 0x06, 0x24, 0x5c, 0xc2, 0xd3, 0xac, 0x62, 0x91, 0x95, 0xe4, 0x79, //A
	0xe7, 0xc8, 0x37, 0x6d, 0x8d, 0xd5, 0x4e, 0xa9, 0x6c, 0x56, 0xf4, 0xea, 0x65, 0x7a, 0xae, 0x08, //B
	0xba, 0x78, 0x25, 0x2e, 0x1c, 0xa6, 0xb4, 0xc6, 0xe8, 0xdd, 0x74, 0x1f, 0x4b, 0xbd, 0x8b, 0x8a, //C
	0x70, 0x3e, 0xb5, 0x66, 0x48, 0x03, 0xf6, 0x0e, 0x61, 0x35, 0x57, 0xb9, 0x86, 0xc1, 0x1d, 0x9e, //D
	0xe1, 0xf8, 0x98, 0x11, 0x69, 0xd9, 0x8e, 0x94, 0x9b, 0x1e, 0x87, 0xe9, 0xce, 0x55, 0x28, 0xdf, //E
	0x8c, 0xa1, 0x89, 0x0d, 0xbf, 0xe6, 0x42, 0x68, 0x41, 0x99, 0x2d, 0x0f, 0xb0, 0x54, 0xbb, 0x16} //F

func Aes_decrypt_scratch_subbytes(ready_to_change_1_dim [][]byte) []byte {
	// hex value sbox

	// ubah jadi 1 dimensi
	var firststep_array = convertToOneDimension(ready_to_change_1_dim)

	// Fungsi subbytes
	var new_array = make([]byte, 16)
	i := 0
	for i < 16 {
		// refference: https://stackoverflow.com/questions/8032170/how-to-assign-string-to-bytes-array
		new_array[i] = subByte(firststep_array[i])
		i++
	}
	return new_array
}

func convertToOneDimension(ready_to_change_1_dim [][]byte) []byte {
	// ubah jadi 1 dimensi
	firststep_array := make([]byte, 16)
	count_convert := 0
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			firststep_array[count_convert] = ready_to_change_1_dim[i][j]
			count_convert++
		}
	}

	return firststep_array
}

func subByte(b byte) byte {
	return sbox[b]
}

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
	var the_result byte = multiply_by_2(the_value) ^ the_value
	return the_result
}

func Aes_decrypt_scratch_mixcolumn(thirdstep_array [][]byte) [][]byte {
	// reference: https://gist.github.com/vwxyzjn/bcac5f97b5abb7708773a28b82a809b4
	// reference: https://blog.tclaverie.eu/posts/understanding-golangs-aes-implementation-t-tables/
	// multiply matrix concept
	// 02 03 01 01
	// 01 02 03 01
	// 01 01 02 03
	// 03 01 01 02

	var new_array = make([][]byte, 4)
	for i := 0; i < 4; i++ {
		new_array[i] = make([]byte, 4)
	}

	new_array[0][0] = multiply_by_2(thirdstep_array[0][0])
	new_array[1][0] = multiply_by_3(thirdstep_array[1][0])
	new_array[2][0] = thirdstep_array[2][0]
	new_array[3][0] = thirdstep_array[3][0]

	new_array[0][1] = thirdstep_array[0][1]
	new_array[1][1] = multiply_by_2(thirdstep_array[1][1])
	new_array[2][1] = multiply_by_3(thirdstep_array[2][1])
	new_array[3][1] = thirdstep_array[3][1]

	new_array[0][2] = thirdstep_array[0][2]
	new_array[1][2] = thirdstep_array[1][2]
	new_array[2][2] = multiply_by_2(thirdstep_array[2][2])
	new_array[3][2] = multiply_by_3(thirdstep_array[3][2])

	new_array[0][3] = multiply_by_3(thirdstep_array[0][3])
	new_array[1][3] = thirdstep_array[1][3]
	new_array[2][3] = thirdstep_array[2][3]
	new_array[3][3] = multiply_by_2(thirdstep_array[3][3])
	// for i := 0; i < 4; i++ {
	// 	for j := 0; j < 4; j++ {
	// 		for k := 0; k < 4; k++ {
	// 			if static_matrix[i][k] == 2 {
	// 				new_array[i][j] = multiply_by_2(thirdstep_array[k][j])
	// 			} else if static_matrix[i][k] == 3 {
	// 				new_array[i][j] = multiply_by_3(thirdstep_array[k][j])
	// 			} else {
	// 				new_array[i][j] = thirdstep_array[k][j]
	// 			}
	// 		}
	// 	}
	// }
	return new_array
}

func expandKey(key []byte) [][]byte {
	var rcon byte = 1
	var expandedKey = make([][]byte, 4)
	for i := 0; i < 4; i++ {
		expandedKey[i] = make([]byte, 44)
	}
	i := 0
	for j := 0; j < 16; j += 4 {
		expandedKey[i][0] = key[j]
		expandedKey[i][1] = key[j+1]
		expandedKey[i][2] = key[j+2]
		expandedKey[i][3] = key[j+3]
		i++
	}
	for i := 4; i <= 40; i += 4 {
		for j := 0; j < 4; j++ {
			if j == 0 {
				// rotate
				temp := expandedKey[0][i-1]
				expandedKey[0][i+j] = expandedKey[1][i+j-1]
				expandedKey[1][i+j] = expandedKey[2][i+j-1]
				expandedKey[2][i+j] = expandedKey[3][i+j-1]
				expandedKey[3][i+j] = temp
				// sub byte
				expandedKey[0][i+j] = subByte(expandedKey[0][i+j])
				expandedKey[1][i+j] = subByte(expandedKey[1][i+j])
				expandedKey[2][i+j] = subByte(expandedKey[2][i+j])
				expandedKey[3][i+j] = subByte(expandedKey[3][i+j])
				// xor dengan i+j-4 dan rcon i+j
				expandedKey[0][i+j] = expandedKey[0][i+j] ^ expandedKey[0][i+j-4] ^ rcon
				expandedKey[1][i+j] = expandedKey[1][i+j] ^ expandedKey[1][i+j-4]
				expandedKey[2][i+j] = expandedKey[2][i+j] ^ expandedKey[2][i+j-4]
				expandedKey[3][i+j] = expandedKey[3][i+j] ^ expandedKey[3][i+j-4]
				// increment rcon
				rcon <<= 1
			} else {
				// xor dengan i-4
				expandedKey[0][i+j] = expandedKey[0][i+j-1] ^ expandedKey[0][i+j-4]
				expandedKey[1][i+j] = expandedKey[1][i+j-1] ^ expandedKey[1][i+j-4]
				expandedKey[2][i+j] = expandedKey[2][i+j-1] ^ expandedKey[2][i+j-4]
				expandedKey[3][i+j] = expandedKey[3][i+j-1] ^ expandedKey[3][i+j-4]
			}
		}
	}
	return expandedKey
}

func AddRoundKey(block [][]byte, round int) [][]byte {
	var result = make([][]byte, 4)
	for i := 0; i < 4; i++ {
		result[i] = make([]byte, 4)
	}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			result[i][j] = block[i][j] ^ expandedKey[i][j+(4*(round))]
		}
	}
	return result
}

func xorBlock(block1, block2 [][]byte) [][]byte {
	var result = make([][]byte, 4)
	for i := 0; i < 4; i++ {
		result[i] = make([]byte, 4)
	}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			result[i][j] = block1[i][j] ^ block2[i][j]
		}
	}
	return result
}

func xorSlice(slice1, slice2 []byte) []byte {
	var result = make([]byte, 16)
	for i := 0; i < len(slice1); i++ {
		result[i] = slice1[i] ^ slice2[i]
	}
	return result
}

func incrementIV(iv *[]byte) {
	// increment
	idx := 0
	for {
		if (*iv)[idx] == 0xff {
			(*iv)[idx]++
			idx++
		} else {
			(*iv)[idx]++
			break
		}
	}
}

func Encrypt(plaintext, iv []byte) []byte {
	ciphertext := make([]byte, 2048)
	// tambahin iv
	ciphertext = append(ciphertext, iv...)
	// if len > 16
	for len(plaintext) >= 16 {
		// ambil 16 awal, enkrip, append
		b := plaintext[:16]
		ciphertext = append(ciphertext, Encrypt16Byte(b, iv)...)
		incrementIV(&iv)
		plaintext = plaintext[16:]
	}
	// if len > 0
	if len(plaintext) > 0 {
		// ambil sebanyak len, enkrip, append
		b := plaintext[:]
		ciphertext = append(ciphertext, Encrypt16Byte(b, iv)...)
	}
	return ciphertext
}

func Decrypt(ciphertext []byte) []byte {
	plaintext := make([]byte, 2048)
	// ambil iv
	iv := ciphertext[:16]
	ciphertext = ciphertext[16:]
	// if len > 16
	for len(ciphertext) > 16 {
		// ambil 16 awal, enkrip, append
		b := ciphertext[:16]
		plaintext = append(plaintext, Decrypt16Byte(b, iv)...)
		incrementIV(&iv)
	}
	// if len > 0
	if len(ciphertext) > 16 {
		// ambil sebanyak len, enkrip, append
		b := ciphertext[:]
		plaintext = append(plaintext, Decrypt16Byte(b, iv)...)
	}
	// udah
	return plaintext
}

func Encrypt16Byte(plaintext, iv []byte) []byte {
	// buat cipher
	blockPlaintext := changeToMultiSlice(plaintext)
	ciphertext := newCipher(iv)

	// xor cipher yang dibuat dengan plaintext
	ciphertext = xorBlock(ciphertext, blockPlaintext)
	finalCiphertext := changeToSingleSlice(ciphertext)
	// add iv ke hasil xor
	result := append(iv, finalCiphertext...)
	// balikin text
	return result
}

func Decrypt16Byte(ciphertext []byte, iv []byte) []byte {
	// ambil iv
	cipherPart := ciphertext[16:]

	// buat cipher
	blockCiphertext := newCipher(iv)

	// xor cipher dan ciphertext dari index blocksize
	newCiphertext := changeToSingleSlice(blockCiphertext)
	plaintext := xorSlice(cipherPart, newCiphertext)

	// balikin text
	return plaintext
}

func newCipher(iv []byte) [][]byte {
	ciphertext := changeToMultiSlice(iv)

	// add round key awal
	ciphertext = AddRoundKey(ciphertext, 0)

	// for 9 kali
	for i := 0; i < 9; i++ {
		// subbyte
		// shiftrow
		ciphertext = Aes_decrypt_scratch_shiftrows(Aes_decrypt_scratch_subbytes(ciphertext))
		// mix column
		ciphertext = Aes_decrypt_scratch_mixcolumn(ciphertext)
		// add round key
		ciphertext = AddRoundKey(ciphertext, i+1)
	}

	// terakhir
	// sub byte
	// shift row
	ciphertext = Aes_decrypt_scratch_shiftrows(Aes_decrypt_scratch_subbytes(ciphertext))
	// add round keyD
	ciphertext = AddRoundKey(ciphertext, 10)

	// return hasilnya
	return ciphertext
}

func changeToMultiSlice(block []byte) [][]byte {
	result := make([][]byte, 4)
	lenBlock := len(block)
	row := lenBlock / 4
	col := lenBlock % 4
	for i := 0; i < 4; i++ {
		result[i] = make([]byte, 4)
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			result[i][j] = block[4*i+j]
		}
	}
	return result
}

func changeToSingleSlice(block [][]byte) []byte {
	var result = make([]byte, 16)
	lenBlock := len(block)
	for i := 0; i < lenBlock; i++ {
		for j := 0; j < len(block[i]); j++ {
			result[4*i+j] = block[i][j]
		}
	}
	return result
}
