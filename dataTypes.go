package main
import "fmt"
import "unsafe"
func main(){
var a bool=false
var b bool=true
fmt.Println("Boolean Datatypes\t\t",a,b)
var str string ="Vinay"
fmt.Println("String ",str)
var  inte int =9
fmt.Println("Integer",inte,"Size :",unsafe.Sizeof(inte))
var inte8 int8=0
fmt.Println("Integer8 ",inte8,"Size of int8 is ",unsafe.Sizeof(inte8))


//var inte16 int16=0
//fmt.Println("Integer16",int16,"Size of int16 is",unsafe.Sizeof(inte16))



// int int8 int16 int32 int64
// uint uint8 uint16 uint32 uint64 uintptr
//byte --> alias for uint8
// rune --> alias for int32
//float32 float64
//complex64 complex128
}
