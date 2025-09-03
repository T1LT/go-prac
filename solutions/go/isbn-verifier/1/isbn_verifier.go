package isbn

import (
    "strings"
    "regexp"
    "strconv"
)

func IsValidISBN(isbn string) bool {
	parsed := strings.ReplaceAll(isbn, "-", "")
    if len(parsed) != 10 {
        return false
    }
    
	re := regexp.MustCompile(`^\d{9}[0-9X]$`)
    if !re.MatchString(parsed) {
        return false
    }

	var sum int
    for i, char := range parsed {
        if i == 9 && char == 'X' {
            sum += 10
        } else {
            val, _ := strconv.Atoi(string(char))
        	sum += val * (10 - i)
        }
    }

    return sum % 11 == 0
}
