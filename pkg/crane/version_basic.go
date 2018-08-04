// +build !pro
package crane

import (
	"fmt"
)

const (
	ProgramName = `crane`
	Version     = "3.4.2"
	Pro         = false
)

func printVersion() {
	fmt.Printf("v%s\n", Version)
}

func printVersionLong() {
	fmt.Printf("%s v%s\n", ProgramName, Version)
}
