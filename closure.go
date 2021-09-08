package main
import "fmt"

func outer() func(){
	count := 0
    return func(){
		count++
        fmt.Println("Now I am here ",count )
    }
}
func main(){
    outer() //nothing is printed
    outer()() // this will print "Now I am here q"
	f := outer()
	f() //Now I am here  1
	f() //Now I am here  1

}