package main

import (
	"fmt"
	"sync"
	"time"
)

var lock sync.Mutex
var x int32
var wg sync.WaitGroup

func main() {
	//Demo()
	start := time.Now()
	DemoRW()
	// 1.467693ms
	fmt.Println(time.Now().Sub(start))
}

func Add() {
	for i := 0; i < 1024; i++ {
		lock.Lock()
		x += 1
		lock.Unlock()
	}
	wg.Done()
}

func Demo() {
	wg.Add(1)
	go Add()
	wg.Add(1)
	go Add()
	wg.Wait()
	fmt.Println(x)
}

func DemoRW() {
	wg.Add(1)
	go func() {
		for i := 0; i < 8; i++ {
			lock.Lock()
			x += 1
			lock.Unlock()
		}
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		for i := 0; i < 8*128; i++ {
			lock.Lock()
			fmt.Println(x)
			lock.Unlock()
		}
		wg.Done()
	}()
	wg.Wait()
}
