package main

import "fmt"

func anagram(s1 string, s2 string) bool {
	m := make(map[rune]int)
	for _, r := range s1 {
		m[r]++
	}

	for _, r := range s2 {
		if _, ok := m[r]; ok {
			m[r]--
			if m[r] == 0 {
				delete(m, r)
			}
		} else {
			return false
		}
	}
	return len(m) == 0
}

func main() {
	fmt.Println(anagram("hello", "ollhe"), "= true")
	fmt.Println(anagram("hello", "olllhe"), "= false")
	fmt.Println(anagram("hello", "olhe"), "= false")
	fmt.Println(anagram("abc", "abcc"), "= false")
}
