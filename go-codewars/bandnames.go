package main
import "fmt"
// Expect(bandNameGenerator(string("knife"))).Should(BeEquivalentTo("The Knife"))
// Expect(bandNameGenerator(string("tart"))).Should(BeEquivalentTo("Tartart"))

//Expected
 //   <string>: The Step-daughter
//to be equivalent to
//	<string>: The Step-Daughte
	
import "strings"
func bandNameGenerator(word string) string {
	hypenation := strings.Split(word,"-")
	newword := ""
	if len(hypenation) >= 2 {
		for index, part := range hypenation {
			charpart := strings.Split(part,"")
			capitalizedword := strings.Replace(word,charpart[0], strings.ToUpper(charpart[0]),1)
			
			if index < len(hypenation) -1 {
				newword = capitalizedword + "-"
			} else {
				newword = capitalizedword
			}
			
		}
		word = newword
	}
	characters := strings.Split(word,"")
	word = strings.Replace(word,characters[0], strings.ToUpper(characters[0]),1)
	if characters[0] == characters[len(word)-1] {
		bandname := word + word[1:]
		return bandname
	} 
	bandname := "The " + word
	return bandname
	}

func main(){
	fmt.Println(bandNameGenerator("step-daughter"))
}