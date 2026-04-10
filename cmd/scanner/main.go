package main

import (
	"fmt"
	"seven.com/sniffer/pkg/structs"
	"seven.com/sniffer/internal/scan"
)


func main() {
	packets := make(chan * structs.Packet)
	statuses := make(chan * structs.IpStatus)

	ip := "192.0.0.1"

	for port:= range 5 {
		go scan.Dial(&structs.Address{Host: ip, Port: string(port)}, packets, statuses, 60)
	}

	for p := range packets {
		fmt.Println(p)
	}
}