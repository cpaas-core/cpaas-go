package protein

import "errors"

type codon struct {
  name string
  protein string
}

var ErrStop = errors.New("ErrStop")
var ErrInvalidBase = errors.New("ErrInvalidBase")

func BuildMap() map[string]string {

	Codon := make(map[string]string)

	Codon["AUG"] = "Methionine"

	Codon["UUU"] = "Phenylalanine"
	Codon["UUC"] = "Phenylalanine"

	Codon["UUA"] = "Leucine"
	Codon["UUG"] = "Leucine"

	Codon["UCU"] = "Serine"
	Codon["UCC"] = "Serine"
	Codon["UCA"] = "Serine"
	Codon["UCG"] = "Serine"

	Codon["UAU"] = "Tyrosine"
	Codon["UAC"] = "Tyrosine"

	Codon["UGU"] = "Cysteine"
	Codon["UGC"] = "Cysteine"
	
	Codon["UGG"] = "Tryptophan"

	Codon["UAA"] = "STOP"
	Codon["UAG"] = "STOP"
	Codon["UGA"] = "STOP"

	return Codon
}

func SliceRNA(rna string) []string {

	codon := ""
	protein := []string{}

	for i, char := range rna {
		codon += string(char)

		if i % 3 == 2 {
			protein = append(protein, codon)
			codon = ""

			continue
		}
	}
	return protein
}

func FromRNA(rna string) ([]string, error) {
	if len(rna) % 3 != 0 { 
		return nil, ErrInvalidBase
	}
	codons := SliceRNA(rna)
	data := BuildMap()

	protein := []string{}
	for _, codon := range codons { 
		if data[codon] == "STOP" {
			break
		}
		protein = append(protein, data[codon])
	}
	return protein, nil
}

func FromCodon(codon string) (string, error) {

	data := BuildMap()
	protein := data[codon]

	if len(protein) == 0 {
		return "", ErrInvalidBase
	}
	if protein == "STOP" {
		// why errors.New("ErrStop") makes it fail?
		return "", ErrStop
	}
	return protein, nil
}
