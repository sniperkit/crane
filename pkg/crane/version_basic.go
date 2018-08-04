// +build !pro

package crane

import (
	"fmt"
)

<<<<<<< HEAD:pkg/version_basic.go
const Version = "3.4.1"
const Pro = true
=======
const Version = "3.4.2"
const Pro = false
>>>>>>> 2a452231eff7a1b185e03a8b848399c52cee9fbd:crane/version_basic.go

func printVersion() {
	fmt.Printf("v%s\n", Version)
}
