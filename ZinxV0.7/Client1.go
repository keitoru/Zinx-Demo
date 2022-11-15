package main

import (
	"fmt"
	"io"
	"net"
	"time"
	"zinx/utils"
	"zinx/znet"
)

// 模拟客户端
func main() {
	fmt.Println("Client Test ... start")

	//1秒之后发起测试请求，给服务端开启服务的机会
	time.Sleep(1 * time.Second)

	addr := fmt.Sprintf("%s:%d", utils.GlobalObject.Host, utils.GlobalObject.TcpPort)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println("client1 start err, exit!")
		return
	}

	for {
		//发封包message消息
		dp := znet.NewDataPack()

		msg, err := dp.Pack(znet.NewMsgPackage(1, []byte("Zinx V0.6 Client1 Test Message")))
		if err != nil {
			fmt.Println("Client pack err:", err)
			return
		}

		if _, err := conn.Write(msg); err != nil {
			fmt.Println("write error err ", err)
			return
		}

		//先读出流中的head部分
		headData := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(conn, headData); err != nil { //ReadFull 会把msg填充满为止
			fmt.Println("read head error")
			return
		}

		//将headData字节流 拆包到msg中
		msgHead, err := dp.UnPack(headData)
		if err != nil {
			fmt.Println("Client unpack err:", err)
			return
		}

		if msgHead.GetMsgLen() > 0 {
			//msg 是有data数据的，需要再次读取data数据
			msg := msgHead.(*znet.Message)
			msg.Data = make([]byte, msg.GetMsgLen())
			if _, err := io.ReadFull(conn, msg.Data); err != nil {
				fmt.Println("read head error")
				return
			}

			fmt.Println("==> Recv Msg: ID=", msg.MsgId, ", len=", msg.MsgLen, ", data=", string(msg.Data))
		}
		// cpu阻塞
		time.Sleep(1 * time.Second)
	}
}
