package main
import (
       "fmt"
       "bufio"       
	   "os"

) 


func main(){
	reader :=bufio.NewReader(os.Stdin)
	fmt.Print("Enter :")
	input,_:=reader.ReadString('\n')
	fmt.Print("You enter :",input)

}