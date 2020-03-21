package iface

import "net"

type IConection interface{
	//开启读写任务
	Start()
	//关闭当前conn
	Stop()
	//获取conn
	GetConn()        net.Conn
	//获取addr
	GetAddr()        net.Addr
	//获取id
	GetID()          uint32
	//发送消息
	SendMsg([]byte) (int,error)
}
