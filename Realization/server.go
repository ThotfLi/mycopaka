package Realization

import (
	"mycopaka/iface"
	"fmt"
	"net"
	"time"
)

type Server struct{
	Host string
	Port int
	Name string
	RouterHandle iface.IRouterHandel
}
func NewServer(host string,port int,name string)iface.IServer{
	return &Server{
		Host:host,
		Port:port,
		Name:name,
		RouterHandle:NewRouterHandle(),
	}
}

func(s *Server)Start(){
	fmt.Printf("[START]Server Name: IP:%s:%d is Running\n",s.Host,s.Port)

	listener,err := net.Listen("tcp",fmt.Sprintf("%s:%d",s.Host,s.Port))
	if err != nil{
		fmt.Println("[ERROR]Server run err:",err)
	}

	//给conn的自增id
	AuToId := uint32(0)

	//循环监听客户端连接
	for {
		conn,err := listener.Accept()
		if err != nil{
			fmt.Println("[ERROR]Conn open failed")
		}

		//创建一个conn对象
		connectionObject := NewConnection(conn,AuToId,s.RouterHandle)
		//启动conn任务
		go connectionObject.Start()
		AuToId += 1
	}
}

func(s *Server)Stop(){}

func(s *Server)Serve(){
	//开始运行服务器
	s.Start()

	//可以添加一些服务器开启之前、开启之后的一些服务

	time.Sleep(20*time.Second)
}

func(s *Server)AddRouter(id uint32,router iface.IRouter){
	s.RouterHandle.AddRouter(id,router)
}

func(s *Server)DelRouter(id uint32){
	s.RouterHandle.DelRouter(id)
}