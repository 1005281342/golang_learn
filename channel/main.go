package main

import "fmt"

func main() {
	//closeWCase()
	//closeCCase()
	closeRCase()
}

/*
panic: send on closed channel

goroutine 1 [running]:
main.closeWCase(...)
        /Users/oyjx/Desktop/gocode/src/golang_learn/channel/main.go:10
main.main()
        /Users/oyjx/Desktop/gocode/src/golang_learn/channel/main.go:4 +0x63
*/
func closeWCase() {
	ch := make(chan int, 16)
	close(ch)
	ch <- 2
}

/*
panic: close of closed channel

goroutine 1 [running]:
main.closeCCase(...)
        /Users/oyjx/Desktop/gocode/src/golang_learn/channel/main.go:26
main.main()
        /Users/oyjx/Desktop/gocode/src/golang_learn/channel/main.go:5 +0x57
*/
func closeCCase() {
	ch := make(chan int, 16)
	close(ch)
	close(ch)
}

/*
Process finished with exit code 0
*/
func closeRCase() {
	ch := make(chan int, 16)
	ch <- 1
	ch <- 1
	ch <- 1
	ch <- 1
	ch <- 1
	close(ch)
	ans := 5
	cnt := 0
	for i := 0; i < 16; i++ {
		a, ok := <-ch
		if !ok {
			break
		}
		cnt += a
	}
	if ans != cnt {
		panic(fmt.Sprintf("should get %d, but get %d", ans, cnt))
	}
}
