package markovChain

import (
	"testing"
)

func TestCreate(t *testing.T){
	mc, err := Create([]string{"a","b"},[][]float64{{0.5, 0.5}, {0.5, 0.5}})

	if mc == nil || err != nil{
		t.Errorf("MarkovChain is nil")
	}

	mc, err = Create(
			[]string{"a","b","c","d","e"},
			[][]float64{
				{0.2, 0.2, 0.2, 0.2, 0.2},
				{0.2, 0.2, 0.2, 0.2, 0.2},
				{0.2, 0.2, 0.2, 0.2, 0.2},
				{0.2, 0.2, 0.2, 0.2, 0.2},
				{0.2, 0.2, 0.2, 0.2, 0.2}})
	if mc == nil || err != nil{
		t.Errorf("MarkovChain is nil")
	}
}