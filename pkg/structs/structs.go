package structs

type Input struct{
	Host *string
	Ports []int
}

type Address struct{
	Host *string
	Port int
}

type Packet struct {
	Ch chan bool
	IP *string
}

type IpStatus struct{
	Ad *Address
	IsOnline bool
	Err string
}
