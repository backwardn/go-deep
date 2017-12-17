package deep

import "math/rand"

type Example struct {
	Input    []float64
	Response []float64
}

type Examples []Example

func (e Examples) Shuffle() {
	for i := range e {
		j := rand.Intn(i + 1)
		e[i], e[j] = e[j], e[i]
	}
}

func (e Examples) Split(p float64) (first, second Examples) {
	for i := 0; i < len(e); i++ {
		if p > rand.Float64() {
			first = append(first, e[i])
		} else {
			second = append(second, e[i])
		}
	}
	return
}

func (e Examples) SplitN(size int) []Examples {
	res := make([]Examples, 0)
	for i := 0; i < len(e); i += size {
		res = append(res, e[i:min(i+size, len(e))])
	}
	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}