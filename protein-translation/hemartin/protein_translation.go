package protein

import (
	"errors"
)

var codons = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}
var ErrStop = errors.New("stop")
var ErrInvalidBase = errors.New("invalid base")

func FromRNA(rna string) ([]string, error) {
	nLoops := len(rna) / 3
	fullProtein := []string{}

	for i := 0; i < nLoops; i++ {
		codon := rna[i*3 : i*3+3]
		protein, err := FromCodon(codon)
		if err == ErrStop {
			break
		}

		if err == ErrInvalidBase {
			return []string{}, err
		}

		fullProtein = append(fullProtein, protein)
	}
	return fullProtein, nil
}

func FromCodon(codon string) (string, error) {
	protein, ok := codons[codon]
	if !ok {
		return "", ErrInvalidBase
	}

	if protein == "STOP" {
		return "", ErrStop
	}
	return protein, nil
}
