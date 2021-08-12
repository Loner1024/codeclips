package sort

import (
	"crypto/rand"
	"math/big"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeSort(t *testing.T) {
	n, err := rand.Int(rand.Reader, big.NewInt(12800))
	if err != nil {
		t.Fatal(err)
	}
	length := n.Int64()
	var a []int
	for i := 0; i < int(length); i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(12800))
		if err != nil {
			t.Fatal(err)
		}
		a = append(a, int(n.Int64()))
	}
	b := a
	sort.Ints(a)
	MergeSort(b, 0, len(a)-1)
	assert.Equal(t, b, a, "not equal")
}

func TestQuickSort(t *testing.T) {
	n, err := rand.Int(rand.Reader, big.NewInt(12800))
	if err != nil {
		t.Fatal(err)
	}
	length := n.Int64()
	var a []int
	for i := 0; i < int(length); i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(12800))
		if err != nil {
			t.Fatal(err)
		}
		a = append(a, int(n.Int64()))
	}
	b := a
	sort.Ints(a)
	QuickSort(b, 0, len(a)-1)
	assert.Equal(t, b, a, "not equal")
}
