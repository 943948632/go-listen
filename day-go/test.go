package main    // 定义包名

import "fmt"    // 告诉编辑器 需要fmt这个包

//三种 变量赋值
var a="我是菜鸟"
var b string="123456789"
var c bool



var x,y init
var (
	k init  //这种写法要是全局声明变量的
	b bool
)

var j,h init=1,2
var e,f=123456,"hello"
 func main() {   // 主函数  有ini会改执行该函数
   fmt.Println("Hello, World!")
   fmt.Println("我的一个go程序")
   fmt.Println(a)
}


