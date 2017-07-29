
package main
//const b string="abc"  //显示声明常量

//const b ="asd" 		//隐是声明


//const c_name,b_name,_d_name=1,2  //多个常量声明
  



import "unsafe"
const (
    a = "abc"
    b = len(a)
    c = unsafe.Sizeof(a)
)

func main(){
    println(a, b, c)
}