package main
import (	
	"fmt"
	"time"
) 
func main(){
	n:=time.Now()
	fmt.Println("This is time : ",n)

	t:=time.Date(2020,time.August,12,12,0,0,0,time.UTC)
	fmt.Println("This is time : ",t)

}