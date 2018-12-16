package main

import(
	"net"
	"fmt"
	"runtime"
)

//type Data struct{
//	C string
//	D *net.UDPAddr
//}

type ConsumeData interface {
	Consume([]byte) error
}

func UdpServer(host string, consumeData ConsumeData){
	runtime.GOMAXPROCS(runtime.NumCPU())
	addr, err := net.ResolveUDPAddr("udp", host)  // Niconiconi～
	if err != nil{
		fmt.Println("!ERR:", err)
		return
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil{
		fmt.Println("!ERR:", err)
		return
	}
	defer conn.Close()

	c := make(chan []byte, 10240)  // 带缓冲的通道
	fmt.Println("Listen at " + addr.String())
	for {
		data := make([]byte, 10240)
		_, _, err := conn.ReadFromUDP(data)
		if err != nil {
			fmt.Println("!ERR:", err)
			continue
		}
		go sets(data, c)  // 通过通道来接收数据
		select{
		case k:= <-c:
			consumeData.Consume(k)
		default:
		}
	}
}

func sets(data []byte, c chan []byte){
	c <- data
}