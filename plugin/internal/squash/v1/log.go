package squash

import (
	"fmt"
	"os"
)

var Verbose bool

func Debugf(format string, args ...interface{}) {
	if verbose {
		fmt.Fprintf(os.Stderr, fmt.Sprintf("%s", format), args...)
	}
}

func Debug(args ...interface{}) {
	if Verbose {
		fmt.Fprintln(os.Stderr, args...)
	}
}

func Fatalf(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, fmt.Sprintf("ERROR: %s", format), args...)
	signals <- os.Interrupt
	wg.Wait()
	os.Exit(1)
}

func Fatal(args ...interface{}) {

	fmt.Fprint(os.Stderr, "ERROR: ")
	fmt.Fprintln(os.Stderr, args...)
	signals <- os.Interrupt
	wg.Wait()
	os.Exit(1)
}
