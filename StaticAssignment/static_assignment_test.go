package static_assignment_test

import (
	"testing"

	nodegenerator "example.com/pci-graph-coloring-alg/NodeGenerator"
	static_assignment "example.com/pci-graph-coloring-alg/StaticAssignment"
)

// function to test StaticPCI

func TestStaticPCI(t *testing.T) {

	// Example nodes with neighbor relationships
	nodes := []nodegenerator.Node{
		{ID: 1, Degree: 3, Neighbors: []int{2, 3, 6}},
		{ID: 2, Degree: 1, Neighbors: []int{1}},
		{ID: 3, Degree: 2, Neighbors: []int{1}},
		{ID: 4, Degree: 2, Neighbors: []int{2, 3}},
		{ID: 5, Degree: 1, Neighbors: []int{4}},
		{ID: 6, Degree: 1, Neighbors: []int{4}},
	}
	//example pciPool
	pciPool := []int{1, 2, 3}

	// Apply static PCI assignment
	static_assignment.StaticPCI(nodes, 3)

	// Log the assigned PCI for each node
	for _, node := range nodes {
		t.Logf("Node %d assigned PCI: %d", node.ID, node.PCI)
	}

	// Check if each node has been assigned a PCI, beginning from the start again after list has been used up
	for i, node := range nodes {
		for j := i + 1; j < len(pciPool); j++ {
			if node.PCI == nodes[j].PCI {
				t.Errorf("Node %d and Node %d have the same PCI: %d", node.ID, nodes[j].ID, node.PCI)
			}
		}
	}

}
