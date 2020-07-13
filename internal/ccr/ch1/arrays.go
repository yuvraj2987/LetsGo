package chapter1

import (
	"unicode"
)

func isUnique(str string) bool {
	if len(str) == 0 {
		return true
	}
	aMap := make(map[rune]bool)
	for _, word := range str {
		_, exist := aMap[word]
		if exist {
			return false
		} else {
			aMap[word] = true
		}

	}
	return true
}

func checkPermutation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	// s1 and s2 have same length
	aMap := make(map[rune]int)
	for _, ch1 := range s1 {
		_, exists := aMap[ch1]
		if exists {
			aMap[ch1]++
		} else {
			aMap[ch1] = 1
		}
	}

	for _, ch2 := range s2 {
		_, exists := aMap[ch2]
		if exists {
			aMap[ch2]--
		} else {
			// ch1 in s1 is not in s2
			return false
		}
	}
	// check all key vals are set to zero
	for _, v := range aMap {
		if v != 0 {
			return false
		}
	}
	return true
}

func IsPalinPermutation(s string) bool {
	if len(s) <= 0 {
		return false
	}

	rs := []rune(s)
	aMap := make(map[rune]int)
	for _, ch := range rs {
		if !unicode.IsSpace(ch) {
			ch = unicode.ToLower(ch)
			_, exists := aMap[ch]
			if exists {
				aMap[ch]++
			} else {
				aMap[ch] = 1
			}

		} // end of non space chars
	} // end of for loop

	even_chars := 0
	odd_chars := 0

	for _, v := range aMap {
		if v%2 == 0 {
			even_chars++
		} else {
			odd_chars++
		}
	}

	return odd_chars <= 1

}
