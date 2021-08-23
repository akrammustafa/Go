

package main


import (
	
	
	"fmt"

) 







func main(){
var colors [3] string
	colors[0]="Blue"
	colors[1]="Yellow"
	colors[2]="Black"

	fmt.Println(colors)
	
	fmt.Println(colors[2])


	var numbers = [6] int{1,2,3,4,5,6}
	fmt.Println(numbers)

	fmt.Println("number of color",len(colors))
	fmt.Println("number of numbers",len(numbers))

}