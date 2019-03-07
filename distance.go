package kmeans


import (
	"math"
)

// L1
func L1(a, b Node) (float64) {
	distance := 0.
	for i := range a {
		distance += math.Abs(a[i] - b[i])
	}
	return distance
}

// L2
func L2(a, b Node) (float64) {
	return math.Sqrt(L2s(a, b))
}

// L2 squared
func L2s(a, b Node) (float64) {
	distance := 0.
	for i := range a {
		distance += (a[i] - b[i])*(a[i] - b[i])
	}
	return distance
}

// Lp returns Lp norm
func Lp(p float64) (Distance) {
	return func (a, b Node) float64 {
		distance := 0.
		for i := range a {
			distance += math.Pow(math.Abs(a[i]-b[i]), p)
		}
		return math.Pow(distance, 1/p)
	}
}

// Lpw returns weighted Lp norm
// Yeah I know, no checks... use it with caution!
func Lpw(w Node, p float64) (Distance) {
	return func (a, b Node) float64 {
		distance := 0.
		for i := range a {
			distance += w[i] * math.Pow(math.Abs(a[i]-b[i]), p)
		}
		return math.Pow(distance, 1/p)
	}
}

// infinity norm distance (l_inf distance)
func ChebyshevDistance(a, b Node) (float64) {
	distance := 0.
	for i := range a {
		if math.Abs(a[i]-b[i]) >= distance {
			distance = math.Abs(a[i] - b[i])
		}
	}
	return distance
}

func HammingDistance(a, b Node) (float64) {
	distance := 0.
	for i := range a {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance
}

func BrayCurtisDistance(a, b Node) (float64) {
	n, d := 0., 0.
	for i := range a {
		n += math.Abs(a[i] - b[i])
		d += math.Abs(a[i] + b[i])
	}
	return n/d
}

func CanberraDistance(a, b Node) (float64) {
	distance := 0.
	for i := range a {
		distance += math.Abs(a[i]-b[i]) / (math.Abs(a[i]) + math.Abs(b[i]))
	}
	return distance
}
