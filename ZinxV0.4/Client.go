package main

import (
	"fmt"
	"net"
	"time"
	"zinx/utils"
)

// 模拟客户端
func main() {
	fmt.Println("Client Test ... start")

	//1秒之后发起测试请求，给服务端开启服务的机会
	time.Sleep(1 * time.Second)

	addr := fmt.Sprintf("%s:%d", utils.GlobalObject.Host, utils.GlobalObject.TcpPort)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println("client start err, exit!")
		return
	}

	for {
		// 链接调用Write写数据
		if _, err := conn.Write([]byte("Hello Zinx V0.1...")); err != nil {
			fmt.Println("write error err ", err)
			return
		}

		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read error err ", err)
			return
		}

		fmt.Printf(" server call back : %s, cnt = %d\n", buf, cnt)

		// cpu阻塞
		time.Sleep(1 * time.Second)
	}
}
