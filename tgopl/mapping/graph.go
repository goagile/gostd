package main

import "fmt"

type Graph map[string]map[string]bool

func (g Graph) hasEdge(from, to string) bool {
	return g[from][to]
}

func (g Graph) addEdge(from, to string) {
	if _, ok := g[from]; !ok {
		g[from] = make(map[string]bool)
		g[from][to] = true
	}
}

func main() {

	g := make(Graph)
	fmt.Println("graph:", g)
	fmt.Println("a--b:", g.hasEdge("a", "b"))

	g.addEdge("a", "b")
	fmt.Println("graph:", g)
	fmt.Println("a--b:", g.hasEdge("a", "b"))

}
