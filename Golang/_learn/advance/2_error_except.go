package advance

import (
	"errors"
	"fmt"
	"net"
	"os"
	"os/exec"
)

func TestError() {
	fmt.Println("\nTestError")

	for _, ele := range []string{"", "hello"} {
		resp, err := echo(ele)
		if err != nil {
			fmt.Println("error ", err)
			continue
		}
		fmt.Println("echo response ", resp)
	}
	// 模板化生成error
	err0 := fmt.Errorf("test %s", "err0 fmt error")
	val, e := interface{}(err0).(error)
	// errT := err0.(type)
	fmt.Println("fmt.Errorf ", err0, val, e)
	err_info := fmt.Errorf("err_info %s", "err_info")
	err1 := os.PathError{Op: "op", Path: "path", Err: err_info}
	fmt.Println("err", underlyingError(&err1))
	fmt.Println("print err ")
	printError(1, &err1)

	// design
	net_error := net.AddrError{Err: "AddrError err", Addr: "AddrError addr"}
	fmt.Println(net_error)
}

func echo(req string) (resp string, err error) {
	if req == "" {
		err = errors.New("empty string")
		return
	}
	resp = fmt.Sprintf("echo : %s", req)
	return
}

func underlyingError(err error) error {
	// 获取和返回已知的操作系统相关错误的潜在错误值
	switch err := err.(type) {
	case *os.PathError:
		return err.Err
	case *os.LinkError:
		return err.Err
	case *os.SyscallError:
		return err.Err
	case *exec.Error:
		return err.Err
	}
	return err
}

func printError(i int, err error) {
	if err == nil {
		fmt.Println("nil error")
	}
	err = underlyingError(err)
	switch err {
	case os.ErrClosed:
		fmt.Printf("error closed %d %s\n", i, err)
	case os.ErrInvalid:
		fmt.Printf("error invalid %d %s\n", i, err)
	default:
		fmt.Printf("error info %d %s\n", i, err)
	}
}
