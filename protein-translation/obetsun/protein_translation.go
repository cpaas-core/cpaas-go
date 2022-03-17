package protein

import (
	"errors"
)

/*
Codon                 | Protein
:---                  | :---
AUG                   | Methionine
UUU, UUC              | Phenylalanine
UUA, UUG              | Leucine
UCU, UCC, UCA, UCG    | Serine
UAU, UAC              | Tyrosine
UGU, UGC              | Cysteine
UGG                   | Tryptophan
UAA, UAG, UGA         | STOP
*/

var ErrStop = errors.New("Error: Stop")
var ErrInvalidBase = errors.New("Error: Invalid Base")

var codonsToRNA = map[string]string{
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

func FromRNA(rna string) ([]string, error) {
	var res = make([]string, 0)
	for i := 0; i < len(rna); i += 3 {
		cod := rna[i : i+3]
		prot, err := FromCodon(cod)
		if err == nil {
			res = append(res, prot)
		} else if err == ErrInvalidBase {
			return nil, err
		} else {
			return res, nil
		}

	}

	return res, nil
}

func FromCodon(codon string) (string, error) {
	if codonsToRNA[codon] == "STOP" {
		return "", ErrStop
	}

	if codonsToRNA[codon] == "" {
		return "", ErrInvalidBase
	}
	return codonsToRNA[codon], nil
}
