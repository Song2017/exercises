package coding

import (
	"fmt"
	"math"
	"testing"
	"time"
)

func hello(name string) string {
	if name == "" {
		return "name is empty"
	}

	return "Hello " + name + "!"
}

func TestHello(t *testing.T) {
	// go test ./coding -run TestHello
	// go test ./coding -cpu=1,2,3 -count=1 -run TestHello
	fmt.Println("\n TestHello----")
	var name string
	resp := hello(name)
	if resp == "" {
		t.Errorf("The error is nil, but it should not be. (name=%q)",
			name)
	}
	name = "robot"
	resp = hello(name)
	expected := "Hello robot wrong"
	if resp != expected {
		t.Errorf("The response is %q, but it should be %q. (name=%q)",
			resp, expected, name)
	}
	t.Logf("End: The expected is %q.\n", expected)
}
func TestHello2(t *testing.T) {
	// go test ./coding -run TestHello2
	// go test ./coding -cpu=1,2,4 -count=3 -run TestHello2
	// go clean -cache
	// go clean -testcache //删除所有的测试结果缓存
	// gocacheverify=1将会导致 go 命令绕过任何的缓存数据
	fmt.Println("\n testCases")
	var name string
	resp := hello(name)
	if resp == "" {
		t.Errorf("The error is nil, but it should not be. (name=%q)",
			name)
	}
	name = "robot"
	resp = hello(name)
	expected := "Hello robot!"
	if resp != expected {
		t.Errorf("The response is %q, but it should be %q. (name=%q)",
			resp, expected, name)
	}
	t.Logf("End: The expected is %q.\n", expected)
}

func TestFail(t *testing.T) {
	// go test ./coding -run TestFail
	t.FailNow()
	t.Log("TestFail")
}

func BenchmarkHello(b *testing.B) {
	// go test ./coding -run=^$ -bench=. BenchmarkHello
	for i := 0; i < b.N; i++ {
		getPrimes(1000)
	}
}
func BenchmarkHello2(b *testing.B) {
	// go test ./coding -run=^$ -bench=. BenchmarkHello2
	b.Log("BenchmarkHello2")
	b.StartTimer()
	time.Sleep(2 * time.Second)
	b.StopTimer()
	max_num := 1000
	for i := 0; i < b.N; i++ {
		getPrimes(max_num)
	}
}

func getPrimes(n int) []int {
	var primes []int
	n = int(math.Sqrt(float64(n)))
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes
}
func isPrime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
