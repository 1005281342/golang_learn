package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	//MapDemo20() // Panic: fatal error: concurrent map writes
	SyncMapDemo20() // OK
}

// Panic: fatal error: concurrent map writes
func MapDemo20() {
	mp := make(map[int]int, 20)
	for i := 0; i < 20; i++ {
		ii := i
		wg.Add(1)
		go func() {
			mp[ii] = i
			wg.Done()
		}()
	}
	wg.Wait()
	for k, v := range mp {
		fmt.Println(k, v)
	}
}

func SyncMapDemo20() {
	mp := sync.Map{}
	for i := 0; i < 20; i++ {
		ii := i
		wg.Add(1)
		go func() {
			mp.Store(ii, ii)
			wg.Done()
		}()
	}
	wg.Wait()
	//Range
	//遍历sync.Map, 要求输入一个func作为参数
	f := func(k, v interface{}) bool {
		//这个函数的入参、出参的类型都已经固定，不能修改
		//可以在函数体内编写自己的代码，调用map中的k,v

		fmt.Println(k, v)
		return true
	}
	mp.Range(f)
}
