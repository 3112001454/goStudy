package channel

import (
	"fmt"
	"time"
)

func Process(ch chan int) {
	// Do some work...
	time.Sleep(time.Second)
	ch <- 1 // 管道中写入一个元素表示当前协程已结束
}

func DemoChannel() {
	count := 10 // 最大支持并发
	sum := 100  // 任务总数

	c := make(chan struct{}, count) // 控制任务并发的chan
	sc := make(chan struct{}, sum)  // 控制任务总数的chan
	defer close(c)
	defer close(sc)

	for i := 0; i < sum; i++ {
		c <- struct{}{} // 作用类似于waitgroup.Add(1)
		go func(j int) {
			fmt.Println(j)
			<-c              // 执行完毕，释放资源
			sc <- struct{}{} // 记录到执行总数里
		}(i)
	}
	for i := sum; i > 0; i-- {
		<-sc
	}
}

func DemoChannel2() {
	channels := make([]chan int, 10) // 创建一个10个元素的切片，元素类型为channel
	for i := 0; i < 10; i++ {
		channels[i] = make(chan int) // 切片中放入一个channel
		go Process(channels[i])      // 启动协程，传一个管道用于通信
	}
	for i, ch := range channels { // 遍历切片，等待子协程结束
		<-ch
		fmt.Println("Routine", i, " quit!")
	}
}
