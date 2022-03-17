package protein

import (
	"errors"
)

// Codon-Protein Map
var codonProteinMap = map[string]string{
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

// Errors
var ErrStop = errors.New("There is a STOP codon, so we need to stop right now!")
var ErrInvalidBase = errors.New("I'm sorry, but I don't know that codon :/")

func FromCodon(codon string) (string, error) {
	// search the codon in the Codon-Protein Map.
	protein, found := codonProteinMap[codon]
	if found {
		if protein == "STOP" {
			return "", ErrStop
		} else {
			return protein, nil
		}
	} else {
		return "", ErrInvalidBase
	}
}

func FromRNA(rna string) ([]string, error) {
	var proteinChain []string

	codons := splitByNumber(rna, 3)

	for _, codon := range codons {
		protein, error := FromCodon(codon)
		if error == ErrInvalidBase {
			return nil, error
		} else if error == ErrStop {
			return proteinChain, nil
		} else {
			proteinChain = append(proteinChain, protein)
		}
	}
	return proteinChain, nil
}

func splitByNumber(line string, num int) []string {
	var codons []string

	for i := 0; i < len(line); i++ {
		if i%num == 0 {
			if i+num < len(line) {
				codons = append(codons, line[i:i+num])
				// last codon group
			} else {
				codons = append(codons, line[i:])
			}
		}
	}
	return codons
}
