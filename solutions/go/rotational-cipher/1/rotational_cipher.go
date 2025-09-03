package rotationalcipher

import (
    "regexp"
    "unicode"
)

func RotationalCipher(plain string, shiftKey int) string {
	if shiftKey == 0 || shiftKey%26 == 0 {
        return plain
    }

    re := regexp.MustCompile(`[a-zA-Z]`)
    var res string
    for _, char := range plain {
        if re.MatchString(string(char)) {
            if unicode.IsUpper(char) {
                base := int('A')
                rotated := (int(char) - base + shiftKey) % 26 + base
                res += string(rune(rotated))
            } else {
                base := int('a')
                rotated := (int(char) - base + shiftKey) % 26 + base
                res += string(rune(rotated))
            }
        } else {
            res += string(char)
        }
    }
    
    return res
}
