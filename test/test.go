package main

// func main() {
// 	key := []byte("1234567890123456")
// 	file, err := os.Open("../dataset/plrabn12.txt")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// 	defer file.Close()
// 	// fileOut, err := os.OpenFile("../dest/something", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
// 	// if err != nil {
// 	// 	fmt.Println(err.Error())
// 	// }

// 	c, err := aes.NewCipher(key)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}

// 	gcm, err := cipher.NewGCM(c)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// 	nonce := make([]byte, gcm.NonceSize())
// 	fmt.Println(gcm.NonceSize(), gcm.NonceSize(), gcm.NonceSize())

// 	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
// 		fmt.Println(err)
// 	}
// 	text := []byte("hello something")
// 	enc := gcm.Seal(nonce, nonce, text, nil)
// 	err = ioutil.WriteFile("myfile.data", enc, 0777)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// }
func main() {
	// test xor
	// var rcon int = 1
	// var random int = 0x10
	// for i := 0; i < 10; i++ {
	// 	// fmt.Printf("%x\n", rcon)
	// 	fmt.Printf("%x\n", random^rcon)
	// 	rcon <<= 1
	// }
	// // for i := 0; i <= 16; i++ {
	// // 	fmt.Printf("%x\n", i)
	// // }
	// row := []int{0, 1, 2, 3}
	// shift := 1
	// var temp []int = make([]int, 4, 4)
	// for i := 0; i < 4; i++ {
	// 	idx := (i + (4 - shift)) % 4
	// 	temp[idx] = row[i]
	// 	// 0 1 2 3
	// 	// 1 2 3 0
	// 	// 2 3 0 1
	// 	// 3 0 1 2
	// }
	// for i := 0; i < 4; i++ {
	// 	fmt.Println(temp[i])
	// }
	var key = []byte("1234567890123456")
	var newKey [][]byte
	i := 0
	for j := 0; j < 4; j += 4 {
		newKey[i][0] = key[j]
		newKey[i][1] = key[j+1]
		newKey[i][2] = key[j+2]
		newKey[i][3] = key[j+3]
		i++
	}

}
