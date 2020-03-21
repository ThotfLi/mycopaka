package Realization

import(
	"bytes"
	"encoding/binary"
	"fmt"
	"mycopaka/iface"
)

type Package struct{

}

func(p *Package)Packet(msg iface.IMessage)[]byte{
	bufer := bytes.NewBuffer([]byte{})
	datalen := msg.GetMsgLen()

	//写入长度
	if err := binary.Write(bufer,binary.LittleEndian,datalen);err != nil {
		fmt.Println("[ERROR]binary datalen bytes writer failed,",err)
		return nil
	}

	//写入id
	if err := binary.Write(bufer,binary.LittleEndian,msg.GetMsgID());err != nil {
		fmt.Println("[ERROR]binary msgID bytes writer failed",err)
		return nil
	}

	//写入msgdata
	if err := binary.Write(bufer,binary.LittleEndian,msg.GetData()); err != nil {
		fmt.Println("[ERROR]binary msgData bytes write failed",err)
		return nil
	}

	return bufer.Bytes()
}

func(p *Package)UnPack(msgbyte []byte)iface.IMessage{
	var newMsg Message
	bufer := bytes.NewBuffer(msgbyte)
	//解包 先反序列化data大小
	binary.Read(bufer,binary.LittleEndian,&newMsg.dataLen)

	//反序列化id
	binary.Read(bufer,binary.LittleEndian,&newMsg.id)

	//返回msg
	return &newMsg

}

func(p *Package)GetHeadLen()uint32{
	return 8
}