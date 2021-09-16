package main
import "fmt"


func someroutine(numbers chan int ){
	for i:=1; i<4; i++{
		numbers<-i
			
		}
		close(numbers)
		numbers<-5
	}
	
func main(){
	numbers := make(chan int, 3)
	go someroutine(numbers)
	for i:= range numbers {
		fmt.Println(i)
	}
}