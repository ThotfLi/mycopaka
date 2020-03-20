package iface

//多路由管理器
type IRouterHandel interface{
	//添加路由
	AddRouter(id uint32,router IRouter)
	//删除路由
	DelRouter(routerid uint32)error
	//执行路由
	RunRouter(request IRequest)

}
