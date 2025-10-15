package main

import(
	"fmt"
	"strings"
)

func charfreq(s string) map[rune]int{
	s=strings.ToLower(s)
	freq:=map[rune]int{}
	for _,ch:=range s{
		if ch!=' '{
			freq[ch]++
		}
	}
	return freq
}

func main(){

	for ch, count := range charfreq("Hello World") {
    fmt.Printf("%q -> %d\n", ch, count) // %q prints the character in quotes
	}
}