package advance

import (
	"fmt"
	"math/rand"
	"time"
)

func TestChannel() {
	fmt.Println("\nChannel Test")

	ch1 := make(chan int, 3)
	ch1 <- 1
	fmt.Println("first ele", <-ch1)
	ch1 <- 3
	ch1 <- 4
	ch1 <- 5
	// ch1 <- 5 //fatal error: all goroutines are asleep - deadlock!

	// ch2 := make(chan int, 1)
	// ele, err := <-ch2
	// fmt.Println("empty channel", ele, err)

	// close channel
	ch := make(chan int, 4)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("sender: send", i)
			ch <- i
		}
		fmt.Println("sender: close channel")
		close(ch)
	}()

	for {
		elem, ok := <-ch
		if !ok {
			fmt.Println("receiver: close channel")
			break
		}

		fmt.Println("receiver ele ", elem)
	}

	// one way channel
	uselessChan := make(chan<- int, 3)
	uselessChan <- 3
	// <-uselessChan //receive from send-only channel uselessChan
	chan1 := make(chan int, 3)
	SendInt(chan1)
	fmt.Println("send only function: ", <-chan1)

	chan2 := getChan(6)
	for ele := range chan2 {
		fmt.Println("receive-only channel", ele)
	}

	// select chan
	SelectChannel()
}

func SendInt(ch chan<- int) {
	ch <- rand.Intn(100)
}

type Notifier interface {
	SendInt(chan<- int)
}

func getChan(num int) <-chan int {
	ch := make(chan int, num)
	for i := 0; i < num; i++ {
		ch <- i
	}
	close(ch)
	return ch
}

func SelectChannel() {
	initChannels := [3]chan int{
		make(chan int, 1),
		make(chan int, 1),
		make(chan int, 1),
	}

	index := rand.Intn(3)
	fmt.Println("SelectChannel ", index)

	initChannels[index] <- index

	select {
	case <-initChannels[0]:
		fmt.Println("SelectChannel 0 ")
	case <-initChannels[1]:
		fmt.Println("SelectChannel 1 ")
	case <-initChannels[2]:
		fmt.Println("SelectChannel 2 ")
	default:
		fmt.Println("SelectChannel default ")
	}

	// select语句只能对其中的每一个case表达式各求值一次
	initChan := make(chan int, 2)
	initChan <- 5
	time.AfterFunc(time.Second, func() {
		close(initChan)
	})
	select {
	case ele, ok := <-initChan:
		if !ok {
			fmt.Println("close select channel")
			break
		}
		fmt.Println("receive channel ", ele)
	default:
		fmt.Println("receive default ")
	}
}
