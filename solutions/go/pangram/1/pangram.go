package pangram

import (
    "regexp"
    "strings"
)

func IsPangram(input string) bool {
	re := regexp.MustCompile(`[^a-zA-Z]`)
    parsedStr := re.ReplaceAllString(input, "")
    charCounter := make(map[string]int)
    
    for c := 'a'; c <= 'z'; c++ {
        charCounter[string(c)] = 0
    }

    for _, char := range parsedStr {
        charCounter[strings.ToLower(string(char))]++
    }

    for _, v := range charCounter {
        if v == 0 {
            return false 
        }
    }

    return true
}
