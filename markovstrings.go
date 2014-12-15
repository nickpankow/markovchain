package markovChain

import (
	"bytes"
	"strconv"
)

func (mi MarkovItem) String() string{
	var buf bytes.Buffer

	buf.WriteString(mi.id)
	buf.WriteString(": [")
	for i,j:= range mi.connections{
		if i > 0{
			buf.WriteString(", ")
		}
		buf.WriteString(strconv.FormatFloat(j, 'f', 4, 64))
	}
	buf.WriteString("]")

	return buf.String()
}

func (mc MarkovChain) String() string{
	var buf bytes.Buffer

	buf.WriteString("Chain:\n")
	for _,y := range mc.chain{
		buf.WriteString(y.String())
		buf.WriteString("\n")
	}

	return buf.String()
}