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

func ConcWrapper(s string, m chan FreqMap) {
	m <- Frequency(s)   
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(l []string) FreqMap {
	var m chan FreqMap
    m = make(chan FreqMap)
    var res FreqMap
    res = make(map[rune]int)
    for _, v := range l { 
        go ConcWrapper(v, m)
        mm, ok := <-m
        if ok == true {
        	for r,f := range mm {
                res[r] += f     
            }
        }
    }
	return res
}
