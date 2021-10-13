package main

import (
	"fmt"
	"io"
	"kij-block-cipher/pkg/aes-lib"
	"log"
	"os"
	"sync"
	"time"

	// aes "kij-block-cipher/encrypt"

	ipc "github.com/james-barrow/golang-ipc"
)

func main() {
	var wg sync.WaitGroup
	config := &ipc.ServerConfig{Encryption: false}
	sc, err := ipc.StartServer("testlib", config)
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
	file, err := os.Open("./dataset/plrabn12.txt")
	if err != nil {
		log.Println("something broke :", err.Error())
	}
	defer file.Close()
	fi, err := file.Stat()
	if err != nil {
		fmt.Println(err.Error())
	}
	sc.Write(69, []byte(fi.Name()))

	const maxChunk = 2048 - 16
	buffer := make([]byte, maxChunk)
	i := 0
	average := []int64{}
	for {
		// fmt.Println(i)
		i++
		n, err := file.Read(buffer)
		buffer := buffer[:n]
		if err != nil {
			if err != io.EOF {
				log.Println("reading broke :", err.Error())
			}
			sc.Write(71, []byte{})
			fmt.Println("done")
			var total int64 = 0
			for i := range average {
				total += average[i]
			}
			hasil := float64(total) / float64(len(average))
			fmt.Println("Average encrypt time :", hasil, "microsecond")
			break
		}
		// fmt.Println(buffer)
		start := time.Now()
		buffer = aes.Encrypt(buffer)
		duration := time.Since(start)
		average = append(average, duration.Microseconds())
		sc.Write(70, buffer)
	}
	return
}
