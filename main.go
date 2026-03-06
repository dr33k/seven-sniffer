package main

type packet struct{

}

func sniff(ip string, packets chan *packet){

}

func main(){
	packets:= make(chan *packet)
	ip:= "192.0.0.1"

	for range 5{
		go sniff(ip, packets)
	}

	for p := range packets{

	}
}