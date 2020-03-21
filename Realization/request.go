package Realization

import (
	"mycopaka/iface"
)

//每个request都是一个请求
//每个请求封装后都交给路由系统处理
type Request struct{
	conn   iface.IConection
	msg    iface.IMessage
}

func NewRequest(conn iface.IConection,msg iface.IMessage)iface.IRequest{
	return &Request{
		conn: conn,
		msg:  msg,
	}
}

func(r *Request)GetConnection()iface.IConection{
	return r.conn
}

func(r *Request)GetConnAddrString()string{
	return r.conn.GetAddr().String()
}

func(r *Request)GetConnID()uint32{
	return r.GetConnection().GetID()
}

func(r *Request)GetMsg()iface.IMessage{
	return r.msg
}

func(r *Request)GetMsgID()uint32{
	return r.msg.GetMsgID()
}

func(r *Request)GetMsgLen()uint32{
	return r.msg.GetMsgLen()
}
