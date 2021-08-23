package main  
  
import "fmt"  
  
func main() {  
 var operation string  
 var input1, input2 int  
 fmt.Print("Please enter First number: ")  
 fmt.Scanln(&input1)  
 fmt.Print("Please enter Second number: ")  
 fmt.Scanln(&input2)  
 fmt.Print("Please enter Operator (+,-,/,%,*):")  
 fmt.Scanln(&operation)  
 output := 0  
 switch operation {  
 case "+":  
  output = input1 + input2  
  break  
 case "-":  
  output = input1 - input2  
 case "*":  
  output = input1 * input2  
 case "/":  
  output = input1 / input2  
 case "%":  
  output = input1 % input2  
 default:  
  fmt.Println("Invalid Operation")  
 }  
 fmt.Printf("%d %s %d = %d", input1, operation, input2, output)  
}  