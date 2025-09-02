// Package bob contains a function for interpretation
package bob

import "regexp"

// Hey transforms Bob's words
func Hey(remark string) string {
    question := regexp.MustCompile(`\?\s*$`)
    yelling := regexp.MustCompile(`^[^a-z]*[A-Z][^a-z]*$`)
    yellingQuestion := regexp.MustCompile(`^[^a-z]*[A-Z][^a-z]*\?\s*$`)
    silence := regexp.MustCompile(`^\s*$`)

    switch {
        case yellingQuestion.MatchString(remark):
        	return "Calm down, I know what I'm doing!"
        case question.MatchString(remark):
        	return "Sure."
        case yelling.MatchString(remark):
        	return "Whoa, chill out!"
        case silence.MatchString(remark):
        	return "Fine. Be that way!"
        default:
        	return "Whatever."
    }
}
