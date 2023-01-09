package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Robot struct {
	name string
}

func (r *Robot) Name() (string, error) {
	strTemp := make([]rune, 2)
	intTemp := make([]int, 3)
	allLetter := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	allNumber := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	rand.Seed(time.Now().UnixNano())

	if r.name != "" {
		return r.name, nil
	}

	for idx := range strTemp {
		strTemp[idx] = allLetter[rand.Intn(len(allLetter))]
	}
	for idx := range intTemp {
		intTemp[idx] = allNumber[rand.Intn(len(allNumber))]
	}

	r.name = string(strTemp) + fmt.Sprintf("%v%v%v", intTemp[0], intTemp[1], intTemp[2])
	return r.name, nil
}

func (r *Robot) Reset() {
	r.name = ""
}
