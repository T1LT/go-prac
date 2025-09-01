// Package leap provides a function to work with leap years.
package leap

// IsLeapYear takes in a year and determines whether it is a leap year or not.
func IsLeapYear(year int) bool {
	if year % 4 == 0 {
        if year % 100 == 0 {
			if year % 400 == 0 {
                return true
            } else {
                return false
            }
        } else {
            return true
        }
    } else {
        return false
    }
}
