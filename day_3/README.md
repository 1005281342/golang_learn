---
title: day_3
---

### Golang函数特点
- 不支持重载，一个包中不能含有名字一样的函数
- 函数是一等公民，函数也是一种类型，可以赋值给变量
- 匿名函数
- 多值返回

### 可变参数
```go
func testCase(arg ...int){
	return
}
```

### 匿名函数
```go
var (
	res = func(a, b int) int {
		return a+b
	}
)

func main() {
    fmt.Println(res(2, 5))
    res2 := func(a, b int) int {
		return a+b
	}(2, 5)
    fmt.Println(res2)
}
```

### defer
- 当函数返回时，执行defer语句。因此，可以用来做资源清理
- 多个defer语句，按先进后出的方式执行
- defer语句中的变量，在defer声明时就决定了
- defer主要用于资源回收，如关闭文件句柄、锁资源释放、数据库连接释放

### 使用管道控制并发
```go
goCount := 2
goCoroutine := make(chan int, goCount)
for i, sHotelId := range ids {
    goCoroutine <- 1
    go func(req, a int) {
        defer func() {
            <-goCoroutine
        }()
        log.Println(a)
        // do main
        DoMain(req)
    }(sHotelId, i)
}

for i := 0; i < goCount; i++ {
    goCoroutine <- 1
}

```