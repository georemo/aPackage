// v0.0.1
package aPackage

import (
	"fmt"
)

func init() {
	fmt.Println("init() a")
}

func A() {
	fmt.Println("This is function A!")
}

func B() {
	fmt.Println("privateConstant:", privateConstant)
}

const MyConstant = 123
const privateConstant = 21
