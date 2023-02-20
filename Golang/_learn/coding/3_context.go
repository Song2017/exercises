package coding

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"
)

func TestContext() {
	fmt.Println("\nTestContext")
	testContext()
}

func testContext() {
	fmt.Println("\ntestContext")
	addNum := func(numP *int32, id int, der func()) {
		defer der()
		for i := 0; ; i++ {
			curr := atomic.LoadInt32(numP)
			time.Sleep(time.Millisecond * 20)
			if atomic.CompareAndSwapInt32(numP, curr, curr+1) {
				fmt.Printf("goroutine id %d: cas %d\n", id, curr+1)
				break
			} else {
				fmt.Printf("goroutine id %d: cas failed %d\n", id, curr+1)
			}
		}
	}

	total := 5
	var num int32
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < total; i++ {
		go addNum(&num, i, func() {
			if atomic.LoadInt32(&num) == int32(total) {
				cancel()
			}
		})
	}
	<-ctx.Done()
}
