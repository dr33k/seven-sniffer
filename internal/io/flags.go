package io

import (
	"errors"
	"flag"
	"fmt"
	"net"
	"regexp"
	"strconv"
	"strings"

	"seven.com/sniffer/pkg/structs"
)


func ParseFlags() (*structs.Input, error){
	hostPtr := flag.String("host", "scanme.nmap.org", "The IPv4, IPv6 or DNS resolvable host name to sniff")
	portsPtr:= flag.String("ports", "80", "Comma-Separated Port numbers of the given host to sniff")
	protocolPtr:= flag.String("ptcl", "tcp", "Indicates either the TCP or UDP Protocol")

	flag.Parse()

	//Validations and transformations
	err := validateHost(hostPtr)
	if err != nil {
		return nil, err
	}

	err = validateProtocol(protocolPtr)
	if err != nil{
		return nil, err
	}

	portsInt, err := transformPorts(portsPtr)
	if err != nil{
		return nil, err
	}

	return &structs.Input{
		Host: hostPtr,
		Protocol: protocolPtr,
		Ports: portsInt,
	}, nil
}

func validateHost(value *string) error{
	if ip := net.ParseIP(*value); ip != nil{
		return nil
	} else if _, dnserr:=net.LookupHost(*value) ; dnserr != nil{
		return fmt.Errorf("Host '%s' is neither a valid IP Host nor DNS resolvable", *value)
	}

	return nil
}

func validateProtocol(value *string) error{
	regex, err := regexp.Compile("(^tcp$|^udp$)")

	if err == nil{
		return err;
	}
	if !regex.MatchString(*value){
		return fmt.Errorf("Invalid protocol: %s; Must be 'tcp' or 'udp'", *value)
	}
	return nil
}

func transformPorts(portPtr *string) ([]int, error) {
	if len(*portPtr) == 0{
		return nil, errors.New("No ports provided")
	}

	portsStr:= strings.Split(*portPtr, ",")
	portsInt:= make([]int, 0)
	for _ , p := range portsStr{
		p = strings.TrimSpace(p)
		i, err := strconv.Atoi(p)
		if err != nil{
			return nil, fmt.Errorf("Unable to parse port: %v. Must be an integer", p)
		}
		portsInt = append(portsInt, i)
	}

	return portsInt, nil
}