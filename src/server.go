/*
1.尝试测试工具
2.git工具尝试
3.文档查看工具使用
*/
package ldlnet

import (
	"bufio"
	"fmt"
	"net"
)

func Start() {
	lister, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Printf("错误信息%v",err)
		return
	}

	fmt.Println("服务启动成功，监听端口8888")
	defer lister.Close()

	for {
		conn, err := lister.Accept()
		if err != nil {
			return
		}

		fmt.Println("用户接入")

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	fmt.Println("nihaoma")
	input := bufio.NewScanner(conn)
	for input.Scan() {
		fmt.Printf("%s\n", input.Text())
	}
	conn.Close()
}
