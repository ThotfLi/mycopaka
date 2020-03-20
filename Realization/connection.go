package Realization

import (
	"fmt"
	"net"
	"mycopaka/iface"
)

type Connection struct{
	Id      uint32
	Conn    net.Conn
	SignOut chan bool  //负责发送断开当前conn及回收资源信号
	IsClose bool       //当前conn是否处于关闭状态

	RouterHandle iface.IRouterHandel
}

//创建Connection 对象
func NewConnection(conn net.Conn, id uint32,routerhandle iface.IRouterHandel)iface.IConection{
	return &Connection{
		Id:      id,
		Conn:    conn,
		SignOut: make(chan bool,1),
		IsClose: false,
		RouterHandle:routerhandle,
	}
}

func(c *Connection)Start(){
	fmt.Println("[START]Runing a Conn  id：",c.Id)

	go c.StartReader()
	go c.StartWriter()

	select {
	case <-c.SignOut:
		c.Stop()
	}
}

//从客户端消息请求进行处理
func(c *Connection)StartReader(){
	fmt.Println("[START]Connection.StartReader is running")
	fmt.Println("conn addr:",c.GetAddr().String())
	defer fmt.Println("exit conn id=",c.Id)
	for {
		buf := make([]byte,512)
		n,err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("[ERROR]recv faild,conn id:", c.Id)
			c.SignOut <- true
			break
		}
		fmt.Println(string(buf[:n]))
		req := NewRequest(c,buf[:n])
		//根据路由选择运行router
		go c.RouterHandle.RunRouter(req)



	}
}

//向客户端发送响应消息
func(c *Connection)StartWriter(){}

func(c *Connection)Stop(){
	if c.IsClose == true{return}

	fmt.Println("[STOP] conn:",c.GetAddr())
	c.IsClose = true
	c.Conn.Close()
	close(c.SignOut)
}

func(c *Connection)GetConn()net.Conn{
	return c.Conn
}

func(c *Connection)GetAddr()net.Addr{
	return c.Conn.RemoteAddr()
}

func(c *Connection)GetID()uint32{
	return c.Id
}