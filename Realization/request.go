package Realization

import (
	"mycopaka/iface"
)

type Request struct{
	conn   iface.IConection
	buf    []byte
}

func NewRequest(conn iface.IConection,buf []byte)iface.IRequest{
	return &Request{
		conn: conn,
		buf:  buf,
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

func(r *Request)GetMsg()[]byte{
	return r.buf
}

func(r *Request)GetMsgID()uint32{
	return 1
}

func(r *Request)GetMsgLen()int{
	return len(r.buf)
}
