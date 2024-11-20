package api_test

import (
	"testing"
	"time"
)

type Foo struct {
	LastTouched time.Time
}

func (f *Foo) UpdateTouched() { // Methods
	f.LastTouched = time.Now()
}

func (f Foo) Error() string {
	return "no!"
}

func TestMutate(t *testing.T) {

	// In typescript
	// type Reader interface {
	// 	Size: number
	// 	Read(i: string ): string
	// }

	// slices, maps

	// io.Reader
	var f error = Foo{}

	var _ = f
}

func Mutate(v []int) { // Functions
	v[0] = 2
}
