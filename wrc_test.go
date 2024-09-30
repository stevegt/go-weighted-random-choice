package wrc

import (
	"math"
	"math/rand"
	"testing"

	. "github.com/stevegt/goadapt"
)

func TestWeightedRandomChoice(t *testing.T) {
	rand.Seed(1)

	// Add some items
	elements := map[string]float64{
		"item1": 1,
		"item2": 2,
		"item3": 3,
	}
	totalWeight := 0.0
	for _, weight := range elements {
		totalWeight += weight
	}

	// Create a new WeightedRandomChoice
	wrc := New()
	// Add items
	wrc.AddElements(elements)

	// Get many random choices
	got := make(map[string]float64)
	N := 10000
	for i := 0; i < N; i++ {
		item := wrc.GetRandomChoice()
		// Pl(item)
		_, ok := got[item]
		if !ok {
			got[item] = 1
		} else {
			got[item]++
		}
	}
	// Check the result
	for item, count := range got {
		weight := elements[item]
		expected := float64(weight) / float64(totalWeight)
		actual := float64(count) / float64(N)
		Pf("item: %s, expected: %f, actual: %f\n", item, expected, actual)
		if math.Abs(expected-actual) > 0.01 {
			t.Errorf("item: %s, expected: %f, actual: %f", item, expected, actual)
		}
	}

}
