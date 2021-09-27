package main

import (
	"fmt"
	"sync"

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
