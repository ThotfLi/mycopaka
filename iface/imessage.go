package iface


//对客户端发送过来的字节序列进行封装

type IMessage interface{
	//获取信息内容
	GetData()   []byte
	//获取信息id
	GetMsgID()  uint32
	//获取信息长度
	GetMsgLen() uint32

	//设置内容
	SetData(data []byte)
	//设置id
	SetID(id uint32)
}
