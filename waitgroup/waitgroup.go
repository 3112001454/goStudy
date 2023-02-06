/*
* @Author: grant
* @Date: 2023/2/6 11:37
* @功能:
 */

package waitgroup

import (
	"fmt"
	"sync"
	"time"
)

func Demo1() {
	var wg sync.WaitGroup

	wg.Add(2) // 设置计数器，数值即为goroutine的个数
	go func() {
		defer wg.Done() // goroutine 执行结束后将计数器减一
		// Do some work
		time.Sleep(1 * time.Second)
		fmt.Println("Goroutine 1 finished!")
	}()

	go func() {
		defer wg.Done()
		time.Sleep(2 * time.Second)
		fmt.Println("Goroutine 2 finished!")
	}()

	wg.Wait() // 主goroutine阻塞等待计数器变为0
	fmt.Println("All Goroutine finished!")

}

func Demo2() {
	work := NewPool(4)
	for i := 0; i < 100; i++ {
		work.Add(1)
		go worker(i, work)
	}
	fmt.Println("waiting...")
	work.Wait()
	fmt.Println("done")
}

func worker(i int, wg *WaitGroup) {
	defer wg.Done()
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "output:", i)
	time.Sleep(time.Second * 1)
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "output", i, "done")
}

// WaitGroup 一个异步结构体
type WaitGroup struct {
	workChan chan int
	wg       sync.WaitGroup
}

// NewPool 生成一个工作池，coreNum限制
func NewPool(coreNum int) *WaitGroup {
	ch := make(chan int, coreNum)
	return &WaitGroup{
		workChan: ch,
		wg:       sync.WaitGroup{},
	}
}

// Add 添加
func (ap *WaitGroup) Add(num int) {
	for i := 0; i < num; i++ {
		ap.workChan <- i
		ap.wg.Add(1)
	}
}

// Done 完结
func (ap *WaitGroup) Done() {
LOOP:
	for {
		select {
		case <-ap.workChan:
			break LOOP
		}
	}
	ap.wg.Done()
}

// Wait 等待
func (ap *WaitGroup) Wait() {
	ap.wg.Wait()
}
