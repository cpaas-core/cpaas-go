package letter

import "sync"

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

func (fm FreqMap) merge(m FreqMap) {
	for k, v := range m {
		fm[k] += v
	}
}

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(text string) FreqMap {
	frequencies := FreqMap{}
	for _, r := range text {
		frequencies[r]++
	}
	return frequencies
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(texts []string) FreqMap {
	c := make(chan FreqMap)
	var wg sync.WaitGroup
	wg.Add(len(texts))

	go func() {
		defer close(c)
		wg.Wait()
	}()

	for _, text := range texts {
		go func(text string) {
			defer wg.Done()
			c <- Frequency(text)
		}(text)
	}

	frequencies := <-c
	for fm := range c {
		frequencies.merge(fm)
	}
	return frequencies
}
