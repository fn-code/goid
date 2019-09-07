package goid

import (
	"testing"
)

var lastId GID

func TestNew(t *testing.T) {
	for i := 0; i < 100000; i++ {
		new := New()
		if new.String() == lastId.String() {
			t.Errorf("error generate unique new: %v - last: %v", new, lastId)
		}
		lastId = new
	}
}

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		New()
	}
}
