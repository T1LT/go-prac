package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	res := map[string]int{}

    for k, v := range in {
        for _, char := range v {
            res[strings.ToLower(string(char))] = k
        }
    }

    return res
}
