package main
import "fmt"
// import "time"

func getAintoB( ab chan int){
	// time.Sleep(2132321)
	fmt.Println("before ab")

	ab<-3*2
	fmt.Println("after ab")

}

func getCintoD( ab chan int){
	fmt.Println("before ab")

	ab<-5-2
	fmt.Println("after ab")
	
}


func main(){
	ab := make(chan int, 2)
	// cd := make(chan int)
	go getCintoD(ab)
	go getAintoB(ab)

	

	fmt.Println(<-ab)


}
// this is not efficient at channels.go here we are waiting for first print statement and as
// go getCintoD(ab) is after that this is kind of executing like a single thread