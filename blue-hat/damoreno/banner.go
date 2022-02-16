package main

import "fmt"

type Banner struct {
	pixelMaps     []*PixelMap
	Height, Width int
	Data          [][]bool
}

func NewBanner(s string) (*Banner, error) {
	if len(s) == 0 {
		return nil, fmt.Errorf("input string cannot be empty")
	}

	var banner Banner
	var rows [RowsPerMap][]bool

	for _, ascii := range s {
		if pixelMap, ok := pixelMaps[ascii]; ok {
			for i, row := range pixelMap.Map {
				rows[i] = append(rows[i], row...)
				rows[i] = append(rows[i], false)
			}
		} else {
			return nil, fmt.Errorf("cannot find a pixel map for ascii '%c'", ascii)
		}
	}

	banner.Data = rows[:]
	banner.Height = RowsPerMap
	banner.Width = len(banner.Data[0])

	return &banner, nil
}

func (b Banner) String() string {
	var repr string

	for i := 0; i < b.Height; i++ {
		for j := 0; j < b.Width; j++ {
			if b.Data[i][j] {
				repr += "x"
			} else {
				repr += " "
			}
		}
		repr += "\n"
	}

	return repr
}
