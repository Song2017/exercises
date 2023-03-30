package pool

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/panjf2000/ants/v2"
)

var sum int32

func myFunc(i interface{}) {
	n := i.(int32)
	atomic.AddInt32(&sum, n)
	fmt.Printf("run with %d\n", n)
}

func demoFunc() string {
	time.Sleep(20 * time.Millisecond)
	fmt.Println("Hello World!")
	return "test"
}

func mainants() {

	runTimes := 1000

	// Use the common pool.
	var wg sync.WaitGroup
	syncCalculateSum := func() {
		demoFunc()
		wg.Done()
	}
	pool, _ := ants.NewPool(10)
	defer pool.Release()

	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		_ = pool.Submit(syncCalculateSum)
	}
	wg.Wait()

	fmt.Printf("running goroutines: %d\n", pool.Running())
	fmt.Printf("finish all tasks.\n")

	// Use the pool with a function,
	// set 10 to the capacity of goroutine pool and 1 second for expired duration.
	// p, _ := ants.NewPoolWithFunc(10, func(i interface{}) {
	// 	myFunc(i)
	// 	wg.Done()
	// })
	// defer p.Release()
	// // Submit tasks one by one.
	// for i := 0; i < runTimes; i++ {
	// 	wg.Add(1)
	// 	_ = p.Invoke(int32(i))
	// }
	// wg.Wait()
	// fmt.Printf("running goroutines: %d\n", p.Running())
	// fmt.Printf("finish all tasks, result is %d\n", sum)
}
