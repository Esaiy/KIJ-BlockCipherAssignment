package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

var key = []byte("1234567890123456")

func Aes_encrypt(chunk []byte) []byte {
	c, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err.Error())
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		fmt.Println(err.Error())
	}
	nonce := make([]byte, gcm.NonceSize())

	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println(err)
	}
	enc := gcm.Seal(nonce, nonce, chunk, nil)
	return enc
}

func Aes_decrypt(chunk []byte) []byte {
	c, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		fmt.Println(err)
	}

	nonceSize := gcm.NonceSize()
	if len(chunk) < nonceSize {
		fmt.Println(err)
	}

	nonce, chunk := chunk[:nonceSize], chunk[nonceSize:]
	dec, err := gcm.Open(nil, nonce, chunk, nil)
	if err != nil {
		fmt.Println(err.Error())
	}

	return dec
}
