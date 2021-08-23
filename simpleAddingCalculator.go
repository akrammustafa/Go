package main
import (
		
	"fmt"	
) 
func main(){
	
	var input1,input2 int
	fmt.Print("Enter First Number :")
	fmt.Scanln(&input1)
	fmt.Print("Enter Second Number :")
	fmt.Scanln(&input2)
	var sum int =input1+input2
	fmt.Print("Sum is :",sum)
}