package main
import "fmt"

func main() {
	var y int = 55
	i := 90
	fmt.Println("Bonjour")
	fmt.Println(y)
	fmt.Println(i)
	for i:=0; i<8; i++{
		fmt.Println(i)
	}
	func adding( a int, b int){
		return a + b
	}
}