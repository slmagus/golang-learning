package main
import "strings"
import "fmt"
//i increments the value (initially 0)
//d decrements the value
//s squares the value
//o outputs the value into the return array

// Parse the data
func Parse(data string) []int{
	commands := strings.Split(data,"")
	
	value := 0
	finalvalues := make([]int, value)
	for _, command := range commands {
		switch command {
			case "i":
				value = increment(value)
			case "d":
				value = decrement(value)
			case "s":
				value = square(value)
			case "o":	
			finalvalues = append(finalvalues, value)
		}
	}
	
	return finalvalues
}

func increment(value int) int {
	return value + 1
}

func decrement(value int) int {
	return value - 1
}
func square(value int) int {
	return value * value	
}
func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
func main1(){
	data := "iiisdoso"
	foo := Parse(data)
	fmt.Println(foo)
}