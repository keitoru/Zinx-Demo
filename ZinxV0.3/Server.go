package main

import (
	"fmt"
	"zinx/ziface"
	"zinx/znet"
)

type PingRouter struct {
	znet.BaseRouter
}

// PreHandle test
func (pr *PingRouter) PreHandle(req ziface.IRequest) {
	fmt.Println("Call Router PreHandle...")
	if _, err := req.GetConnection().GetTcpConnect().Write([]byte("before ping...")); err != nil {
		fmt.Println("call back before ping error")
	}
}

// Handle test
func (pr *PingRouter) Handle(req ziface.IRequest) {
	fmt.Println("Call Router Handle...")
	if _, err := req.GetConnection().GetTcpConnect().Write([]byte("ping...")); err != nil {
		fmt.Println("call back ping error")
	}
}

// PostHandle test
func (pr *PingRouter) PostHandle(req ziface.IRequest) {
	fmt.Println("Call Router PostHandle...")
	if _, err := req.GetConnection().GetTcpConnect().Write([]byte("after ping...")); err != nil {
		fmt.Println("call back after ping error")
	}
}

func main() {
	// 创建一个server句柄，使用Zinx的api
	s := znet.NewServer("[zinx V0.1]")

	s.AddRouter(&PingRouter{})

	// 启动server
	s.Serve()
}
