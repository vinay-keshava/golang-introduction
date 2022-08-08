package main
import "fmt"
func main(){
	fmt.Printf("Hello World");
	var a="This is a String";
	var ab int=10;
	//here either the type or the expression is ignored
	fmt.Printf("%s %d",a,ab);
	//multiple declarations
	var b,c,d="String",34.343,'5';
	var marks int;
	fmt.Printf("%s\n%f\n%d",b,c,d);
	fmt.Printf("Enter the your marks");
	fmt.Scanln(&marks);
	fmt.Println(marks);
	fmt.Printf("",marks);
	

	//Short Variable Declaration is used mainly in local variables declaration
	vinay:="This is Vinay";
	fmt.Println(vinay);
	i,j:=2,3;//Multiple shorthand declarations of the variables i,j
	fmt.Println(i,j);
	i,j=j,i;
	fmt.Println(i,j);


	// This is for Variable Assignment with function calls;
	var dd,ee=hello();
	fmt.Println(dd,ee);
}
func hello()(int,int){
	a,b:=10,20;
	a=a+b;
	return a, b;
}
