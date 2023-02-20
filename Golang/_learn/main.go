package main

import (
	"flag"
	"fmt"

	ad "hello/advance"
	co "hello/coding"
	lib "hello/package_demo"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "this is name comment")
}

func main() {
	flag.Parse()
	fmt.Printf("this is name %s \n", name)
	// type
	lib.TestPackage()

	ad.TestContainer()

	ad.TestMap()

	ad.TestChannel()

	ad.TestFunction()

	ad.TestStruct()

	ad.TestInterface()

	ad.TestPointer()

	// advance
	ad.TestStatement()

	ad.TestFlowControl()

	ad.TestError()

	ad.TestPRD()

	// coding
	co.TestCases() // tests go run

	co.TestSync()

	co.TestAtom()

	co.TestContext()
}
