// Package triangle provides functions to identify the type of a triangle
package triangle

type Kind = string

const (
    // Pick values for the following identifiers used by the test program.
    NaT Kind = "NaT" // not a triangle
    Equ Kind = "Equ" // equilateral
    Iso Kind = "Iso" // isosceles
    Sca Kind = "Sca" // scalene
)

// KindFromSides determines what kind of triangle is given.
func KindFromSides(a, b, c float64) Kind {
	var k Kind

    switch {
    	case a <= 0.0 || b <= 0.0 || c <= 0.0:
            k = NaT
        case a + b < c || b + c < a || a + c < b:
            k = NaT
        case a == b && b == c:
            k = Equ
        case a == b || b == c || c == a:
            k = Iso
        default:
            k = Sca
    }
    
	return k
}
