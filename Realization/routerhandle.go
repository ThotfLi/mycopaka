package Realization

import(
	"fmt"
	"github.com/pkg/errors"
	"mycopaka/iface"
)
//路由处理器
type RouterHandle struct{
	apis map[uint32]iface.IRouter
}
func NewRouterHandle()iface.IRouterHandel{
	return &RouterHandle{make(map[uint32]iface.IRouter)}
}

func(p *RouterHandle)AddRouter(id uint32,router iface.IRouter){
	//判断要添加的router是否存在
	if _,ok := p.apis[id];ok == true{
		panic(fmt.Sprint("router id:",id,"Already exist"))
	}

	//添加路由
	p.apis[id] = router
}

func(p *RouterHandle)DelRouter(routerid uint32)error{
	//存在 id = routerid 的路由就删除
	if _,ok := p.apis[routerid];ok == true{
		delete(p.apis,routerid)
		return nil
	}
	return errors.New("要删除的路由不存在")
}

func(p *RouterHandle)RunRouter(request iface.IRequest){
	router,ok := p.apis[request.GetMsgID()]
	if ok != true {
		fmt.Println("Router id",request.GetMsgID(),"Non-existent")
	}
	//运行 router
	func(request iface.IRequest) {
		router.PreHandle(request)
		router.Handle(request)
		router.PostHandle(request)
	}(request)
}