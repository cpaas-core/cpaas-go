// Package grep provides a Search function for the Grep Exercism exercise
package grep

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
)

// grepper provides a Grep style match function
type grepper struct {
	lineNums     bool
	onlyFileName bool
	ignoreCase   bool
	invert       bool
	wholeLine    bool
	multiFile    bool
	pattern      string
}

// newGrepper takes a list of string flags and returns a new grepper struct.
func newGrepper(flags []string, pattern string, multiFile bool) (*grepper, error) {
	matcher := &grepper{multiFile: multiFile}

	for _, flag := range flags {
		switch flag {
		case "-n":
			matcher.lineNums = true
		case "-l":
			matcher.onlyFileName = true
		case "-i":
			matcher.ignoreCase = true
		case "-v":
			matcher.invert = true
		case "-x":
			matcher.wholeLine = true
		default:
			return nil, fmt.Errorf("unrecognized flag: %s", flag)
		}
	}

	// to match the whole line pattern must start with '^' and end with '$'
	if matcher.wholeLine {
		if pattern[0] != '^' {
			pattern = "^" + pattern
		}
		if pattern[len(pattern)-1] != '$' {
			pattern += "$"
		}
	}

	// to ingore case in a regexp match pattern must start with (?i)
	if matcher.ignoreCase && (len(pattern) < 4 || pattern[:4] != "(?i)") {
		pattern = "(?i)" + pattern
	}
	matcher.pattern = pattern

	return matcher, nil
}

func (g *grepper) match(fileName string) ([]string, error) {
	matches := []string{}
	file, err := os.Open(fileName)
	if err != nil {
		return []string{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	regex, err := regexp.Compile(g.pattern)
	lineNum := 1
	for scanner.Scan() {
		line := scanner.Text()
		match := regex.MatchString(line)
		if g.invert {
			match = !match
		}
		if match {
			if g.lineNums {
				line = fmt.Sprintf("%d:%s", lineNum, line)
			}
			if g.multiFile {
				line = fmt.Sprintf("%s:%s", fileName, line)
			}
			if g.onlyFileName {
				line = fileName
				matches = append(matches, line)
				break
			}
			matches = append(matches, line)
		}
		lineNum++
	}
	return matches, nil
}

// test case expect results in the same order as file input so need id for ordering results
type job struct {
	id   int
	file string
}

type payload struct {
	id      int
	matches []string
}

// interface for sorting a slice
type idSortedPayloads []payload

func (a idSortedPayloads) Len() int {
	return len(a)
}

func (a idSortedPayloads) Less(i, j int) bool {
	return a[i].id < a[j].id
}

func (a idSortedPayloads) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func worker(id int, grep *grepper, jobs <-chan job, results chan<- payload) {
	for job := range jobs {
		matches, err := grep.match(job.file)
		if err != nil {
			panic(fmt.Sprintf("worker %d: Error matching in file %s: %s\n", id, job.file, err))
		}
		result := payload{id: job.id, matches: matches}
		results <- result
	}
}

func Search(pattern string, flags, files []string) []string {
	numJobs := len(files)
	jobs := make(chan job, numJobs)
	results := make(chan payload, numJobs)
	payloads := []payload{}
	lines := []string{}

	matcher, err := newGrepper(flags, pattern, len(files) > 1)
	if err != nil {
		fmt.Printf("Error creating grepper: %s\n", err)
		return []string{}
	}

	// start workers threads
	for w := 1; w <= numJobs; w++ {
		go worker(w, matcher, jobs, results)
	}

	for i, file := range files {
		jobs <- job{id: i, file: file}
	}

	for i := 0; i < numJobs; i++ {
		payloads = append(payloads, <-results)
	}

	// insert match results in order
	sort.Sort(idSortedPayloads(payloads))
	for _, result := range payloads {
		if len(result.matches) > 0 {
			lines = append(lines, result.matches...)
		}
	}
	return lines
}
