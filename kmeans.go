// kmeans is a package to calculate k-means clustering
// This package implements sequential functions
// so you can calculate k-means from stream of data
package kmeans

import(
	"math"
	"math/rand"
	"errors"
)

// Node represents a feature vector
type Node []float64

// Cluster is info about one cluster
type Cluster struct {
	Center Node
	Count int
}

type Kmeans struct {
	C []Cluster
	k int
	d Distance
	a float64
}

// Distance Function: To compute the distanfe between observations
type Distance func(Node, Node) float64


// Sum sums two vectors
func (n Node) Add(o Node) Node {
	r := make([]float64, len(n))
	for i, j := range o {
		r[i] = n[i]+j
	}
	return r
}

// Sub subtracts two vectors
func (n Node) Sub(o Node) Node {
	r := make([]float64, len(n))
	for i, j := range o {
		r[i] = n[i]-j
	}
	return r
}

// Mul multiplicates a vector with a scalar
func (n Node) Mul(s float64) Node {
	r := make([]float64, len(n))
	for i := range n {
		r[i] = n[i]*s
	}
	return r
}

func New(k int, d Distance, a float64) (Kmeans) {
	return Kmeans{make([]Cluster, k), k, d, a}
}

// Add adds a vector to the internal clusters on the frequently
// This function uses the count of elements in cluster as a weight
// See http://www.cs.princeton.edu/courses/archive/fall08/cos436/Duda/C/sk_means.htm
func (k Kmeans) Add(o Node) {
	add(o, k.C, k.d)
}

// Addf adds a vector to the internal clusters on the frequently
// This function uses constant pass during creation as a weight
// See http://www.cs.princeton.edu/courses/archive/fall08/cos436/Duda/C/sk_means.htm
func (k Kmeans) Addf(o Node) {
	addf(o, k.C, k.d, k.a)
}

// Sequential calculates clusters in sequential manner. It estimates clusters
func (k Kmeans) Sequential(o []Node) error {
	err := k.Seed(o)
	if err != nil {
		return err
	}
	for _, n := range o {
		k.Add(n)
	}
	return nil
}

func (k Kmeans) Seed(o []Node) error {
	if len(o) < k.k {
		return errors.New("There is not enough data")
	}
	seed(o, k.C, k.d, k.k)
	return nil
}

// Near finds the closest observation and returns the index and distance
func near(o Node, cs []Cluster, df Distance) (int, float64) {
	d := math.MaxFloat64
	i := -1

	for j, c := range cs {
		dd := df(o, c.Center)
		if dd < d {
			d = dd
			i = j
		}
	}
	return i, d
}

// Seed initializes clusters Instead of initializing randomly the seeds, make a sound decision of initializing
// Seed does not add 1 to cluster because we then proceed with Add() for each node
func seed(n []Node, cs []Cluster, df Distance, k int) {
	l := len(n)

	cs[0].Center = n[rand.Intn(l)]
	for i := 1; i < k; i++ {
		d2 := make([]float64, l)
		var sum float64
		for j, o := range n {
			_, d := near(o, cs[:i], df)
			d2[j] = d
			sum += d2[j]
		}
		target := rand.Float64() * sum
		j := 0
		for sum = d2[0]; sum < target; sum += d2[j] {
			j++
		}
		cs[i].Center = n[j]
	}
}

func add(o Node, cs []Cluster, df Distance) {
	// Get closest cluster
	i, _ := near(o, cs, df)
	// Update it's properties
	cs[i].Count += 1
	cs[i].Center = cs[i].Center.Add(o.Sub(cs[i].Center).Mul(1/float64(cs[i].Count)))
}

func addf(o Node, cs []Cluster, df Distance, a float64) {
	// Get closest cluster
	i, _ := near(o, cs, df)
	// Update it's properties
	cs[i].Count += 1
	cs[i].Center = cs[i].Center.Add(o.Sub(cs[i].Center).Mul(a))
}
