package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"time"
)

const (
	size = 10000000
)

var (
	smap map[string]int
	imap map[int]int
)

func main() {

	if os.Args[1] == "int" {
		fmt.Println("Int")
		imap = make(map[int]int, size)
		for i := 0; i < size; i++ {
			imap[i] = i
		}
	} else {
		fmt.Println("String")
		smap = make(map[string]int, size)
		for i := 0; i < size; i++ {
			smap[strconv.Itoa(i)] = i
		}
	}

	for {
		t := time.Now()
		runtime.GC()
		fmt.Printf("gc for %s took: %s\n", os.Args[1], time.Since(t))
		time.Sleep(1 * time.Second)
	}
}
