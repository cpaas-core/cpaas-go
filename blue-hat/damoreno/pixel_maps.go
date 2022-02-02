package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"strings"
)

const ColumnsPerMap = 5
const RowsPerMap = 7

var (
	//go:embed ascii_maps
	asciiMaps string

	pixelMaps = func() map[rune]PixelMap {
		maps := make(map[rune]PixelMap)

		scanner := bufio.NewScanner(strings.NewReader(asciiMaps))
		scanner.Split(bufio.ScanLines)

		for scanner.Scan() {
			pixelMap, err := loadPixelMap(scanner.Text())

			if err == nil {
				maps[pixelMap.ASCII] = *pixelMap
			}
		}

		return maps
	}()
)

type PixelMap struct {
	ASCII rune
	Map   [][]bool
}

func loadPixelMap(mapData string) (*PixelMap, error) {
	var pixelMap PixelMap

	mapFields := strings.Split(mapData, " ")
	if len(mapFields) != 2 || len(mapFields[0]) != 1 || len(mapFields[1]) != ColumnsPerMap*RowsPerMap {
		return nil, fmt.Errorf("invalid map data")
	}

	pixelMap.ASCII = []rune(mapFields[0])[0]

	var row []bool
	for i, character := range strings.TrimSpace(mapFields[1]) {
		if i != 0 && i%ColumnsPerMap == 0 {
			pixelMap.Map = append(pixelMap.Map, row)
			row = []bool{}
		}

		row = append(row, character == '1')
	}

	pixelMap.Map = append(pixelMap.Map, row)

	return &pixelMap, nil
}
