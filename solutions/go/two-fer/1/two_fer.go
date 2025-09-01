// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package twofer should have a package comment that summarizes what it's about.
// Package twofer allows you to share items with others.
package twofer

import "fmt"

// ShareWith takes in a name and allows you to share.
func ShareWith(name string) string {
    var parsedName string
    
    if name == "" {
        parsedName = "you"
    } else {
        parsedName = name
    }
    
	return fmt.Sprintf("One for %s, one for me.", parsedName)
}
