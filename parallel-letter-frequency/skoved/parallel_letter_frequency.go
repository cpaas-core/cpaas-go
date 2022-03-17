package letter

import "sync"

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(l []string) FreqMap {
	freq := FreqMap{}
	freqResults := make(chan FreqMap)
	var wg sync.WaitGroup

	// waits for all workers to finish and close the channel
	go func() {
		wg.Wait()
		close(freqResults)
	}()

	for _, line := range l {
		wg.Add(1)

		// creates a line variable specific to this iteration
		line := line
		go func() {
			defer wg.Done()
			freqResults <- Frequency(line)
		}()
	}

	// this loop only terminates when freqResults is closed
	for fmap := range freqResults {
		for key, value := range fmap {
			totalValue, ok := freq[key]
			if ok {
				freq[key] = totalValue + value
			} else {
				freq[key] = value
			}
		}
	}

	return freq
}
