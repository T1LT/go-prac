package strand

func ToRNA(dna string) string {
	var str string
    mapping := map[string]string{
        "G": "C",
        "C": "G",
        "T": "A",
        "A": "U",
    }
    
    for _, c := range dna {
        str += mapping[string(c)]
    }

    return str
}
