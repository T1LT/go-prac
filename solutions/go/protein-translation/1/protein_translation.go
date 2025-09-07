package protein

import "errors"

var (
    ErrStop        = errors.New("STOP")
    ErrInvalidBase = errors.New("Invalid")
)

func FromRNA(rna string) ([]string, error) {
    res := []string{}
    if len(rna)%3 != 0 {
        return res, ErrInvalidBase
    }

    for i := 0; i < len(rna); i += 3 {
        codon := rna[i : i+3]
        protein, err := FromCodon(codon)
        if err == ErrStop {
            return res, nil
        } else if err == ErrInvalidBase {
            return res, ErrInvalidBase
        }
        res = append(res, protein)
    }
    return res, nil
}

func FromCodon(codon string) (string, error) {
    switch codon {
    case "AUG":
        return "Methionine", nil
    case "UUU", "UUC":
        return "Phenylalanine", nil
    case "UUA", "UUG":
        return "Leucine", nil
    case "UCU", "UCC", "UCA", "UCG":
        return "Serine", nil
    case "UAU", "UAC":
        return "Tyrosine", nil
    case "UGU", "UGC":
        return "Cysteine", nil
    case "UGG":
        return "Tryptophan", nil
    case "UAA", "UAG", "UGA":
        return "", ErrStop
    default:
        return "", ErrInvalidBase
    }
}
