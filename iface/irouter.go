package iface

//路由
type IRouter interface {
	//任务开始前所要执行的处理
	PreHandle(request IRequest)
	//主任务函数
	Handle(request IRequest)
	//任务执行后所要执行的处理
	PostHandle(request IRequest)
}
