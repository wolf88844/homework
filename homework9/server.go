package main

import (
	"homework/homework9/proto"
	"net"
)

func main() {
	netListen, err := net.Listen("tcp", ":9988")
	proto.CheckError(err)

	defer netListen.Close()

	proto.Log("Waiting for clients")

	for {
		conn, err := netListen.Accept()
		if err != nil {
			continue
		}
		proto.Log(conn.RemoteAddr().String(), "tcp connect success")

		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {
	tmpBuffer := make([]byte, 0)
	readerChannel := make(chan []byte, 16)

	go reader(readerChannel)

	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			proto.Log(conn.RemoteAddr().String(), "connection err:", err)
			return
		}

		tmpBuffer = proto.Unpack(append(tmpBuffer, buffer[:n]...), readerChannel)
	}
}

func reader(readerChannel chan []byte) {
	for {
		select {
		case data := <-readerChannel:
			headerLengthInt := proto.BytesToInt16(data[proto.PackageLength : proto.PackageLength+proto.HeaderLength])
			protocol := data[proto.PackageLength+proto.HeaderLength : proto.PackageLength+proto.HeaderLength+proto.ProtocolVersion]
			operation := data[proto.PackageLength+proto.HeaderLength+proto.ProtocolVersion : proto.PackageLength+proto.HeaderLength+proto.ProtocolVersion+proto.Operation]
			sequenceId := data[proto.PackageLength+proto.HeaderLength+proto.ProtocolVersion+proto.Operation : proto.PackageLength+proto.HeaderLength+proto.ProtocolVersion+proto.Operation+proto.SequenceId]
			message := data[proto.PackageLength+proto.HeaderLength+headerLengthInt:]
			proto.Log(proto.BytesToInt16(protocol))
			proto.Log(proto.BytesToInt32(operation))
			proto.Log(proto.BytesToInt32(sequenceId))
			proto.Log(string(message))
		}
	}
}
