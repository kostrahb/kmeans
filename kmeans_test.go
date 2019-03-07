package kmeans

import(
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
)

func TestKmeans(t *testing.T) {
	filePath, err := filepath.Abs("data/iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// Best Distance for Iris is Canberra Distance
	k := New(3, CanberraDistance, 0.1)

	lines := strings.Split(string(content), "\n")

	seed := make([]Node, 3)
	for i, line := range lines[:3] {
		vector := strings.Split(line, ",")
//		label := vector[len(vector)-1]
		vector = vector[:len(vector)-1]
		n := make([]float64, len(vector))
		for jj := range vector {
			n[jj], err = strconv.ParseFloat(vector[jj], 64)
		}
		seed[i] = n
	}
	k.Seed(seed)

	for _, line := range lines[3:] {
		if line == "" {
			break
		}
		vector := strings.Split(line, ",")
//		label := vector[len(vector)-1]
		vector = vector[:len(vector)-1]
		n := make([]float64, len(vector))
		for jj := range vector {
			n[jj], err = strconv.ParseFloat(vector[jj], 64)
		}
		k.Addf(n)
	}

	// Check
	c := map[string][]Node{}
	for _, line := range lines {
		if line == "" {
			break
		}
		vector := strings.Split(line, ",")
		label := vector[len(vector)-1]
		vector = vector[:len(vector)-1]
		n := make([]float64, len(vector))
		for jj := range vector {
			n[jj], err = strconv.ParseFloat(vector[jj], 64)
		}
		c[label] = append(c[label], n)
	}

	for _, nodes := range c {
		n := make([]float64, len(nodes[0]))
		for i := 0; i<len(nodes[0]); i++ {
			for j := 0; j<len(nodes); j++ {
				n[i] += nodes[j][i]
			}
			n[i] = n[i]/float64(len(nodes))
		}
		fmt.Println(n)
	}

	fmt.Println(k)
/*	labels := Kmeans(irisData, 3, CanberraDistance, threshold)

	misclassifiedOnes := 0
	for ii, jj := range labels {
		if ii < 50 {
			if jj != 2 {
				misclassifiedOnes++
			}
		} else if ii < 100 {
			if jj != 1 {
				misclassifiedOnes++
			}
		} else {
			if jj != 0 {
				misclassifiedOnes++
			}
		}
	}
	fmt.Println(misclassifiedOnes)*/
}
