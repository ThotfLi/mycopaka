package iface

//请求对象
//每次从客户端接收到信息都要封装成请求对象发送给路由
type IRequest interface{
	//获取当前请求的connection
	GetConnection()IConection
	//获取conn的addr
	GetConnAddrString()string
	//获取connID
	GetConnID()uint32

	//获取当前请求信息
	GetMsg()[]byte
	//获取msgID
	GetMsgID()uint32
	//获取Msglen
	GetMsgLen()int

}
