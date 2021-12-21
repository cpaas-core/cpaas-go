package raindrops

import (
	"sort"
	"strconv"
)

var rainFactors = map[int]string{
	3: "Pling",
	5: "Plang",
	7: "Plong",
}

func Convert(number int) string {
	rainSound := ""

	// keys have to be sorted bc maps iterate in a random order
	keys := make([]int, 0)
	for key := range rainFactors {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	for _, factor := range keys {
		if number%factor == 0 {
			rainSound += rainFactors[factor]
		}
	}
	if rainSound == "" {
		rainSound = strconv.Itoa(number)
	}
	return rainSound
}
