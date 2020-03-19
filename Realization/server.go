package Realization

import (
	"mycopaka/iface"
	"fmt"
	"net"
)

type Server struct{
	Host string
	Port int
	Name string
}
func NewServer(host string,port int,name string)iface.IServer{
	return &Server{
		Host:host,
		Port:port,
		Name:name,
	}
}

func(s *Server)Start(){
	fmt.Printf("[START]Server Name: IP:%s:%d is Running\n",s.Host,s.Port)

	listener,err := net.Listen("tcp",fmt.Sprintf("%s:%d",s.Host,s.Port))
	if err != nil{
		fmt.Println("[ERROR]Server run err:",err)
	}
	for {
		conn,err := listener.Accept()
		if err != nil{
			fmt.Println("[ERROR]Conn open failed")
		}
		buf := make([]byte,512)
		n,err := conn.Read(buf)
		if err != nil {
			fmt.Println("recv faild")
		}
		fmt.Println(string(buf[:n]))

	}
}

func(s *Server)Stop(){}

func(s *Server)Serve(){
	//开始运行服务器
	s.Start()

	//可以添加一些服务器开启之前、开启之后的一些服务
}