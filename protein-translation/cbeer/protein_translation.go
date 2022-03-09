package protein

import (
	"errors"
)

var ErrStop = errors.New("Invalid Stop")
var ErrInvalidBase = errors.New("Invalid codon")

func FromRNA(rna string) ([]string, error) {
	var proteins []string

	// Make sure the length is a multiple of 3
	if (len(rna) % 3) != 0 {
		return proteins, ErrInvalidBase
	}

	currentLen := 0
	currentStart := 0
	for i := range rna {
		if currentLen == 3 {
			protein, err := FromCodon(rna[currentStart:i])
			if err != nil {
				if err == ErrStop {
					return proteins, nil
				} else {
					return proteins, err
				}
			}
			proteins = append(proteins, protein)
			currentLen = 0
			currentStart = i
		}
		currentLen++
	}
	protein, err := FromCodon(rna[currentStart:])
	if err != nil {
		if err == ErrStop {
			return proteins, nil
		} else {
			return proteins, err
		}
	}
	proteins = append(proteins, protein)

	return proteins, nil
}

func FromCodon(codon string) (string, error) {
	var protein string

	switch codon {
	case "AUG":
		protein = "Methionine"
	case "UUU", "UUC":
		protein = "Phenylalanine"
	case "UUA", "UUG":
		protein = "Leucine"
	case "UCU", "UCC", "UCA", "UCG":
		protein = "Serine"
	case "UAU", "UAC":
		protein = "Tyrosine"
	case "UGU", "UGC":
		protein = "Cysteine"
	case "UGG":
		protein = "Tryptophan"
	case "UAA", "UAG", "UGA":
		return "", ErrStop
	default:
		return "", ErrInvalidBase
	}

	return protein, nil
}
