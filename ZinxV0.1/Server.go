package main

import "zinx/znet"

func main()  {
	// 创建一个server句柄，使用Zinx的api
	s := znet.NewServer("[zinx V0.1]")
	// 启动server
	s.Serve()
}