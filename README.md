# golang-introduction
Getting started with GO programming language.
# https://boot.dev/learn/learn-golang

# Execution of Golang 
go run filename.go

go is Strongly Typed language which means that variable can have only single type string cannot be converted to int;

The package “main” tells the Go compiler that the package should compile as an executable program instead of a shared library. 

# Basic Data types available
bool string int int8 int16 int32 int64 uint uint8 uint16 uint32 uint64

byte(alias for uint8) rune(alias for int32) float32 float64 complex64 complex128


# Declaration 
 var name type= expression

either the type or the expression is igmored 

# Short Variable Declaration
name:=expression  -> mainly used for local variables;
Here the type of the variable is automatically detected 

Examples are:
hello:=20;     Integer type

helloString:="String variable";

helloFloat:=23.43;

: is the declaration where = is the assignment

	var dd,ee=hello();
	fmt.Println(dd,ee);

func hello()(int,int){
	a,b:=10,20;
	a=a+b;
	return a, b;
}

