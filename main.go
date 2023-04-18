package main

import (
  "fmt"
  "strings"
  "flag"
)

func main (){
	word := flag.String("word", "Abba", "word")	
	fmt.Print(isPalindrom(*word))

}	

func isPalindrom (some string) bool {
	l :=strings.TrimSpace(strings.ToLower(some))
    for i:=0; i < len(l)/2; i++ {
    	if (l[i] != l[len(l) -1 - i])  {
    		return false
    	}
  	}
  	return true
}