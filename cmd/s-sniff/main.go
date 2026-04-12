package main

import (
	"fmt"
	"strings"
	"seven.com/sniffer/internal/scan"
	"seven.com/sniffer/internal/io"
	"seven.com/sniffer/pkg/structs"
	"slices"
)


func main() {	
	statuses := make(chan * structs.IpStatus)
	addressMap := make(map[int]bool)
	
	//Parse Input
	input, err := io.ParseFlags()
	if err != nil{
		fmt.Println(err)
		return
	}

	
	slices.Sort(input.Ports)


	//Poll ports
	for _, port:= range input.Ports {
		go scan.Sniff(
			&structs.Address{Host: input.Host, Port: port}, 
			statuses,
		)
	}

	//Print Statuses

	// \033[2K clears the entire line
    // \r moves the cursor to the start of that line
	const FORMAT_SPECIFIER = "\033[2K\r[%-5v]: %v\n"
	
	fmt.Printf("Scanning ports for host: %v ...\n", *input.Host)
	fmt.Printf(FORMAT_SPECIFIER, "PORT", "STATUS")

	fmtPorts := func(){
		var displayString strings.Builder
		var statusText string

		for _, port :=range  input.Ports{
			statusText = "\033[31mOFFLINE\033[0m" //Red Text
			if addressMap[port] {
				statusText = "\033[32mONLINE\033[0m"//Green Text
			}

			//Display current status
			fmt.Fprintf(&displayString, FORMAT_SPECIFIER, port, statusText)
		}

		fmt.Print(displayString.String())

		// Move back up to the start of the block
		fmt.Printf("\033[%dA", len(input.Ports))
	}
	
	for status := range statuses {
		addressMap[status.Ad.Port] = status.IsOnline

		fmtPorts()
	}

}

