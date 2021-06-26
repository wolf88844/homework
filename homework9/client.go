package main

import (
	"homework/homework9/proto"
	"net"
	"strconv"
	"strings"
	"time"
)

func main() {
	server := "127.0.0.1:9988"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
	proto.CheckError(err)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)

	proto.CheckError(err)

	defer conn.Close()

	proto.Log("connect success")

	go sender(conn)

	for {
		time.Sleep(1 * 1e9)
	}
}

func sender(conn net.Conn) {
	for i := 0; i < 100; i++ {
		message := []string{"{`id`:", strconv.Itoa(i), ":,`name`:`golang`}"}
		words := strings.Join(message, "")
		conn.Write(proto.Packet([]byte(words)))
	}
	proto.Log("send over")
}
