package coding

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"sync"
	"sync/atomic"
	"time"
)

func TestSync() {
	fmt.Println("\n testSync")

	testMutex()

	testCond()

	testWaitGroup()

	testOnce()
}

func testOnce() {
	fmt.Println("\ntestOnce")
	var once sync.Once
	var wg sync.WaitGroup
	wg.Add(3)
	// wg.Add(-4) // panic: sync: negative WaitGroup counter
	var once2 sync.Once
	go func() {
		defer wg.Done()
		once2.Do(func() {
			time.Sleep(time.Millisecond * 50)
			fmt.Println("Once Do another task Done...")
		})
	}()

	go onceFunc(&once, wg.Done)
	go func() {
		defer wg.Done()
		once.Do(func() {
			time.Sleep(time.Millisecond * 50)
			fmt.Println("Once Do another task can not run...")
		})
	}()

	wg.Wait()
}

func onceFunc(i_once *sync.Once, def func()) {
	defer func() { def() }()
	i_once.Do(func() {
		for i := 0; i < 3; i++ {
			fmt.Printf("Once Do task. [%d]\n", i)
			time.Sleep(time.Millisecond * 50)
		}
	})
	fmt.Println("Once Do task Done...")
}

func testWaitGroup() {
	fmt.Println("\ntestWaitGroup---------")
	var wg sync.WaitGroup
	wg.Add(2)
	num := int32(0)
	max := int32(10)

	go addNum(&num, 1, max, wg.Done)
	go addNum(&num, 2, max, wg.Done)

	wg.Wait()
}

func addNum(num *int32, id, max int32, deferFunc func()) {
	defer func() { deferFunc() }()

	for i := 0; ; i++ {
		curr := atomic.LoadInt32(num)
		if curr >= max {
			break
		}
		time.Sleep(time.Millisecond * 50)
		new := curr + 1
		if atomic.CompareAndSwapInt32(num, curr, new) {
			fmt.Printf("The number done: %d [id %d- iter %d]\n", new, id, i)
		} else {
			fmt.Printf("The number swap failed: %d [id %d- iter %d]\n", new, id, i)
		}
	}
}

func testCond() {
	fmt.Println("testCond")
	// mailbox 代表信箱。
	// 0代表信箱是空的，1代表信箱是满的。
	var mailbox uint8
	// lock 代表信箱上的锁。
	var lock sync.Mutex

	var sendCond = sync.NewCond(&lock)
	receiveCond := sync.NewCond(&lock)

	send := func(id, index int) {
		lock.Lock()
		for mailbox == 1 {
			sendCond.Wait()
		}
		log.Printf("sender [id %d-%d]: the mailbox is empty.",
			id, index)
		mailbox = 1
		log.Printf("sender [id %d-%d]: the letter has been sent.",
			id, index)
		lock.Unlock()
		// 发送信件，需要通知接收者
		receiveCond.Broadcast()
	}

	recv := func(id, index int) {
		lock.Lock()
		for mailbox == 0 {
			receiveCond.Wait()
		}
		log.Printf("receiver [id %d-%d]: the mailbox is full.", id, index)
		mailbox = 0
		log.Printf("receiver [id %d-%d]: the letter has been received.", id, index)
		lock.Unlock()
		// 已接收信件，需要通知发送者
		sendCond.Signal()
	}
	// sign 用于传递演示完成的信号。
	sign := make(chan struct{}, 3)
	max := 6
	// 发信者
	go func(id, iter int) {
		defer func() {
			sign <- struct{}{}
		}()
		for i := 1; i <= iter; i++ {
			time.Sleep(time.Millisecond * 50)
			send(id, i)
		}
	}(0, max)

	// 收信者
	receiver := func(id, iter int) {
		defer func() {
			sign <- struct{}{}
		}()
		for i := 1; i <= iter; i++ {
			time.Sleep(time.Millisecond * 20)
			recv(id, i)
		}
	}
	go receiver(1, max/2)
	go receiver(2, max/2)

	// 等待演示完成。
	<-sign
	<-sign
	<-sign
}

func testMutex() {
	fmt.Println("testMutex")
	// protecting 用于指示是否使用互斥锁来保护数据写入。
	// 若值等于0则表示不使用，若值大于0则表示使用。
	// 改变该变量的值，然后多运行几次程序，并观察程序打印的内容。
	protecting := 1
	var buffer bytes.Buffer

	const (
		gr_num   = 5
		data_num = 10
		data_len = 10
	)
	var mu sync.Mutex
	// sign 代表信号的通道。 同步信息
	sign := make(chan struct{}, gr_num)

	for i := 1; i < gr_num; i++ {
		go func(id int, w io.Writer) {
			defer func() {
				sign <- struct{}{}
			}()

			for j := 0; j < data_num; j++ {
				header := fmt.Sprintf("\nid %d, itera %d: ", id, j)
				data := fmt.Sprintf("%d*%ddata %d; ", id, j, id*j)
				if protecting > 0 {
					mu.Lock()
				}

				_, err := w.Write([]byte(header))
				if err != nil {
					log.Printf("-----write header error: %v", err)
				}
				for k := 0; k < data_len; k++ {
					_, err := w.Write([]byte(data))
					if err != nil {
						log.Printf("-----write data error: %v", err)
					}
				}
				if protecting > 0 {
					mu.Unlock()
				}
			}
		}(i, &buffer)
	}

	for i := 1; i < gr_num; i++ {
		<-sign
	}

	data, err := io.ReadAll(&buffer)
	if err != nil {
		log.Printf("read data error: %v", err)
	}
	log.Printf("all data: %s", data)
}

type OnceDemo struct {
	done int32
	m    sync.Mutex
}

func (od *OnceDemo) Do(f func()) {
	// 如果done的值为0，则执行f，并将done设置为1。
	// 第一次原子操作只需要保证f被执行了
	if atomic.LoadInt32(&od.done) == 0 {
		// 互斥锁保证f只执行一次
		od.m.Lock()
		defer od.m.Unlock()
		if od.done == 0 {
			f()
			atomic.StoreInt32(&od.done, 1)
		}
	}
}
