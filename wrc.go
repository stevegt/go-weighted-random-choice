package wrc

import (
	"math/rand"
	"sort"
	// . "github.com/stevegt/goadapt"
)

type WeightedRandomChoice struct {
	elements    map[string]float64
	totalWeight float64
	// index is a list of the keys in elements, sorted by descending weight
	index []string
}

func New() WeightedRandomChoice {
	return WeightedRandomChoice{
		elements: make(map[string]float64),
	}
}

// AddElement adds an element to the WeightedRandomChoice
func (wrc *WeightedRandomChoice) AddElement(element string, weight float64) {
	wrc.elements[element] = weight
	// reset index and totalWeight so we'll recalculate next time we call GetRandomChoice
	wrc.index = nil
	wrc.totalWeight = 0
}

func (wrc *WeightedRandomChoice) AddElements(elements map[string]float64) {
	for element, weight := range elements {
		wrc.AddElement(element, weight)
	}
}

func (wrc *WeightedRandomChoice) recalc() {
	if wrc.index != nil && wrc.totalWeight != 0 {
		return
	}
	// build an unsorted index and calculate totalWeight
	wrc.index = make([]string, 0, len(wrc.elements))
	wrc.totalWeight = 0
	for name := range wrc.elements {
		wrc.index = append(wrc.index, name)
		wrc.totalWeight += wrc.elements[name]
	}
	// sort the index by descending weight
	sort.Slice(wrc.index, func(i, j int) bool {
		return wrc.elements[wrc.index[i]] > wrc.elements[wrc.index[j]]
	})
	return
}

func (wrc *WeightedRandomChoice) GetRandomChoice() string {
	wrc.recalc()
	value := rand.Float64() * wrc.totalWeight
	name := ""
	for _, name = range wrc.index {
		weight := wrc.elements[name]
		value -= weight
		if value <= 0 {
			break
		}
	}
	return name
}
