package main
import "fmt"

func getAintoB( ab chan int){
	ab<-3*2
}

func getCintoD( ab chan int){
	ab<-5-2
	
}


func main(){
	ab := make(chan int)
	// cd := make(chan int)
	go getAintoB(ab)
	fmt.Println(<-ab)

	go getCintoD(ab)

	

	fmt.Println(<-ab)


}
// this is not efficient at channels.go here we are waiting for first print statement and as
// go getCintoD(ab) is after that this is kind of executing like a single thread