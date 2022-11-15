package main

import (
	"fmt"
	"zinx/ziface"
	"zinx/znet"
)

type PingRouter struct {
	znet.BaseRouter
}

// Handle test
func (pr *PingRouter) Handle(req ziface.IRequest) {
	fmt.Println("Call Router Handle...")

	//先读取客户端的数据，再回写ping...ping...ping
	fmt.Println("recv from client : msgId=", req.GetMsgID(), ", data=", string(req.GetMsgData()))

	//回写数据
	if err := req.GetConnection().SendMsgData(1, []byte("ping...ping...ping")); err != nil {
		fmt.Println(err)
	}
}

func main() {
	// 创建一个server句柄，使用Zinx的api
	s := znet.NewServer("[zinx V0.5]")

	s.AddRouter(&PingRouter{})

	// 启动server
	s.Serve()
}
