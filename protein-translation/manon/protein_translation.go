package protein

import (
	"errors"
)

var ErrStop = errors.New("found end of sequence codon")
var ErrInvalidBase = errors.New("found invalid base")

// FromRNA extracts the sequence of proteins from the rna string
func FromRNA(rna string) ([]string, error) {
	var proteins []string
	// There is no Map method for slices. See https://github.com/golang/go/discussions/47203#discussioncomment-1034432
	for i := 0; i < len(rna); i += 3 {
		protein, err := FromCodon(rna[i : i+3])
		if errors.Is(err, ErrStop) {
			break
		}
		if err != nil {
			return nil, err
		}
		proteins = append(proteins, protein)
	}
	return proteins, nil
}

// FromCodon extracts the specific protein of a codon triplet
func FromCodon(codon string) (string, error) {
	// Switch is slightly more efficient than mapping
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
