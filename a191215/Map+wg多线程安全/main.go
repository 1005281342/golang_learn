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
var wg sync.WaitGroup

func main() {
	m = SyncMap{map[string]int{}, new(sync.RWMutex)}
	wg.Add(1)
	go func() {
		writeA()
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		writeB()
		wg.Done()
	}()
	wg.Wait()
	Read()
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
}

func writeB() {
	t := [3]string{"b1", "b2", "b3"}
	for i, v := range t {
		m.Lock()
		m.MyMap[v] = i
		m.Unlock()
	}
}
