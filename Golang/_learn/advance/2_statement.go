package advance

import (
	"fmt"
	"sync/atomic"
	"time"
)

func TestStatement() {
	fmt.Println("\n\nTestStatement")
	test1()
	orderRoutine()
}
func test1() {
	num := 10
	sign := make(chan struct{}, num)
	for it1 := 0; it1 < 10; it1++ {
		go func() {
			fmt.Println(it1)
			sign <- struct{}{}
		}()
	}
	// 1. time.Sleep函数会在被调用时用当前的绝对时间，再加上相对时间计算出在未来的恢复运行时间
	// time.Sleep(time.Second)
	// 2. 让其他的 goroutine 在运行完毕的时候告诉我们
	for j := 0; j < 10; j++ {
		<-sign
	}
}

func orderRoutine() {
	count := uint32(0) //信号量
	trigger := func(n uint32, fn func()) {
		for {
			if nc := atomic.LoadUint32(&count); n == nc {
				fn()
				atomic.AddUint32(&count, 1) //信号的值总是下一个可以调用打印函数的go函数的序号
				break
			}
			time.Sleep(time.Nanosecond) //进入下一个迭代
		}
	}

	print := func(ii int) {
		fn := func() {
			fmt.Println(ii)
		}
		trigger(uint32(ii), fn) // 安全进行并发打印
	}
	for i := 0; i < 10; i++ {
		go print(i)
	}

	trigger(uint32(10), func() {}) // 确保count达到10后退出
}
