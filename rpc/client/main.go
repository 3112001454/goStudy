/*
* @Author: 龚国宁
* @Date: 2023/3/31 11:36
* @功能:
 */

package main

import (
	"fmt"
	"log"
	"myrpc/controller"
	"net/rpc"
)

func main01() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("Dialing: ", err)
	}
	var replay string
	err = client.Call("HelloService.Hello"," world", &replay)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(replay)
}

type HelloServiceClient struct {
	*rpc.Client
}

var _ controller.HelloServiceInterface = (*HelloServiceClient)(nil)

func DialHelloService(network, address string)(*HelloServiceClient, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{c}, nil
}

func(cli *HelloServiceClient)Hello(request string, replay *string) error {
	return cli.Client.Call(controller.HelloServiceName+".Hello", replay, replay)
}