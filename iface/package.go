package iface

//包操作

type IPackage interface{
	//封包
	//将msg按照 len id Data 的顺序序列化为[]byte
	Packet(msg IMessage)   []byte
	//解包
	//msgbytes为一个长度为8的字节 包括datalen 和 id
	//返回的IMessage对象只有datalen和id字段已有赋值
	//外部通过io.ReadFull(conn,buf[:datalen]) 读取准确的数据内容
	UnPack(msgbytes []byte)IMessage

	//获取datalen+len(id)的长度
	GetHeadLen()uint32
}
