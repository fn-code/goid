package goid

import (
	"math/rand"
	"time"
)

const (
	// Base64 web-safe chars, but ordered by ASCII.
	charID = "-0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz"
)

// RandID is convert byte to string
type RandID string

var lastTime int64
var lastID [12]int

// New generate 20 random id
func New() RandID {
	var randNumber *rand.Rand
	var id [20]byte

	randNumber = rand.New(rand.NewSource(time.Now().UnixNano()))
	timeMs := time.Now().UTC().UnixNano() / 1e6
	if timeMs == lastTime {
		// incremen last rand id
		lastID = inLastRandID(lastID)
	} else {
		for i := 0; i < len(lastID); i++ {
			lastID[i] = randNumber.Intn(64)
		}
	}

	lastTime = timeMs
	// generate last 12 byte of id
	for i := 0; i < 12; i++ {
		id[19-i] = charID[lastID[i]]
	}

	// Genererate first 8 byte of id
	for i := 7; i >= 0; i-- {
		n := int(timeMs % 64)
		id[i] = charID[n]
		timeMs = timeMs / 64
	}
	return RandID(id[:])
}

func (s RandID) String() string {
	return string(s)
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
