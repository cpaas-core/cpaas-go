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
func ConcurrentFrequency(texts []string) FreqMap {
	var wg sync.WaitGroup

	maps := make(chan FreqMap, len(texts))

	for _, text := range texts {
		wg.Add(1)
		go func(text string) {
			defer wg.Done()
			maps <- Frequency(text)
		}(text)
	}

	wg.Wait()
	close(maps)

	sumMap := FreqMap{}
	for freqMap := range maps {
		for key, value := range freqMap {
			if finalValue, ok := sumMap[key]; ok {
				sumMap[key] = finalValue + value
			} else {
				sumMap[key] = value
			}
		}
	}

	return sumMap
}
