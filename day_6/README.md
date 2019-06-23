---
title: day_6
---

### 接口
- 可以通过接口定义一组方法对类型进行抽象
- 比如汽车接口，有Run、Name等方法，宝马、比亚迪类型汽车通过实现方法Run、Name
    从而符合汽车接口，可以赋值给汽车实例
- Golang接口不需要显式实现，只要一个变量，没有implement关键字, 一个变量含有接口预定义的所有方法，即实现了这个接口

### 对自定义类型进行排序
- 查看官方排序sort的接口定义
    ```go
    type Interface interface {
        // Len方法返回集合中的元素个数
        Len() int
        // Less方法报告索引i的元素是否比索引j的元素小
        Less(i, j int) bool
        // Swap方法交换索引i和j的两个元素
        Swap(i, j int)
    }
    ```
- 自定义的结构体以及方法实现
    ```go
    package main
    
    import (
        "fmt"
        "math/rand"
        "sort"
    )
    
    type Student struct {
        Name     string
        Id       string
        Age      int
        sortType int
    }
    
    type Book struct {
        Name   string
        Author string
    }
    
    type StudentArray []Student
    
    func (p StudentArray) Len() int {
        return len(p)
    }
    
    func (p StudentArray) Less(i, j int) bool {
        return p[i].Name < p[j].Name
    }
    
    func (p StudentArray) Swap(i, j int) {
        p[i], p[j] = p[j], p[i]
    }
    
    func main() {
        var stus StudentArray
        for i := 0; i < 10; i++ {
            stu := Student{
                Name: fmt.Sprintf("stu%d", rand.Intn(100)),
                Id:   fmt.Sprintf("110%d", rand.Int()),
                Age:  rand.Intn(100),
            }
        
            stus = append(stus, stu)
        }
        
        for _, v := range stus {
            fmt.Println(v)
        }
        
        fmt.Println("\n\n")
        sort.Sort(stus)
        for _, v := range stus {
            fmt.Println(v)
        }
    }
    ```

### 接口嵌套
- 可以实现多个接口，定义一个新接口时可以引用这些接口
    ```go
    package main
    
    import "fmt"
    
    type Reader interface {
        Read()  
    }
    
    type Writer interface {
        Write()
    }
    
    type ReadWriter interface {
        Reader
        Writer
    }
    
    type File struct {
    }
    
    func (f *File) Read() {
        fmt.Println("read data")
    }
    
    func (f *File) Write() {
        fmt.Println("write data")
    }
    
    func Test(rw ReadWriter) {
        rw.Read()
        rw.Write()
    }
    
    func main() {
        var f *File
        var b interface{}
        b = f
        //Test(&f)
        v, ok := b.(ReadWriter)
        fmt.Println(v, ok)
    }
    ```
 - 对于嵌套接口的实现，同样需要实现接口的所有方法
 
 
### 一般类型转具体类型
```go
package main

import "fmt"

type Student struct {
    Name string
    Sex  string
}

func Test(a interface{}) {
    b, ok := a.(Student)
    if ok == false {
        fmt.Println("convert failed")
        return
    }
    //b += 3
    fmt.Println(b)
}

func just(items ...interface{}) {
    for index, v := range items {
        switch v.(type) {
        case bool:
            fmt.Printf("%d params is bool, value is %v\n", index, v)
        case int, int64, int32:
            fmt.Printf("%d params is int, value is %v\n", index, v)
        case float32, float64:
            fmt.Printf("%d params is float, value is %v\n", index, v)
        case string:
            fmt.Printf("%d params is string, value is %v\n", index, v)
        case Student:
            fmt.Printf("%d params student, value is %v\n", index, v)
        case *Student:
            fmt.Printf("%d params *student, value is %v\n", index, v)
        }
    }
}

func main() {
    var b Student = Student{
        Name: "stu01",
        Sex:  "female",
    }
    Test(b)
    just(28, 8.2, "this is a test", b, &b)
}

``` 

### 反射
- 在运行时动态获取变量的相关信息
- import ("reflect")