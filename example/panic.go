package main

import (
	"fmt"
	"log"

	"github.com/itcuihao/gcrash"
)

func init() {
	if err := gcrash.Catch("panic.log"); err != nil {
		log.Fatalf("er:%v", err)
	}
}

// go run main.go
func main() {
	var s []int
	fmt.Println(s[1])
}
