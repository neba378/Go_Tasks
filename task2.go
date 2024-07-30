package main

import (
	"strings"
	"unicode"
)

// sub task 1
func word_count(word string) map[rune]int {
	count := map[rune]int{}
	word_array := []rune(word)
	for _, c := range word_array {
		count[c]++
	}

	return count
}

// sub task 2

func is_palindrome(word string) bool {
	word = strings.ReplaceAll(word, " ", "")
	
	word = strings.ToLower(word)
	word_array := []rune(word)
	length := len(word_array)

	arr := []rune{}

	for i:=0; i<length; i++{
		if !unicode.IsPunct(word_array[i]){
			arr = append(arr,word_array[i])
		}
	}
	length = len(arr)
	for i:=0; i<length; i++{
		if arr[i]!=arr[length-i-1]{
		return false
		}
	}
	return true
}

/* Uncomment the following to test the function*/
/* ALSO Do not forget to import fmt package after package main */

// func main(){
// 	// first subtask test
// 	result := word_count("Hello, A2SVans!")
//     for char, freq := range result {
//         fmt.Printf("Character: %c, Count: %d\n", char, freq)
//     }
// 	fmt.Println("\n-----------------------------\n")
// 	// Second sub task test
// 	result2 := is_palindrome("A man, a plan, a canal: Panama!")
// 	fmt.Println(result2)
// }