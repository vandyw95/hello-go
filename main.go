package main

import (
	"fmt"

	"github.com/hello-go/utils"
)

func main() {
	printVersion()
	fmt.Println(utils.Reverse("hello, world\n"))
	fmt.Println(learnMemory())
	tryLoop1()
	fmt.Println()
}
