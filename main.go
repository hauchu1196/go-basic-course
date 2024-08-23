package main

import (
	"fmt"
	_ "log"
)

const (
	StatusCancel  = iota + 1 // 1
	StatusSuccess            // 2
	StatusDraft              // 3
)

var (
	var1 = 0
	var2 = 0
)

// heap memory
// stack memory

type Test struct {
	Val *int64
}

// 2 bien chung 1 value
func main() {
	var myVar1, myVar2 *int64

	var myVar3 = myVar2

	fmt.Println(myVar1, myVar2, myVar3)
}
