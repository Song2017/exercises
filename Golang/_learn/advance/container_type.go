package advance

import (
	"container/heap"
	"container/list"
	"container/ring"
	"fmt"
)

func TestContainer() {
	fmt.Println("test Container")

	Slice()
	Container()
}

func Slice() {
	// 怎样判断一个变量的类型 类型断言
	ss := "test"
	value, ok := interface{}(ss).(string)
	fmt.Printf("vakue %s, result %v \n", value, ok)
	// 数组（array）类型和切片（slice）类型
	// slice
	s1 := make([]int, 5, 8)
	fmt.Printf("length %d, capacity %d, value %d \n", len(s1), cap(s1), s1)
	a1 := [5]int{1, 2, 3, 4}
	// 数组的长度是其类型的一部分
	value2, err := interface{}(a1).([5]int)
	fmt.Printf("length %d, capacity %d, value %d \n", len(a1), cap(a1), a1)
	fmt.Println(value2, err)
}

func Container() {
	// list
	l := list.New()
	for i := 97; i < 100; i++ {
		l.PushFront(i) //PushFront()代表从头部插入，同样PushBack()代表从尾部插入
	}
	for it := l.Front(); it != nil; it = it.Next() {
		v, _ := interface{}(it.Value).(int)
		fmt.Println(string(v))
	}
	var l2 list.List
	fmt.Println("var l2 list.List", l2, " end")
	var e1 list.Element
	fmt.Println("var e1 list.Element", e1, " end")
	// ring
	r := ring.New(5)
	fmt.Println("var e1 ring.New(5)", r, " end")

	// heap A heap is a common way to implement a priority queue
	h := &StuHeap{
		{name: "name1", socre: 90},
		{name: "name2", socre: 89},
		{name: "name3", socre: 99},
	}

	heap.Init(h)
	heap.Push(h, Student{name: "name4", socre: 80})
	h.Push(Student{name: "name5", socre: 100})
	for _, ele := range *h {
		fmt.Printf("name %s, score %d \n", ele.name, ele.socre)
	}
	for i, ele := range *h {
		if ele.name == "name2" {
			(*h)[i].socre = 60
			// 在修改第i个元素后，调用本函数修复堆，比删除第i个元素后插入新元素更有效率。
			// 复杂度O(log(n))，其中n等于h.Len()。
			heap.Fix(h, i)
		}
	}
	for _, ele := range *h {
		fmt.Printf("name %s, score %d \n", ele.name, ele.socre)
	}

	for h.Len() > 0 {
		// 删除并返回堆h中的最小元素（取决于Less函数，最大堆或最小堆）（不影响堆de约束性）
		// 复杂度O(log(n))，其中n等于h.Len()。该函数等价于Remove(h, 0)
		item := heap.Pop(h).(Student)
		fmt.Printf("student name %s,score %d\n", item.name, item.socre)
	}
}

type Student struct {
	name  string
	socre int
}

type StuHeap []Student

func (h StuHeap) Len() int {
	return len(h)
}

func (h StuHeap) Less(i, j int) bool {
	return h[i].socre < h[j].socre // zui xiao dui
}

func (h StuHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *StuHeap) Push(x interface{}) {
	*h = append(*h, x.(Student))
}

func (h *StuHeap) Pop() interface{} {
	old := *h
	n := len(old)
	c := old[n-1]
	*h = old[0 : n-1]
	return c
}
