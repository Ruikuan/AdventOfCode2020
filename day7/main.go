package main

import (
	"fmt"
	"strings"
)

type colorBag struct {
	color      string
	containers []*colorBag
}

func traverse(node *colorBag, containerColors map[string]bool) {
	if containerColors[node.color] {
		return
	}
	containerColors[node.color] = true
	for _, c := range node.containers {
		traverse(c, containerColors)
	}
}

func part1() {
	bagMap := map[string]*colorBag{}
	for _, line := range strings.Split(input, "\n") {

		idx := strings.Index(line, " bags")
		parentColor := line[:idx]
		if bagMap[parentColor] == nil {
			bagMap[parentColor] = &colorBag{color: parentColor}
		}
		container := bagMap[parentColor]
		line = line[strings.Index(line, "contain")+7:]
		for _, childLine := range strings.Split(line, ",") {
			childLine = strings.Trim(childLine, " ")
			if childLine != "" {
				idx = strings.Index(childLine, " ")
				childLine = childLine[idx+1:]
				idx = strings.Index(childLine, " bag")
				childColor := childLine[:idx]
				if bagMap[childColor] == nil {
					bagMap[childColor] = &colorBag{color: childColor}
				}
				containee := bagMap[childColor]
				containee.containers = append(containee.containers, container)
			}
		}
	}

	containerColors := map[string]bool{}
	traverse(bagMap["shiny gold"], containerColors)
	fmt.Println(len(containerColors) - 1)
}

func main() {
	part1()

}
