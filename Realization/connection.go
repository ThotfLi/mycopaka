package Realization

import (
	"fmt"
	"io"
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

	pk := Package{}

	//持续接受客户端消息
	for {
		buf := make([]byte,pk.GetHeadLen())
		n,err := io.ReadFull(c.GetConn(),buf)
		if err != nil || n == 0 {
			fmt.Println("[ERROR]recv faild,conn id:", c.Id)
			c.SignOut <- true
			break
		}

		//解包拿到msg.Data
		newMsg := pk.UnPack(buf)
		data := make([]byte,newMsg.GetMsgLen())
		_,err = io.ReadFull(c.GetConn(),data)
		if err != nil {
			fmt.Println("[ERROR]recv msgData faild,conn id",c.Id)
		}
		newMsg.SetData(data)
		//对从客户端接收的字节序列进行封装
		//newMsg := NewMessage(buf[:n],1)
		req := NewRequest(c,newMsg)
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

func(c *Connection)SendMsg(data []byte)(int,error){
	if n,err := c.Conn.Write(data);err != nil {
		return 0,err
	}else {
		return n, nil
	}
}