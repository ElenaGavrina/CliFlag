package main

import (
  "fmt"
  "strings"
  "flag"
)

func main (){
	typeOfAnimal := flag.String("type", "dog", "type")	
	name := flag.String("name", "Abba", "name")	
	fmt.Println(isPalindrom(*typeOfAnimal))
	fmt.Println(isPalindrom(*name))

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
