package main

import (
	"fmt"
	"math"
)

type NodeId string
type Graph struct {
	nodes map[NodeId]*Node
}
type Node struct {
	key       NodeId
	neighbors map[NodeId]int
}

func (g *Graph) AddNode(key NodeId) {
	g.nodes[key] = &Node{key: key, neighbors: make(map[NodeId]int)}
}

func (g *Graph) GetNode(key NodeId) *Node {
	return g.nodes[key]
}

func (g *Graph) Exists(key NodeId) bool {
	_, ok := g.nodes[key]
	return ok
}

const InvalidNodeKey = "InvalidNodeKey"

func (g *Graph) AddEdge(from, to NodeId, distance ...int) error {
	if !g.Exists(from) || !g.Exists(to) {
		return fmt.Errorf("addEdge error: %v", InvalidNodeKey)
	}
	if len(distance) > 0 {
		g.nodes[from].neighbors[to] = distance[0]
	} else {
		g.nodes[from].neighbors[to] = 0
	}
	return nil
}

func (g *Graph) BreadthFirstPrint(startingPoint NodeId) []NodeId {
	visited := make(map[NodeId]bool)
	res := []NodeId{}
	queue := []NodeId{startingPoint}
	for len(queue) > 0 { //get first element of queue (X), add X to res, then add neighbors of X to queue
		current := queue[0]
		queue = queue[1:]
		if !visited[current] {
			visited[current] = true
			res = append(res, current)
		} else {
			continue
		}
		for neighbor := range g.nodes[current].neighbors {
			queue = append(queue, neighbor)
		}
	}
	return res
}

func (g *Graph) DepthFirstPrint(startingPoint NodeId) []NodeId {
	visited := make(map[NodeId]bool)
	res := []NodeId{}
	stack := []NodeId{startingPoint}
	for len(stack) > 0 { //pop stack (X), add X to res, then add neighbors of X to stack
		l := len(stack)
		current := stack[l-1]
		stack = stack[:l-1]
		if !visited[current] {
			visited[current] = true
			res = append(res, current)
		} else {
			continue
		}

		for neighbor := range g.nodes[current].neighbors {
			stack = append(stack, neighbor)
		}
	}
	return res
}

//Kosaraju's algorithm for finding strongly connected components (scc)
func (g *Graph) Kosaraju() [][]NodeId {
	//1 go through every element in graph by dfs recursively. then push onto stack (of course check visted)
	//2 transpose graph (reverse all edges)
	//3 pop from stack every components, after each dfs return a scc (of course check visted)

	//1
	stack, visited := []NodeId{}, make(map[NodeId]bool)
	for id := range g.nodes {
		g.DfsAndPushStack(id, visited, &stack)
	}

	//2
	transposedGraph := g.transpose()

	//3
	sccs := [][]NodeId{}
	visited = make(map[NodeId]bool)
	for len(stack) > 0 {
		l := len(stack) - 1
		current := stack[l]
		stack = stack[:l]
		if !visited[current] {
			sccs = append(sccs, transposedGraph.Dfs(current, visited))
		}
	}
	return sccs
}
func (g *Graph) Dfs(node NodeId, visited map[NodeId]bool) []NodeId {
	res := []NodeId{}
	if !visited[node] {
		res = append(res, node)
		visited[node] = true
		for neighbor := range g.nodes[node].neighbors {
			res = append(res, g.Dfs(neighbor, visited)...)
		}
	}
	return res
}

//recursive dfs. push node finishes dfs onto stack. this way, last node in a dfs will be put first, on so on
func (g *Graph) DfsAndPushStack(node NodeId, visited map[NodeId]bool, stack *[]NodeId) {
	if !visited[node] {
		visited[node] = true
		for neighbor := range g.nodes[node].neighbors {
			g.DfsAndPushStack(neighbor, visited, stack)
		}
		*stack = append(*stack, node)
	}
}

//transpose is reversing edge direction
func (g *Graph) transpose() *Graph {
	transposedGraph := &Graph{nodes: make(map[NodeId]*Node)}
	for node := range g.nodes {
		transposedGraph.AddNode(node)
	}
	for id, node := range g.nodes {
		for neighbor := range node.neighbors {
			from, to := id, neighbor
			transposedGraph.AddEdge(to, from)
		}
	}
	return transposedGraph
}

func (g *Graph) GetDistance(from, to NodeId) (int, error) {
	if !g.Exists(from) || !g.Exists(to) {
		return -1, fmt.Errorf("GetDistance error: %v, %v, %v", from, to, InvalidNodeKey)
	}
	if from == to {
		return 0, nil
	}
	if v, ok := g.nodes[from].neighbors[to]; ok {
		return v, nil
	}
	for i := range g.nodes[from].neighbors {
		d1, err := g.GetDistance(from, i)
		if err != nil {
			return -1, err
		}
		if d2, err := g.GetDistance(i, to); d2 != -1 {
			if err != nil {
				return -1, err
			}
			return d1 + d2, nil
		}
	}
	return -1, nil
}

const max = math.MaxInt64

func (g *Graph) GetMinDistance(from, to NodeId) (int, error) {
	if !g.Exists(from) || !g.Exists(to) {
		return -1, fmt.Errorf("GetMinDistance error: %v, %v, %v", from, to, InvalidNodeKey)
	}
	if from == to {
		return 0, nil
	}
	if v, ok := g.nodes[from].neighbors[to]; ok {
		return v, nil
	}
	minDistance := max
	for i := range g.nodes[from].neighbors {
		d1, err := g.GetMinDistance(from, i)
		if err != nil {
			return -1, err
		}
		if d2, err := g.GetMinDistance(i, to); d2 != -1 {
			if err != nil {
				return -1, err
			}
			minDistance = min(minDistance, d1+d2)
		}
	}
	if minDistance == max {
		return -1, nil
	}
	return minDistance, nil
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
