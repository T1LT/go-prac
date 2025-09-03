package rotationalcipher

import "unicode"

func RotationalCipher(plain string, shiftKey int) string {
    if shiftKey == 0 || shiftKey % 26 == 0 {
        return plain
    }

    shiftKey = shiftKey % 26
    res := []rune{}

    for _, char := range plain {
        if unicode.IsUpper(char) {
            base := 'A'
            rotated := (char - base + rune(shiftKey)) % 26 + base
            res = append(res, rotated)
        } else if unicode.IsLower(char) {
            base := 'a'
            rotated := (char - base + rune(shiftKey)) % 26 + base
            res = append(res, rotated)
        } else {
            res = append(res, char)
        }
    }
    return string(res)
}