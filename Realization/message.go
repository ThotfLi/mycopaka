package Realization
import(
	"mycopaka/iface"
)
type Message struct{
	data []byte
	id   uint32
	dataLen uint32
}

func NewMessage(data []byte,id uint32)iface.IMessage{
	return &Message{
		data: data,
		id:   id,
		dataLen:uint32(len(data)),
	}
}

func(m *Message)GetData()[]byte{
	return m.data
}
func(m *Message)GetMsgID() uint32{
	return m.id
}
func(m *Message)GetMsgLen() uint32{
	return m.dataLen
}

func(m *Message)SetData(data []byte){
	m.data = data
}

func(m *Message)SetID(id uint32){
	m.id = id
}
