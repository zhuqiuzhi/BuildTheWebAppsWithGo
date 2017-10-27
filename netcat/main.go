package main


import (
	"io"
	"net"
	"log"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")

	if err != nil {
		log.Fatal(err)
	}

	done := make(chan struct{})
	go func() {
		// 如果调用conn.Close(), Copy 返回一个 "read from closed connection"
		// 如果调用CloseWrite()
		io.Copy(os.Stdout,conn)
		log.Println("Done")
		done <- struct{}{}
	}()

	// 当按下ctrl + d 时， mustCopy return
	// 然后调用 Close()
	mustCopy(conn, os.Stdin)
	if  tcpcoon, ok := conn.(*net.TCPConn); ok {
		log.Println("This is a TCP connect, close write")
		// client 主动关闭发送连接
		// client ---关闭--> server
		// 		  <--打开--
		tcpcoon.CloseWrite()
	} else {
		// clinet ---关闭---> server
		//        <--关闭----
		// Any blocked Read or Write operations will be unblocked and return errors.
		conn.Close()
	}
	<- done
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _,err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}