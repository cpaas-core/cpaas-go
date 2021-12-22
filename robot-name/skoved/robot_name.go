package robotname

import "fmt"

// Define the Robot type here.
type Robot struct {
	Id string
}

// max three degit number (999) + 1
var max = 1000
var alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
var robotNames = func() map[string]struct{} {
	names := make(map[string]struct{})
	for _, c1 := range alphabet {
		for _, c2 := range alphabet {
			for i := 0; i < max; i++ {
				name := fmt.Sprintf("%c%c%03d", c1, c2, i)
				names[name] = struct{}{}
			}
		}
	}
	return names
}()

func (r *Robot) Name() (string, error) {
	if r.Id == "" {
		if len(robotNames) <= 0 {
			return "", fmt.Errorf("There are no more available robot names.")
		}
		for name := range robotNames {
			delete(robotNames, name)
			r.Id = name
			return name, nil
		}
	}
	return r.Id, nil
}

func (r *Robot) Reset() {
	r.Id = ""
}
