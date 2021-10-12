package decrypt

// refference: https://www.hrpub.org/download/20171130/CSIT2-13510193.pdf
// refference: https://www.commonlounge.com/discussion/e32fdd267aaa4240a4464723bc74d0a5
var key = []byte("1234567890123456")
var expandedKey = expandKey(key)

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

func Aes_decrypt_scratch_subbytes(ready_to_change_1_dim [4][4]byte) []byte {
	// hex value sbox

	// ubah jadi 1 dimensi
	var firststep_array = convertToOneDimension(ready_to_change_1_dim)

	// Fungsi subbytes
	var new_array []byte
	i := 0
	for i < 16 {
		// refference: https://stackoverflow.com/questions/8032170/how-to-assign-string-to-bytes-array
		new_array[i] = subByte(firststep_array[i])
		i++
	}
	return new_array
}

func convertToOneDimension(ready_to_change_1_dim [4][4]byte) []byte {
	// ubah jadi 1 dimensi
	var firststep_array []byte
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
	var secondstep_array [][]byte
	count_convert := 0
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			secondstep_array[i][j] = ready_to_change_2_Dim[count_convert]
			count_convert++
		}
	}

	// fungsi shift rows
	var new_array [][]byte
	i := 0
	j := 0
	// shifting x kali = shift pada baris ke x
	for i < 4 {
		for j < 4 {
			// shifting 0x
			if i == 0 {
				new_array[i][j] = secondstep_array[i][j]
			}
			// shifting 1x
			if i == 1 {
				// moving [1,0] ->[1,3]
				if j == 0 {
					new_array[i][3] = secondstep_array[i][j]
				} else {
					new_array[i][j-1] = secondstep_array[i][j]
				}
			}
			// shifting 2x
			if i == 2 {
				// moving [2,0] -> [2,2]; [2,1] -> [2,3]
				if j == 0 {
					new_array[i][2] = secondstep_array[i][j]
				} else if j == 1 {
					new_array[i][3] = secondstep_array[i][j]
				} else {
					new_array[i][j-2] = secondstep_array[i][j]
				}
			}
			// shifting 3x
			if i == 3 {
				// moving [3,0] -> [3,1]; [3,1] -> [3,2]; [3,2] -> [3,3];
				if j == 0 {
					new_array[i][1] = secondstep_array[i][j]
				} else if j == 1 {
					new_array[i][2] = secondstep_array[i][j]
				} else if j == 2 {
					new_array[i][3] = secondstep_array[i][j]
				} else {
					new_array[i][j-3] = secondstep_array[i][j]
				}
			}
			j++

		}
		i++
	}
	return new_array
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
				new_array[i][j] = static_matrix[i][k] * thirdstep_array[k][j]
			}
		}
	}
	return new_array
}

func expandKey(key []byte) [][]byte {
	var rcon byte = 1
	var expandedKey [][]byte
	i := 0
	for j := 0; j < 16; j += 4 {
		expandedKey[i][0] = key[j]
		expandedKey[i][1] = key[j+1]
		expandedKey[i][2] = key[j+2]
		expandedKey[i][3] = key[j+3]
		i++
	}
	for i := 0; i < 10; i++ {
		for j := 0; j < 4; j++ {
			if j%4 == 0 {
				// rotate
				temp := expandedKey[0][i]
				expandedKey[0][i] = expandedKey[1][i]
				expandedKey[1][i] = expandedKey[2][i]
				expandedKey[2][i] = expandedKey[3][i]
				expandedKey[3][i] = temp
				// sub byte
				expandedKey[0][i] = subByte(expandedKey[0][i])
				expandedKey[1][i] = subByte(expandedKey[1][i])
				expandedKey[2][i] = subByte(expandedKey[2][i])
				expandedKey[3][i] = subByte(expandedKey[3][i])
				// xor dengan i-4 dan rcon i
				expandedKey[0][i] = expandedKey[0][i] ^ expandedKey[0][i-4] ^ rcon
				expandedKey[1][i] = expandedKey[1][i] ^ expandedKey[1][i-4]
				expandedKey[2][i] = expandedKey[2][i] ^ expandedKey[2][i-4]
				expandedKey[3][i] = expandedKey[3][i] ^ expandedKey[3][i-4]
				// increment rcon
				rcon <<= 1
			} else {
				// xor dengan i-4
				expandedKey[0][i] = expandedKey[0][i] ^ expandedKey[0][i-4]
				expandedKey[1][i] = expandedKey[1][i] ^ expandedKey[1][i-4]
				expandedKey[2][i] = expandedKey[2][i] ^ expandedKey[2][i-4]
				expandedKey[3][i] = expandedKey[3][i] ^ expandedKey[3][i-4]
			}
		}
	}
	return expandedKey
}

func addRoundKey(block [][]byte, round int) [][]byte {
	var result [][]byte
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			result[i][j] = block[i][j] ^ expandedKey[i+(4*(round-1))][j]
		}
	}
	return result
}

func rotateRow(row []byte, shift int) []byte {
	var temp []byte = make([]byte, 4, 4)
	for i := 0; i < 4; i++ {
		idx := (i + (4 - shift)) % 4
		temp[idx] = row[i]
	}
	return temp
}
