package main

import (
	"fmt"

	"github.com/metakeule/mutex"
)

type tester struct {
	mutex.RWMutex
	sth bool
}

func main() {
	var t tester
	t.RWMutex = mutex.NewRWMutex("tester mutex", true)
	aFunc(&t)
}

func aFunc(t *tester) {
	t.Lock()
	t.sth = true
	t.Unlock()
	t.RLock()
	sth := t.sth
	t.RUnlock()
	fmt.Printf("sth: %v\n", sth)
}
