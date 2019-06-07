package package_1

func Goroute(a, b int, c chan int) {
	c <- a + b
}
