package isogram

import "strings"

func IsIsogram(word string) bool {
	tracker := map[string]int{}

    for _, c := range word {
		char := strings.ToLower(string(c))
        
        if char != " " && char != "-" {
            _, prs := tracker[char]
            
            if prs {
                return false
            } else {
                tracker[char] = 0
            }
        }
    }

    return true
}
