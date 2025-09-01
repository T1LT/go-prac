package raindrops

import "fmt"

func Convert(num int) string {
	res := ""
    
    if num % 3 == 0 {
        res += "Pling"
    }
    
    if num % 5 == 0 {
        res += "Plang"
    } 
    
    if num % 7 == 0 {
        res += "Plong"
    }

    if num % 3 != 0 && num % 5 != 0 && num % 7 != 0 {
    	res += fmt.Sprintf("%d", num)
    }

    return res
}
