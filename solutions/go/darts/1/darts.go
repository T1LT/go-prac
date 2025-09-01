package darts

import "math"

func Score(x, y float64) int {
	distanceFromOrigin := math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))

    switch {
        case distanceFromOrigin <= 1.0:
        	return 10
        case distanceFromOrigin > 1 && distanceFromOrigin <= 5:
        	return 5
        case distanceFromOrigin > 5 && distanceFromOrigin <= 10:
        	return 1
        default:
        	return 0
    }
}
