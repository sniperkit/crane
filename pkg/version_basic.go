// +build !pro

package crane

import (
	"fmt"
)

const Version = "3.4.1"
const Pro = true

func printVersion() {
	fmt.Printf("v%s\n", Version)
}
