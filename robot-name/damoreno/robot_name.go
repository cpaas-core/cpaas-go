package main

import (
	"fmt"
)

type Robot struct {
	Id string
}

var availableRobotNames = func() map[string]struct{} {
	names := make(map[string]struct{})

	for r := 'A'; r <= 'Z'; r++ {
		for m := 'A'; m <= 'Z'; m++ {
			for number := 0; number < 1000; number++ {
				robotName := fmt.Sprintf("%c%c%03d", r, m, number)
				names[robotName] = struct{}{}
			}
		}
	}

	return names
}()

func (r *Robot) Name() (string, error) {
	if r.Id != "" {
		return r.Id, nil
	}

	for name := range availableRobotNames {
		delete(availableRobotNames, name)
		r.Id = name
		return name, nil
	}

	return "", fmt.Errorf("no more names")
}

func (r *Robot) Reset() {
	r.Id = ""
}
