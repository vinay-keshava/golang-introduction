package main
import "fmt"
func main(){
var a float64;
fmt.Printf("Enter the temperature in Fahrenheit : ");
fmt.Scanln(&a);
calcCelsi(a);
}

func calcCelsi(a float64){
//c=f+32*(9/5)
fmt.Println(a+32*(9/5));

}
