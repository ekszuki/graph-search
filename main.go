package main

import "fmt"

//               graph - representation
//
//                         1
//               |--------/|\-------|
//               |         |        |
//          |----2---|     3     |--4--|
//          |        |           |     |
//       |--5--|     6           7     8
//       |     |
//       9     10

type Node struct {
	Value    int
	Children []*Node
}

func createGraph() Node {
	return Node{
		// node - 1
		1,
		[]*Node{
			// node - 2
			{
				Value: 2,
				Children: []*Node{
					// node - 5
					{
						5,
						[]*Node{
							{
								9,
								nil,
							},
							{
								10,
								nil,
							},
						},
					},
					// node - 6
					{
						Value: 6,
					},
				},
			},
			// node - 3
			{
				Value: 3,
			},
			// node - 4
			{
				Value: 4,
				Children: []*Node{
					// node - 7
					{
						Value: 7,
					},
					// node - 8
					{
						Value: 8,
					},
				},
			},
		},
	}
}

func main() {
	graph := createGraph()

	a := []int{}
	d := graph.DepthFirstSearch(a)
	b := graph.BreadthFirstSearch(a)

	fmt.Printf("DepthFirstSearch   Results: => %v \n\n", d)
	fmt.Printf("BreadthFirstSearch Results: => %v \n\n", b)
}

func (n *Node) DepthFirstSearch(array []int) []int {
	array = append(array, n.Value)

	for _, child := range n.Children {
		array = child.DepthFirstSearch(array)
	}

	return array
}

func (n *Node) BreadthFirstSearch(array []int) []int {
	queue := []*Node{n}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		array = append(array, current.Value)
		queue = append(queue, current.Children...)
	}

	return array
}
