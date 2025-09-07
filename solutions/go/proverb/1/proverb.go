// Package proverb has a function to build rhymes
package proverb

import "fmt"

// Proverb builds the full text from a slice of words.
func Proverb(rhyme []string) []string {
	res := []string{}
    
    if len(rhyme) == 0 {
        return res
    }

    if len(rhyme) == 1 {
        res = append(res, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
        return res
    }

    for i := 0; i < len(rhyme) - 1; i++ {
        res = append(res, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i + 1]))
    }

    res = append(res, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
    return res
}
