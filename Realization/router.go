package Realization

import(
	"mycopaka/iface"
)

type Router struct{

}

func(r *Router)PreHandle(request iface.IRequest){}

func(r *Router)Handle(request iface.IRequest){}

func(r *Router)PostHandle(request iface.IRequest){}