package day23

import (
	"fmt"
	"slices"
	"sort"
	"strings"

	"github.com/mikeconroy/advent-of-code-24/utils"
)

func Run() (string, string) {
	input := utils.ReadFileIntoSlice("day23/input")
	return part1(input), part2(input)
}

type Computer struct {
	id        string
	connected map[string]bool
}

func parseInput(input []string) map[string]Computer {
	computers := make(map[string]Computer)

	for _, line := range input {
		if line == "" {
			continue
		}
		split := strings.Split(line, "-")
		c1 := split[0]
		c2 := split[1]

		if _, ok := computers[c1]; !ok {
			computers[c1] = Computer{c1, make(map[string]bool)}
		}
		if _, ok := computers[c2]; !ok {
			computers[c2] = Computer{c2, make(map[string]bool)}
		}

		computers[c1].connected[c2] = true
		computers[c2].connected[c1] = true

	}

	return computers

}

// Could be done with DFS -> Define a level
// Start at Computer (a) -> Visit each Connection -> Visit another Connection (increase level)
// -> When at Target Level check if origin node is connected
func part1(input []string) string {
	computers := parseInput(input)

	v1 := make(map[string]bool)
	sets := make(map[string][]string)

	for id, comp := range computers {
		v2 := make(map[string]bool)
		for connectionId := range comp.connected {
			if v1[connectionId] {
				continue
			}

			for _, val := range getSharedNodes(comp, computers[connectionId]) {
				if v1[val] || v2[val] {
					continue
				}
				key := id + connectionId + val
				sets[key] = []string{id, connectionId, val}
			}
			v2[connectionId] = true
		}
		v1[id] = true
	}

	count := 0
	for _, set := range sets {
		for _, id := range set {
			if strings.Index(id, "t") == 0 {
				count++
				break
			}
		}

	}

	return fmt.Sprint(count)
}

func getSharedNodes(c1 Computer, c2 Computer) []string {
	nodes := []string{}
	for compId := range c1.connected {
		if c2.connected[compId] {
			nodes = append(nodes, compId)
		}
	}
	return nodes
}

func validConnection(nodes []string, connId string, comps map[string]Computer) bool {
	for _, node := range nodes {
		if comps[node].connected[connId] != true {
			return false
		}
	}
	return true
}

func getNetwork(id string, nodes []string, connections map[string]bool, comps map[string]Computer, visited map[string]bool) []string {
	if len(connections) == 0 {
		return nodes
	}

	// Generate Connections between all current nodes
	nodes = append(nodes, id)
	longest := 0
	var longestPath []string
	for connId := range connections {
		if visited[connId] {
			continue
		}

		visited[connId] = true

		if slices.Contains(nodes, connId) {
			continue
		}

		// Could be changed so we only pass the intersections of all connections
		// instead of checking every ID every time.
		if validConnection(nodes, connId, comps) {
			result := getNetwork(connId, nodes, comps[connId].connected, comps, visited)
			if len(result) > longest {
				longest = len(result)
				longestPath = result
			}
		}
	}
	if longest == 0 {
		return nodes
	}
	return longestPath
}

func part2(input []string) string {
	comps := parseInput(input)
	longest := 0
	longestPath := []string{}

	for id, comp := range comps {
		visited := make(map[string]bool)
		visited[id] = true
		result := getNetwork(id, []string{}, comp.connected, comps, visited)
		if len(result) > longest {
			longest = len(result)
			longestPath = result
		}
	}
	sort.Strings(longestPath)
	res := strings.Join(longestPath, ",")
	return fmt.Sprint(res)
}
