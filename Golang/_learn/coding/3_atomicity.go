package coding

import (
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"
	"sync/atomic"
	"time"
)

type atomicValue struct {
	v atomic.Value
	t reflect.Type
}

func NewAtomValue(value interface{}) (*atomicValue, error) {
	if value == nil {
		return nil, errors.New("nil value")
	}
	return &atomicValue{t: reflect.TypeOf(value)}, nil
}

func (av *atomicValue) Load() interface{} {
	return av.v.Load()
}

func (av *atomicValue) Store(v interface{}) error {
	if v == nil {
		fmt.Println("nil value")
	}
	if av.t != reflect.TypeOf(v) {
		return fmt.Errorf("the type of value is not the same : %v != %v", av.t, reflect.TypeOf(v))
	}
	av.v.Store(v)
	return nil
}

func (av *atomicValue) TypeOfValue() reflect.Type {
	return av.t
}

func TestAtom() {
	fmt.Println("\nTestAtom")
	testAtomAdd()

	testCAS()

	testAtomValue()
}

func testAtomValue() {
	fmt.Println("\ntestAtomValue")
	// 原子值在真正使用前可以被复制
	var box atomic.Value
	box2 := box
	v1 := [...]int{1, 2, 3}
	fmt.Printf("Store %v to box2.\n", v1)
	box2.Store(v1)
	fmt.Printf("box: %v\n", box.Load())
	fmt.Printf("box2: %v\n", box2.Load())

	v2 := "123"
	fmt.Printf("Store %q to box2.\n", v2)
	box.Store(v2) // 这里并不会引发panic。
	fmt.Printf("The value load from box is %v.\n", box.Load())
	fmt.Printf("The value load from box2 is %q.\n", box2.Load())

	var box4 atomic.Value
	v4 := errors.New("error wrong")
	box4.Store(v4)
	fmt.Printf("The value load from box4 is %v.\n", box4.Load())
	v41 := io.EOF
	box4.Store(v41)
	fmt.Printf("The value load from box4 is %v.\n", box4.Load())
	v42, ok := interface{}(&os.PathError{}).(error)
	if ok {
		fmt.Println("panic: sync/atomic: store of inconsistently typed value into Value: ", v42)
		// box4.Store(v42)
	}

	box5, err := NewAtomValue(v4)
	if err != nil {
		fmt.Println("NewAtomValue error: ", err)
	}
	fmt.Printf("The legal type in box5 is %s.\n", box5.TypeOfValue())
	fmt.Println("Store a value of the same type to box5.")
	err = box5.Store(v41)
	if err != nil {
		fmt.Printf("error: %s\n", err)
	}
	fmt.Printf("The legal type in box5 is %s.\n", box5.TypeOfValue())
	fmt.Println("Store a value of the same type to box5.")
	err = box5.Store(v42)
	if err != nil {
		fmt.Printf("error: %s\n", err)
	}

	var box6 atomic.Value
	v6 := []int{1, 2, 3}
	box6.Store(v6)
	fmt.Printf("Store %v to box6.\n", v6)
	// 切片类型属于引用类型。所以，我在外面改动这个切片值，就等于修改了box6中存储的那个值。这相当于绕过了原子值而进行了非并发安全的操作
	v6[1] = 4 // 非并发安全
	fmt.Printf("box6: %v\n", box6.Load())
	// 正确
	store := func(i_box *atomic.Value, v []int) {
		replica := make([]int, len(v))
		copy(replica, v)
		i_box.Store(replica)
	}
	fmt.Printf("Store %v to box6.\n", v6)
	store(&box6, v6)
	v6[1] = 5
	fmt.Printf("box6: %v\n", box6.Load())
}

func testCAS() {
	// 用于展示一种简易的（且更加宽松的）互斥锁的模拟
	sign := make(chan struct{}, 2)
	num := int32(0)
	max := int32(20)
	// 定时增加num的值。
	atomIncrease := func(id, max int32) {
		defer func() {
			sign <- struct{}{}
		}()
		for {
			curr := atomic.LoadInt32(&num)
			if curr >= max {
				break
			}
			new_num := curr + 1
			time.Sleep(time.Millisecond * 2)
			if atomic.CompareAndSwapInt32(&num, curr, new_num) {
				fmt.Printf("The number swapped: id %d [%d-%d]\n", id, curr, new_num)
			} else {
				fmt.Printf("The CAS operation failed: id %d [%d-%d]\n", id, curr, new_num)
			}
		}
	}
	go atomIncrease(1, max)
	go atomIncrease(2, max)

	<-sign
	<-sign
}

func testAtomAdd() {
	ui1 := uint32(18)
	fmt.Printf("ui1: %d\n", ui1)
	delta := int32(-4)
	uidelta := uint32(delta)
	atomic.AddUint32(&ui1, uidelta)
	fmt.Printf("atom ui1: %d;delta: %d; uidelta: %d\n", ui1, delta, uidelta)
	atomic.AddUint32(&ui1, ^uint32(-delta-1))
	fmt.Printf("atom ui1: %d; ^uint32(-(-4)-1)的补码: %b\n", ui1, ^uint32(-delta-1)) // 与-3的补码相同。
}
