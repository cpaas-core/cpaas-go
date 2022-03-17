package protein

import (
	"errors"
)

var (
	ErrStop        error = errors.New("STOP")
	ErrInvalidBase error = errors.New("Invalid codon")
)

func FromRNA(rna string) ([]string, error) {
	ret := make([]string, 0)
	if len(rna) == 0 || len(rna)%3 != 0 {
		return nil, ErrInvalidBase
	}
	for i := 0; i < len(rna); i += 3 {
		protein, err := FromCodon(rna[i : i+3])
		if err == ErrStop {
			break
		}
		if err != nil {
			return nil, err
		}
		ret = append(ret, protein)
	}
	return ret, nil
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
	}
	return "", ErrInvalidBase
}
