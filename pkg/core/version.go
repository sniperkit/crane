// +build basic
package core

import (
	"fmt"
)

const Version = "3.4.2"
const Pro = false

func printVersion() {
	fmt.Printf("v%s\n", Version)
}
