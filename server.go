package main

import (
	"log"
	"time"

	ipc "github.com/james-barrow/golang-ipc"
)

func main() {
	config := &ipc.ServerConfig{Encryption: false}
	sc, err := ipc.StartServer("testtest", config)
	if err != nil {
		log.Println(err)
		return
	}

	go func() {

		for {
			m, err := sc.Read()

			if err == nil {
				if m.MsgType > 0 {
					log.Println("Server recieved: "+string(m.Data)+" - Message type: ", m.MsgType)
				}

			} else {

				log.Println("Server error")
				log.Println(err)
				break
			}
		}
	}()

	serverSend(sc)
}

func serverSend(sc *ipc.Server) {

	for {

		err := sc.Write(3, []byte("Hello Client 4"))
		err = sc.Write(23, []byte("Hello Client 5"))
		err = sc.Write(65, []byte("Hello Client 6"))

		if err != nil {
			//fmt.Println(err)
		}

		time.Sleep(2 * time.Second)

	}
}

func serverSend1(sc *ipc.Server) {

	for {

		sc.Write(5, []byte("Hello Client 1"))
		sc.Write(7, []byte("Hello Client 2"))
		sc.Write(9, []byte("Hello Client 3"))

		time.Sleep(time.Second / 30)

	}

}

func serverSend2(sc *ipc.Server) {

	for {

		err := sc.Write(88, []byte("Hello Client 7"))
		err = sc.Write(99, []byte("Hello Client 8"))
		err = sc.Write(22, []byte("Hello Client 9"))

		if err != nil {
			//fmt.Println(err)
		}

		time.Sleep(time.Second / 30)

	}
}
