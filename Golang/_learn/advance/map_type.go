package advance

import (
	"fmt"
	"time"
)

func TestMap() {
	fmt.Println("\ntest map")

	testMap()
}

func testMap() {
	map2 := map[interface{}]interface{}{
		"1":    123,
		34:     "re",
		"dsda": []string{"asd", "dsdf"},
		// []int{2}: 3, //panic: runtime error: hash of unhashable type []int
	}

	fmt.Println(map2["1"])

	var map0 map[string]int
	key := "key1"
	elem, ok := map0[key]
	fmt.Printf("The element paired with key %q in nil map: %d (%v)\n",
		key, elem, ok)
	fmt.Printf("The length of nil map: %d\n",
		len(map0))
	delete(map0, key)
	// map0[key] = 10 //panic: assignment to entry in nil map

	// map 并发读写要加原子锁
	// mapr := map[int]string{1: "str"}

	// go read(mapr)
	// time.Sleep(time.Second)
	// fatal error: concurrent map read and map write
	// go write(mapr)
	// time.Sleep(30 * time.Second)
}

func read(m map[int]string) {
	for {
		_ = m[1]
		time.Sleep(time.Second)
	}
}

func write(m map[int]string) {
	for {
		m[1] = "test"
		time.Sleep(time.Second)
	}
}
