/*
* @Author: 龚国宁
* @Date: 2023/3/31 11:33
* @功能:
 */

package main

import (
	"log"
	"myrpc/controller"
	"net"
	"net/rpc"
)

func main01()  {
	rpc.RegisterName("HelloService", new(controller.HelloService))
	listener, err := net.Listen("tcp",":1234")
	if err != nil {
		log.Fatalf("ListenTCP error: %v", err)
	}
	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("Accept error: ", err)
	}
	rpc.ServeConn(conn)
}
