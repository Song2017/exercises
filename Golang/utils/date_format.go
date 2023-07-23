package main

import (
	"fmt"
	"time"
)

func main() {
	timeString := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(timeString)
	fmt.Println(time.Now().Format("2006_01_02"))
	fmt.Println(time.Now().Format("2006-01-02_15:04:05"))
}
