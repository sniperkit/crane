// +build !pro

package core

import (
	"fmt"
)

const Version = "3.4.1a"
const Pro = true

func printVersion() {
	fmt.Printf("v%s\n", Version)
}
