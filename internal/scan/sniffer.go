package scan

import (
	"bufio"
	"fmt"
	"net"
	"time"
	"seven.com/sniffer/pkg/structs"
)

func Sniff(ad * structs.Address, statuses chan * structs.IpStatus) {
	for{
		conn, err := net.Dial("tcp", fmt.Sprintf("%v:%v", ad.Host, ad.Port),)

		//Ping IP Address
		if err != nil{
			statuses <- &structs.IpStatus{Ad:ad, IsOnline: false, Err: fmt.Sprintf("Unable to ping %v:%v. Error: %v", ad.Host, ad.Port, err)}
			time.Sleep(1 * time.Second) //chill
			continue
		}

		statuses <- &structs.IpStatus{Ad:ad, IsOnline: true}

		//Read TCP Packets
		reader := bufio.NewReader(conn)
		for {
			reader.Buffered()

			//Write TCP Packets to File 
		}
	}
}