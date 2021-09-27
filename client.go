package main

import (
	"fmt"

	ipc "github.com/james-barrow/golang-ipc"
)

func main() {
	config := &ipc.ClientConfig{Encryption: false}

	cc, err := ipc.StartClient("testtest", config)
	if err != nil {
		fmt.Println(err)
		return
	}

	go func() {

		for {
			m, err := cc.Read()

			if err != nil {
				// An error is only returned if the recieved channel has been closed,
				//so you know the connection has either been intentionally closed or has timmed out waiting to connect/re-connect.
				break
			}

			if m.MsgType == -1 { // message type -1 is status change
				//fmt.Println("Status: " + m.Status)
			}

			if m.MsgType == -2 { // message type -2 is an error, these won't automatically cause the recieve channel to close.
				fmt.Println("Error: " + err.Error())
			}

			if m.MsgType > 0 { // all message types above 0 have been recieved over the connection

				// fmt.Println(" Message type: ", m.MsgType)
				fmt.Println("Client recieved: " + string(m.Data))
			}
			//}
		}

	}()

	clientSend(cc)
}

func clientSend(cc *ipc.Client) {
	var input string
	for {
		fmt.Scanln(&input)
		_ = cc.Write(14, []byte(input))
	}

}
