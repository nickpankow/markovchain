package markovChain

import (
	"fmt"
	"errors"
)

type MarkovItem struct {
	id string
	connections []float64
}

type MarkovChain struct {
	chain []MarkovItem
}

func Create (ids []string, items [][]float64) (chain *MarkovChain, err error){
	if len(ids) != len(items){
		return nil, errors.New("Number of ids does not match the number of probability matrices.")
	}
	mc := MarkovChain{make([]MarkovItem, len(items), len(items))}
	for x,y := range items{
		mi := MarkovItem{ids[x], y}
		mc.chain[x] = mi
	}
	fmt.Println(mc)
	return &mc, nil
}

func (mc *MarkovChain) Next() (s string){
	return ""
}