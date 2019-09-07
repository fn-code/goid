package goid

import (
	"math/rand"
	"time"
)

const (
	// Base64 web-safe chars, but ordered by ASCII.
	charID = "-0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz"
)

// GID is convert byte to string
type GID string

var lastTime int64

// New is used for generate 20 random id base on base64 web save chars
func New() GID {
	id := [20]byte{}
	timeMs := time.Now().UTC().UnixNano() / 1e6
	lastID := generateLastID(timeMs, lastTime)

	lastTime = timeMs
	// generate last 12 byte of id
	for i := 0; i < 12; i++ {
		id[(len(id)-1)-i] = charID[lastID[i]]
	}

	// Genererate first 8 byte of id
	for i := 7; i >= 0; i-- {
		n := int(timeMs % 64)
		id[i] = charID[n]
		timeMs = timeMs / 64
	}
	return GID(id[:])
}

func (s GID) String() string {
	return string(s)
}

var lastIDS [12]int

func generateLastID(timeMs, lastTime int64) [12]int {
	var randNumber *rand.Rand
	randNumber = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	if timeMs == lastTime {
		// incremen last rand id
		lastIDS = inLastRandID(lastIDS)
	} else {
		for i := 0; i < len(lastIDS); i++ {
			lastIDS[i] = randNumber.Intn(64)
		}
	}
	return lastIDS
}

func inLastRandID(r [12]int) [12]int {
	for i := 0; i < 12; i++ {
		r[i]++
		if r[i] < 64 {
			break
		}
		r[i] = 0
	}
	return r
}
