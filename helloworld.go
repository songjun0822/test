package main
import (
	"fmt"
)

var j int = 5

func test() func(){
	var i int = 10
	return func(){
		fmt.Printf("i, j: %d, %d\n", i, j)
	}
}

func main() {
	a := func () (func()){
		var i int = 10
		return func(){
			fmt.Printf("i, j: %d, %d\n", i, j)
		}
	}

	a()

}