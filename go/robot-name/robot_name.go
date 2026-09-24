package robotname

import (
	"errors"
	"fmt"
	"math/rand"
)

type Robot struct {
	name string
}

const MaxNames = 26 * 26 * 10 * 10 * 10

var (
	allNames      = initNames()
	nextNameIndex = 0
)

func (r *Robot) Name() (string, error) {
	if r.name == "" {
		if nextNameIndex >= MaxNames {
			return "", errors.New("all possible names have been issued")
		}

		r.name = allNames[nextNameIndex]
		nextNameIndex++
	}

	return r.name, nil
}

func (r *Robot) Reset() {
	r.name = ""
}

func initNames() []string {
	names := make([]string, MaxNames)
	i := 0

	for l1 := 'A'; l1 <= 'Z'; l1++ {
		for l2 := 'A'; l2 <= 'Z'; l2++ {
			for n := range 1000 {
				names[i] = fmt.Sprintf("%s%s%03d", string(l1), string(l2), n)
				i++
			}
		}
	}

	rand.Shuffle(len(names), func(i, j int) {
		names[i], names[j] = names[j], names[i]
	})

	return names
}
