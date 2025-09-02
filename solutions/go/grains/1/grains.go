package grains

import "errors"

func Square(number int) (uint64, error) {
	if number < 1 || number > 64 {
        return 0, errors.New("invalid square number")
    }

    return uint64(1) << (number - 1), nil
}

func Total() uint64 {
	var sum uint64
    
    for i := 1; i <= 64; i++ {
        val, _ := Square(i)
        sum += val
    }

    return sum
}
