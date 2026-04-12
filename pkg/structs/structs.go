package structs

type Address struct{
	Host string
	Port int
}

type Packet struct {
	Ch chan bool
	IP string
}

type IpStatus struct{
	Ad *Address
	IsOnline bool
	Err string
}
