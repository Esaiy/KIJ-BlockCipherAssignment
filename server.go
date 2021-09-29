package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"sync"

	"kij-block-cipher/encrypt"

	ipc "github.com/james-barrow/golang-ipc"
)

func main() {
	var wg sync.WaitGroup
	config := &ipc.ServerConfig{Encryption: false}
	sc, err := ipc.StartServer("testtest", config)
	if err != nil {
		fmt.Println(err)
		return
	}
	wg.Add(1)
	go func() {

		for {
			m, err := sc.Read()
			if err == nil {
				if m.MsgType > 0 {
					fmt.Println("Server recieved: " + string(m.Data))
					go sendFile(sc)
				}
			} else {
				fmt.Println("Server error")
				fmt.Println(err)
				break
			}
		}
	}()
	wg.Wait()
}

func sendFile(sc *ipc.Server) {
	file, err := os.Open("./dataset/Fatin - Aku memilih setia.mp3")
	if err != nil {
		log.Println("something broke :", err.Error())
	}
	defer file.Close()
	fi, err := file.Stat()
	sc.Write(69, []byte(strconv.Itoa(int(fi.Size()))))

	const maxChunk = 2048
	buffer := make([]byte, maxChunk)

	for {
		_, err := file.Read(buffer)
		if err != nil {
			if err != io.EOF {
				log.Println("reading broke :", err.Error())
			}
			sc.Write(71, []byte{})
			break
		}
		buffer = encrypt.Aes_encrypt(buffer)
		sc.Write(70, buffer)
	}
	return
}
