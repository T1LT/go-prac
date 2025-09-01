package reverse

func Reverse(input string) string {
    runes := []rune(input)
	var str string

    for i := len(runes) - 1; i >= 0; i-- {
        str += string(runes[i])
    }

    return str
}
