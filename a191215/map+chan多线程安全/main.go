package main

import (
	"fmt"
	"sync"
)

type SyncMap struct {
	MyMap map[string]int
	*sync.RWMutex
}

var m SyncMap
var ch chan bool

func main() {
	m = SyncMap{
		MyMap:   map[string]int{},
		RWMutex: new(sync.RWMutex),
	}
	ch = make(chan bool, 2)
	for {
		if len(ch) == 2 {
			Read()
			break
		} else {
			go writeA()
			go writeB()
		}
	}
}

func Read() {
	m.RLock()
	for k, v := range m.MyMap {
		fmt.Println(k, v)
	}
	m.RUnlock()
}

func writeA() {
	t := [3]string{"a1", "a2", "a3"}
	for i, v := range t {
		m.Lock()
		m.MyMap[v] = i
		m.Unlock()
	}
	ch <- true
}

func writeB() {
	t := [3]string{"b1", "b2", "b3"}
	for i, v := range t {
		m.Lock()
		m.MyMap[v] = i
		m.Unlock()
	}
	ch <- true
}
