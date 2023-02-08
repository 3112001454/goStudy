package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	conn := make(chan string)
	WaitChannel(conn)
	AfterDemo()
	AfterFuncDemo()
	TickerDemo()
}

func WaitChannel(conn <-chan string) bool {
	timer := time.NewTimer(1 * time.Second)
	select {
	case <-conn:
		timer.Stop()
		return true
	case <-timer.C: // 超时
		fmt.Println("WaitChannel timeout!")
		return false
	}
}

// 有时候我们就是想等待指定的时间，没有需求提前停止定时器，也没有需求复用该定时器，那么可以
// 使用匿名的定时器。

func AfterDemo() {
	log.Println(time.Now())
	<- time.After(1*time.Second)
	log.Println(time.Now())
}

func AfterFuncDemo(){
	log.Println("AfterFuncDemo start:", time.Now())
	time.AfterFunc(1*time.Second, func() {
		log.Println("AfterFuncDemo end:", time.Now())
	})
	time.Sleep(2*time.Second) // 等待协程退出
}

func TickerDemo() {
	ticker := time.NewTicker(1*time.Second)
	defer ticker.Stop()
	for range ticker.C {
		log.Println("Ticker tick.")
	}
}

func fun1() {
	i := 0
	i, j := 1, 2
	fmt.Printf("i = %d, j = %d\n", i, j)
}
