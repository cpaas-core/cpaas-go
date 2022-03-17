package protein

import (
	"errors"
	"sort"
)

const stop string = "STOP"
const codonLength int = 3

var ErrStop, ErrInvalidBase error = errors.New("Provided codon is a stop codon"), errors.New("Provided base is invalid")
var codonMap map[string]string = createCodonMap()

type job struct {
	id      int
	codon   string
	protein string
	err     error
}

func newJob(id int, codon string) *job {
	return &job{id: id, codon: codon}
}

func (j *job) addResult(protein string, err error) {
	j.protein = protein
	j.err = err
}

// interface to sort job slices
type idSortedJobs []job

func (a idSortedJobs) Len() int {
	return len(a)
}

func (a idSortedJobs) Less(i, j int) bool {
	return a[i].id < a[j].id
}

func (a idSortedJobs) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func createCodonMap() map[string]string {
	return map[string]string{
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
		"UAA": stop,
		"UAG": stop,
		"UGA": stop,
	}
}

/* sequential solution
func FromRNA(rna string) ([]string, error) {
	proteins := []string{}

	for begin, end := 0, 3; end <= len(rna); begin, end = end, end+3 {
		protein, err := FromCodon(rna[begin:end])

		if err != nil {
			if err == ErrInvalidBase {
				return []string{}, err
			}
			return proteins, nil
		}

		proteins = append(proteins, protein)
	}
	return proteins, nil
}*/

func worker(id int, codon string, results chan<- job) {
	work := newJob(id, codon)
	protein, err := FromCodon(work.codon)
	work.addResult(protein, err)
	results <- *work
}

func FromRNA(rna string) ([]string, error) {
	proteins := []string{}
	jobs := []job{}
	numJobs := len(rna) / codonLength
	results := make(chan job, numJobs)

	for begin, end := 0, codonLength; end <= len(rna); begin, end = end, end+codonLength {
		go worker(begin, rna[begin:end], results)
	}

	for i := 0; i < numJobs; i++ {
		result := <-results

		// if any worker returns ErrInvalidBase we can retunr immediately
		if result.err == ErrInvalidBase {
			return []string{}, result.err
		}
		jobs = append(jobs, result)
	}

	// sort results
	sort.Sort(idSortedJobs(jobs))
	for _, curJob := range jobs {
		if curJob.err == ErrStop {
			return proteins, nil
		}
		proteins = append(proteins, curJob.protein)
	}

	return proteins, nil
}

func FromCodon(codon string) (string, error) {
	protein, ok := codonMap[codon]

	if !ok {
		return "", ErrInvalidBase
	} else if protein == stop {
		return "", ErrStop
	}
	return protein, nil
}
