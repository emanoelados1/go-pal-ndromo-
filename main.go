package main

import "fmt"

func isPalindrome(s string) bool {
    n := len(s)
    for i := 0; i < n/2; i++ {
        if s[i] != s[n-1-i] {
            return false
        }
    }
    return true
}

func main() {
	var possiblesNames = [5]string{ "arara", "radar", "ovo", "Jonathan", "Manu" }

	for _, name := range possiblesNames {
		if isPalindrome(name) {
			fmt.Println(name, "é um palíndromo.")
		} else {
			fmt.Println(name, "não é um palíndromo.")
		}
	}
}
