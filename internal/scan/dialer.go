package scan

import (
	"bufio"
	"fmt"
	"net"
	"time"
	"seven.com/sniffer/pkg/structs"
)

func Dial(ad * structs.Address, packets chan * structs.Packet, statuses chan * structs.IpStatus, timeoutSeconds uint8 ) {
	for{
		conn, err := net.Dial("tcp", fmt.Sprintf("%v:%v", ad.Host, ad.Port))

		if err != nil{
			statuses <- &structs.IpStatus{Ad:ad, IsOnline: false, Err: fmt.Sprintf("Unable to ping %v:%v. Error: %v", ad.Host, ad.Port, err)}
			time.Sleep(1 * time.Second) //chill
			continue
		}

		statuses <- &structs.IpStatus{Ad:ad, IsOnline: true}
		reader := bufio.NewReader(conn)
		for {
			
		}
	}
}