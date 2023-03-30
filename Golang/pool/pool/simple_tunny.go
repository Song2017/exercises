package pool

import (
	"fmt"
	"sync"

	"github.com/parnurzeal/gorequest"
)

func FunPool() {
	jobs := []string{
		"a", "b", "c", "d", "e", "2a", "2b", "2c", "2d", "2e"}
	wg := new(sync.WaitGroup)
	worker := make(chan string)
	result := make(chan string)

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for req := range worker {
				result <- callHttp(req)
			}
		}()
	}
	fmt.Println("send worker ")
	go func() {
		for _, job := range jobs {
			worker <- job
		}
		fmt.Println("worker close")
		close(worker)
		wg.Wait()
		fmt.Println("result close")
		close(result)
	}()
	fmt.Println("result print")
	for r := range result {
		fmt.Println(r)
	}

}

func callHttp(req string) string {
	request := gorequest.New()
	resp, _, errs := request.Get("http://www.baidu.com").
		// RedirectPolicy(redirectPolicyFunc).
		// Set("If-None-Match", `W/"wyzzy"`).``
		End()
	fmt.Println(req, errs)
	return resp.Request.Host
}
