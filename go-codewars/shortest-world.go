// shortest-word
//Instructions
//Output

//Simple, given a string of words, (s).

//String will never be empty and you do not need to account for different data types.
package main
import "strings"
// FindShort return the length of the shortest word
func FindShort(s string) int {
	words := strings.Fields(s)
	shortest := 0
	for i, word := range words {
		if i == 0 {
			shortest = len(word)
		} else {
			if len(word) < shortest {
				shortest = len(word)	
			}
		}
		
	}
	return shortest
  }