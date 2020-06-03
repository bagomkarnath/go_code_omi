package main

import (
	"fmt"
)

const (
	_  = iota // ignore
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {
	fileSize := 400000000.
	fmt.Printf("%.2fGB", fileSize/GB)
}
