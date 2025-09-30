package main

import (
	"ChatClient/data"
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

func read(conn net.Conn) {
	for {
		res := make([]byte, 1024)
		n, err := conn.Read(res)
		if err != nil {
			fmt.Println(err)
			return
		}

		result := res[:n]
		var message data.Message
		json.Unmarshal(result, &message)
		if len(message.User) > 0 {
			fmt.Printf("[%s]:%s\n", message.User, message.Content)
		} else {
			fmt.Printf("%s\n", message.Content)
		}
	}

}

func write(conn net.Conn) {
	for {
		reader := bufio.NewReader(os.Stdin)
		content, _ := reader.ReadString('\n')
		_, err := conn.Write([]byte(content))
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func main() {
	fmt.Println("建立连接...")
	conn, err := net.Dial("tcp", "127.0.0.1:9000")
	if err != nil {
		fmt.Println("连接失败!")
		return
	}
	defer conn.Close()

	go read(conn)

	write(conn)
}
