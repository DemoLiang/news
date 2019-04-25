package base

import "fmt"

func Log(format string, a ...interface{}) {
	fmt.Printf(format, a...)
	return
}
