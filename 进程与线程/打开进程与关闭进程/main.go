package main

import "os/exec"

func OpenProcess() {
	cmd := exec.Command("open", "/Users/ouyangjuxiang/Desktop/gocode/src/golang_learn/进程与线程/打开进程与关闭进程/test.txt")
	cmd.Run()
}

func main() {
	OpenProcess()
}
