//Pointer Example Program 
package main
import "fmt"
func main(){
var x int=1;
var p *int;

//p is a pointer of Integer type
p=&x;
fmt.Println("The address where the variable x:",p,"\n Value of X is: \t",x);

s:="Vinay";
sptr:=&s;
fmt.Println(sptr,"\t <- Address of the String",s);

var m,n int=10,20;
m=m-n;
fmt.Println(&m==&n);

//Naked Return of Pointer:
var v=f()
fmt.Println(v,"<- Naked Return Pointer ");

//Passing Pointer through Argument for Updation:
fmt.Println(paasbyReference(&x),"Is updated value after pass by reference");
}
func f() *int{
	v:=1
	return &v;
}

//passing function for updation of value passed through parameters;
func paasbyReference(p *int) int {
	*p++;
	return *p;
}
