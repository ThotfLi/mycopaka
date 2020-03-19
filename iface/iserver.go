package iface

type IServer interface {
	//开启服务
	Start()
	//停止服务
	Stop()
	//运行服务
	Serve()
}
