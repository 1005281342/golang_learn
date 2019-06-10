package add2

var Name string
var Age int

// 执行顺序为 导包 初始化全局变量 init函数 main函数
// init函数是在main函数之前运行的
func init() {
	Name = "22"
	Age = 22
}

// iota

// go类型推导

// 值类型和引用类型
