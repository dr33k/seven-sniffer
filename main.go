package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

type address struct{
	Host string
	Port string
}

type packet struct {
	Ch chan bool
	IP string
}

type ipStatus struct{
	Ad *address
	IsOnline bool
	Err string
}

func sniff(ad *address, packets chan *packet, statuses chan *ipStatus, timeoutSeconds uint8 ) {
	for{
		conn, err := net.Dial("tcp", fmt.Sprintf("%v:%v", ad.Host, ad.Port))

		if err != nil{
			statuses <- &ipStatus{Ad:ad, IsOnline: false, Err: err.Error()}
			time.Sleep(1 * time.Second) //chill
			continue
		}

		statuses <- &ipStatus{Ad:ad, IsOnline: true, Err: ""}
		reader := bufio.NewReader(conn)
		for {
			
		}
	}
}

func main() {
	packets := make(chan *packet)
	statuses := make(chan *ipStatus)

	ip := "192.0.0.1"

	for port:= range 5 {
		go sniff(&address{ip, string(port)}, packets, statuses, 60)
	}

	for p := range packets {
		fmt.Println(p)
	}
}