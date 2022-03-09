package main

import "fmt"

func main() {
	graph := &Graph{nodes: make(map[NodeId]*Node)}
	graph.AddNode("1")
	graph.AddNode("2")
	graph.AddNode("3")
	graph.AddNode("4")
	graph.AddNode("5")

	graph.AddEdge("1", "2", 2)
	graph.AddEdge("1", "3", 3)
	graph.AddEdge("2", "4", 1)
	graph.AddEdge("3", "5", 2)
	graph.AddEdge("2", "5", 4)
	graph.AddEdge("4", "1", 9)
	graph.AddEdge("3", "4", 5)

	fmt.Println(graph.BreadthFirstPrint("1"))
	fmt.Println(graph.BreadthFirstPrint("2"))
	fmt.Println(graph.BreadthFirstPrint("3"))

	fmt.Println(graph.DepthFirstPrint("1"))
	fmt.Println(graph.DepthFirstPrint("2"))
	fmt.Println(graph.DepthFirstPrint("3"))

	for _, v := range [][]NodeId{{"1", "2"}, {"1", "3"}, {"1", "4"}, {"1", "5"}, {"0", "6"}, {"1", "7"}} {
		d1, e1 := graph.GetDistance(v[0], v[1])
		if e1 != nil {
			fmt.Println(e1)
		}
		d2, e2 := graph.GetMinDistance(v[0], v[1])
		if e2 != nil {
			fmt.Println(e2)
			continue
		}
		fmt.Printf("distance %v - %v: %v, min %v\n", v[0], v[1], d1, d2)
	}

	fmt.Println(graph.Kosaraju())
}
