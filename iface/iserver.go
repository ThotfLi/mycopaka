package iface

type IServer interface {
	//开启服务
	Start()
	//停止服务
	Stop()
	//运行服务
	Serve()

	//添加新路由
	AddRouter(id uint32,router IRouter)
	//删除路由
	DelRouter(id uint32)
}
