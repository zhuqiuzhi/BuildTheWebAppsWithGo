package main

import (
	"io"
	"time"
	"net"
	"log"
	"context"
	"os/signal"
	"os"
	"syscall"
)

func main() {
	listenaddr := "localhost:8000"
	listener, err := net.Listen("tcp", listenaddr)
	ctx, cancel := context.WithCancel(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Time Server is listen:",listenaddr)

	go AcceptAndHandleConn(ctx, listener)

	// 等待中断信号，当接收到中断信号，调用 cancel 函数
	WaitForInterrupt(func() {
		cancel()
	})
}

func AcceptAndHandleConn(ctx context.Context, listener net.Listener) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}
		go HandleConn(ctx, conn)
	}
}

func HandleConn(ctx context.Context, conn net.Conn) {
	defer conn.Close()
	select {
	case <-ctx.Done():
	default:
		log.Println(conn.RemoteAddr())
		_ ,err := io.WriteString(conn, time.Now().Format(time.RFC1123))
		if err != nil {
			log.Println("Write Connnect error:", err)
		}
	}
}

func WaitForInterrupt(release func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill,syscall.SIGTERM)
	s := <-c
	log.Println("Receiving signal:", s)
	release()
}
