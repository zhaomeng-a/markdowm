package main

import (
	"fmt"
	"strconv"
	"sync"
)

var mu sync.Mutex
var userMap = make(map[string]int)

func mapConcurrent(i int) {
	mu.Lock()
	userMap[strconv.Itoa(i)] = i
	mu.Unlock()
}

func main() {
	for i := 1; i < 10; i++ {
		go mapConcurrent(i)
	}
	fmt.Println("userMap", userMap)
}
