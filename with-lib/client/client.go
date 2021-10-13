package main

import (
	"fmt"
	"time"

	// aes "kij-block-cipher/encrypt"
	"kij-block-cipher/pkg/aes-lib"
	"os"

	ipc "github.com/james-barrow/golang-ipc"
)

func main() {
	config := &ipc.ClientConfig{Encryption: false}

	cc, err := ipc.StartClient("testlib", config)
	if err != nil {
		fmt.Println(err)
		return
	}

	var file *os.File
	go func(file *os.File) {
		average := []int64{}
		for {
			m, err := cc.Read()

			if err != nil {
				// An error is only returned if the recieved channel has been closed,
				//so you know the connection has either been intentionally closed or has timmed out waiting to connect/re-connect.
				break
			}

			// if m.MsgType == -1 { // message type -1 is status change
			// 	//fmt.Println("Status: " + m.Status)
			// }

			if m.MsgType == -2 { // message type -2 is an error, these won't automatically cause the recieve channel to close.
				fmt.Println("Error: " + err.Error())
			}

			if m.MsgType == 69 {
				// filename := time.Now().Unix()
				file, err = os.OpenFile("./dest/"+string(m.Data), os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					fmt.Println(err.Error())
				}
				continue
			}

			if m.MsgType == 70 {
				start := time.Now()
				data := aes.Decrypt(m.Data)
				duration := time.Since(start)
				average = append(average, duration.Microseconds())
				// data = m.Data
				file.Write(data)
				continue
			}

			if m.MsgType == 71 {
				// data := aes.Decrypt(m.Data)
				// file.Write(data)
				file.Close()
				fmt.Println("File downloaded successfully")
				var total int64 = 0
				for i := range average {
					total += average[i]
				}
				hasil := float64(total) / float64(len(average))
				fmt.Println("Average decrypt time :", hasil, "microsecond")
				continue
			}

			if m.MsgType > 0 { // all message types above 0 have been recieved over the connection
				fmt.Println("Client recieved: " + string(m.Data))
			}
		}

	}(file)

	clientSend(cc)
}

func clientSend(cc *ipc.Client) {
	var input string
	for {
		fmt.Scanln(&input)
		_ = cc.Write(14, []byte(input))
	}

}
