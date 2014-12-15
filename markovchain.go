package markovChain

import (
	"errors"
	"math/rand"
	"time"
)

var randSource *rand.Rand

type MarkovItem struct {
	id string
	connections []float64
}

type MarkovChain struct {
	chain []MarkovItem
	state *MarkovItem
}

func Create (ids []string, items [][]float64) (chain *MarkovChain, err error){
	if len(ids) != len(items){
		return nil, errors.New("Number of ids does not match the number of probability matrices.")
	}

	mc := MarkovChain{make([]MarkovItem, len(items), len(items)), nil}
	for x,y := range items{
		mi := MarkovItem{ids[x], y}
		mc.chain[x] = mi
	}
	mc.state = &mc.chain[0]

	return &mc, nil
}

func (mc *MarkovChain) Next() (s string){
	item := mc.state
	x := item.Next()
	mc.state = &mc.chain[x]
	return mc.state.id
}

func (mi *MarkovItem) Next() (int){
	if randSource == nil{
		randSource = rand.New(rand.NewSource(time.Now().UnixNano()))
	}
	rVal := randSource.Float64()
	p := 0.0
	for i,x := range mi.connections{
		if rVal < (float64(x) + p){
			return i
		}
		p += x
	}
	return len(mi.connections) - 1
}