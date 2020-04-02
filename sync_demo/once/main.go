package main

import "sync"

const Length = 8

var ch = make(chan struct{}, Length)
var once = sync.Once{}

func main() {
	Close()
}

func Close() {
	for i := 0; i < Length>>1; i++ {
		ch <- struct{}{}
	}
	once.Do(func() { close(ch) })
}
