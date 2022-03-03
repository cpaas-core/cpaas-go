package letter

import (
	"sync"
)

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
	f := FreqMap{}
	freqValues := make(chan FreqMap)
	var wg sync.WaitGroup

	go func() {
		wg.Wait()
		close(freqValues)
	}()

	for _, line := range l {
		wg.Add(1)

		line := line
		go func() {
			defer wg.Done()
			freqValues <- Frequency(line)
		}()
	}

	for fmap := range freqValues {
		for key, value := range fmap {
			totalValue, ok := f[key]
			if ok {
				f[key] = totalValue + value
			} else {
				f[key] = value
			}
		}
	}
	return f
}
