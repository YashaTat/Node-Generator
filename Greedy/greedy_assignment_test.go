package greedy_test

import (
	"testing"

	greedy "example.com/pci-graph-coloring-alg/Greedy"
	nodegenerator "example.com/pci-graph-coloring-alg/NodeGenerator"
)

func TestGreedyPCI(t *testing.T) {
	// Define test nodes with neighbors
	nodes := []nodegenerator.Node{
		{ID: 1, Neighbors: []int{2, 3}},
		{ID: 2, Neighbors: []int{1, 3}},
		{ID: 3, Neighbors: []int{1, 2, 4}},
		{ID: 4, Neighbors: []int{3}},
	}

	// Call the GreedyPCI function
	greedy.GreedyPCI(nodes)

	// Verify that no two neighbors have the same PCI
	for _, node := range nodes {
		for _, neighborID := range node.Neighbors {
			for _, neighbor := range nodes {
				if neighbor.ID == neighborID && neighbor.PCI == node.PCI {
					t.Errorf("PCI conflict detected between Node %d and Node %d", node.ID, neighborID)
				}
			}
		}
	}
}
