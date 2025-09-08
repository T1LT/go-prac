package anagram

import (
    "strings"
)

// charCount returns a frequency map of runes (characters) in the word, case-insensitive.
func charCount(word string) map[rune]int {
    counts := map[rune]int{}
    
    for _, r := range strings.ToLower(word) {
        counts[r]++
    }
    
    return counts
}

// equalCounts compares two rune frequency maps and returns true if they are equal.
func equalCounts(a, b map[rune]int) bool {
    if len(a) != len(b) {
        return false
    }
    
    for k, v := range a {
        if b[k] != v {
            return false
        }
    }
    
    return true
}

// Detect finds the candidate words that are anagrams of the subject.
func Detect(subject string, candidates []string) []string {
    subjectLower := strings.ToLower(subject)
    subjectCount := charCount(subject)

    res := []string{}
    for _, candidate := range candidates {
        candidateLower := strings.ToLower(candidate)
        if candidateLower == subjectLower {
            continue
        }
        
        candidateCount := charCount(candidate)
        if equalCounts(subjectCount, candidateCount) {
            res = append(res, candidate)
        }
    }
    
    return res
}
