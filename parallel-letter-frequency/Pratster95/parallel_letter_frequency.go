package letter

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
	channel := make(chan FreqMap)
	for _, str := range l {
		go func(str string) {
			channel <- Frequency(str)
		}(str)

	}
	freqMap := FreqMap{}
	for range l {
		for k, v := range <-channel {
			freqMap[k] += v
		}
	}
	return freqMap

}
