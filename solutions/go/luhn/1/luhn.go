package luhn

import (
    "strings"
    "strconv"
	"regexp"
)

func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")
    re := regexp.MustCompile(`^\d+$`)
    match := re.MatchString(id)
    
	if !match || len(id) <= 1 {
        return false
    }

    var total int
    flag := false
    
	for i := len(id) - 1; i >= 0; i-- {
        charInt, _ := strconv.Atoi(string(id[i]))
        
        if flag {
            doubled := charInt * 2
            
            if doubled > 9 {
                doubled -= 9
            }

            total += doubled
            flag = false
        } else {
			total += charInt
            flag = true
        }
    }

    return total % 10 == 0
}
