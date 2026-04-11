package main

import (
	"fmt"
	"strings"
	"seven.com/sniffer/internal/scan"
	"seven.com/sniffer/pkg/structs"
)


func main() {	
	statuses := make(chan * structs.IpStatus)
	addressMap := make(map[string]bool)

	// host := "172.19.9.245"
	host := "scanme.nmap.org"
	ports := []string{"22", "53", "80", "443", "631" }
	// Todo sort ports


	for _, port:= range ports {
		go scan.Sniff(
			&structs.Address{Host: host, Port: port}, 
			statuses,
		)
	}

	//Print Statuses

	// \033[2K clears the entire line
    // \r moves the cursor to the start of that line
	const FORMAT_SPECIFIER = "\033[2K\r[%-5s]: %v\n"
	
	fmt.Printf("Scanning ports for host: %v ...\n", host)
	fmt.Printf(FORMAT_SPECIFIER, "PORT", "STATUS")

	fmtPorts := func(){
		var displayString strings.Builder
		var statusText string

		for _, port :=range  ports{
			statusText = "\033[32mONLINE\033[0m" //Red Text
			if addressMap[port] {
				statusText = "\033[31mOFFLINE\033[0m" //Green Text
			}

			//Display current status
			fmt.Fprintf(&displayString, FORMAT_SPECIFIER, port, statusText)
		}

		fmt.Print(displayString.String())

		// Move back up to the start of the block
		fmt.Printf("\033[%dA", len(ports))
	}
	
	for status := range statuses {
		addressMap[status.Ad.Port] = status.IsOnline

		fmtPorts()
	}

}

