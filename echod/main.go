package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	listenaddr := "localhost:8000"
	listener, err := net.Listen("tcp", listenaddr)
	ctx, cancel := context.WithCancel(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Echo Server is listen:", listenaddr)

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
		select {
		case <-ctx.Done():
			conn.Close()
			listener.Close()
			return
		default:
		}
		go HandleConn(ctx, conn)
	}
}

func HandleConn(ctx context.Context, c net.Conn) {
	defer c.Close()
	deadline := time.Now().Add(3 * time.Minute)
	err := c.SetReadDeadline(deadline)
	if err != nil {
		log.Println("Set ",c,err)
	}
	log.Println(c.RemoteAddr())
	input := bufio.NewScanner(c)
	for input.Scan() {
		select {
		case <-ctx.Done():
			c.Close()
			break
		default:
		}
		echo(c, input.Text(), 1*time.Second)
	}
}

func echo(c net.Conn, shout string, delay time.Duration) {
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
}

func WaitForInterrupt(release func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)
	s := <-c
	log.Println("Receiving signal:", s)
	release()
}
