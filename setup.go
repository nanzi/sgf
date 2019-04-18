package sgf

import (
	"fmt"
)

func NewTree(size int) *Node {

	if size < 1 || size > 52 {
		panic(fmt.Sprintf("NewTree(): invalid size %v", size))
	}

	properties := make(map[string][]string)

	properties["GM"] = []string{"1"}
	properties["FF"] = []string{"4"}
	properties["SZ"] = []string{fmt.Sprintf("%d", size)}

	return NewNode(nil, properties)
}

func NewSetup(size int, black, white []Point, next_player Colour) *Node {

	if size < 1 || size > 52 {
		panic(fmt.Sprintf("NewSetup(): invalid size %v", size))
	}

	properties := make(map[string][]string)

	properties["GM"] = []string{"1"}
	properties["FF"] = []string{"4"}
	properties["SZ"] = []string{fmt.Sprintf("%d", size)}

	if next_player == WHITE {
		properties["PL"] = []string{"W"}
	} else if next_player == BLACK {
		properties["PL"] = []string{"B"}
	}

	if len(black) > 0 {
		properties["AB"] = []string{}
	}

	if len(white) > 0 {
		properties["AW"] = []string{}
	}

	for _, point := range black {
		properties["AB"] = append(properties["AB"], SGFFromPoint(point))
	}

	for _, point := range white {
		properties["AW"] = append(properties["AW"], SGFFromPoint(point))
	}

	return NewNode(nil, properties)
}