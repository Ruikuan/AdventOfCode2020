package main

import (
	"fmt"
	"strconv"
	"strings"
)

type containeeBag struct {
	count int
	bag   *colorBag
}
type colorBag struct {
	color      string
	containers []*colorBag
	containees []*containeeBag
}

func constructGraph() *colorBag {
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
				count, _ := strconv.Atoi(childLine[:idx])
				childLine = childLine[idx+1:]
				idx = strings.Index(childLine, " bag")
				childColor := childLine[:idx]
				if bagMap[childColor] == nil {
					bagMap[childColor] = &colorBag{color: childColor}
				}
				containee := bagMap[childColor]
				containee.containers = append(containee.containers, container)

				eeBag := &containeeBag{count: count, bag: containee}
				container.containees = append(container.containees, eeBag)
			}
		}
	}
	return bagMap["shiny gold"]
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
	shinyGoldBag := constructGraph()
	containerColors := map[string]bool{}
	traverse(shinyGoldBag, containerColors)
	fmt.Println(len(containerColors) - 1)
}

func traverseCount(bag *colorBag) int {
	count := 1
	for _, ee := range bag.containees {
		count += ee.count * traverseCount(ee.bag)
	}
	return count
}

func part2() {
	shinyGoldBag := constructGraph()
	count := traverseCount(shinyGoldBag) - 1 // -1 to exclude shiny gold bag itself
	fmt.Println(count)
}

func main() {
	part1()
	part2()
}
