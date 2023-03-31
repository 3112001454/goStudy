/*
* @Author: 龚国宁
* @Date: 2023/3/31 11:30
* @功能:
 */

package controller

import "net/rpc"

type HelloService struct {

}

func(ctrl *HelloService)Hello(request string, replay *string) error {
	*replay = "hello" + request
	return nil
}

const HelloServiceName = "HelloService"

type HelloServiceInterface interface {
	Hello(request string, replay *string) error
}

func RegisterHelloService(svc HelloServiceInterface) error {
	return rpc.RegisterName(HelloServiceName, svc)
}