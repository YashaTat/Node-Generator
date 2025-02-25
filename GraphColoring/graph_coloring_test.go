package graph_coloring_test

import (
	"fmt"
	"testing"

	graph_coloring "example.com/pci-graph-coloring-alg/GraphColoring"
	nodegenerator "example.com/pci-graph-coloring-alg/NodeGenerator"
)

// Test for isPCIAvailable when PCI is not available
func TestIsPCIAvailableFalse(t *testing.T) {
	PCI := 8                              //PCI to check availability of
	usedPCIs := [6]int{2, 3, 11, 9, 8, 4} //a list of used PCIs that contains the PCI we are checking => should return false

	expectedAvailability := false // is not available
	actualAvailability := graph_coloring.IsPCIAvailable(PCI, usedPCIs[:])

	if expectedAvailability != actualAvailability {
		t.Errorf("expected %t (is in use) but got %t (is available)", expectedAvailability, actualAvailability)
	}

}

// Test for isPCIAvailable when PCI is available
func TestIsPCIAvailableTrue(t *testing.T) {
	PCI := 8                              //PCI to check availability of
	usedPCIs := [6]int{2, 3, 11, 9, 1, 4} //a list of used PCIs that does not contain the PCI we are checking => should return true

	expectedAvailability := true //is available
	actualAvailability := graph_coloring.IsPCIAvailable(PCI, usedPCIs[:])

	if expectedAvailability != actualAvailability {
		t.Errorf("expected %t (is in use) but got %t (is available)", expectedAvailability, actualAvailability)
	}

}

// Test for removePCIConflict function
func TestRemovePCIConflict(t *testing.T) {
	// Define a node with a PCI conflict with its neighbors
	node := nodegenerator.Node{
		ID:        1,
		PCI:       2, // Initial PCI conflicts with a neighbor
		Neighbors: []int{2, 3},
	}

	// Define neighbors with conflicting PCIs
	neighbors := []nodegenerator.Node{
		{ID: 2, PCI: 2, Neighbors: []int{1}}, // Conflict here
		{ID: 3, PCI: 3, Neighbors: []int{1}},
	}

	// Run the removePCIConflict function to resolve the conflict
	graph_coloring.RemovePCIConflict(&node, neighbors)

	// The expected PCI for the node after conflict resolution
	expectedPCI := 1 // Since 1 is the lowest available PCI

	if node.PCI != expectedPCI {
		t.Errorf("Expected PCI to be %d, but got %d", expectedPCI, node.PCI)
	}
}

// Test for assignPCI function
func TestAssignPCI(t *testing.T) {
	// Define a set of nodes with their neighbors
	nodes := []nodegenerator.Node{
		{ID: 1, Degree: 2, Pu: []int32{1, 2, 3}, Pn: map[int32]bool{1: true}, Neighbors: []int{2, 3}},
		{ID: 2, Degree: 1, Pu: []int32{1, 2, 3}, Pn: map[int32]bool{1: true}, Neighbors: []int{1}},
		{ID: 3, Degree: 1, Pu: []int32{1, 2, 3}, Pn: map[int32]bool{1: true}, Neighbors: []int{1}},
		//{ID: 4, Degree: 2, Pu: []int32{1, 2, 3}, Neighbors: []int{2, 3}},
		//{ID: 5, Degree: 1, Pu: []int32{1, 2, 3}, Neighbors: []int{4}},
	}

	// Run PCI assignment
	adjustments := graph_coloring.AssignPCI(nodes)

	// Print results
	fmt.Println("Results after PCI Assignment:")
	for _, node := range nodes {
		fmt.Printf("Node %v assigned PCI: %v\n", node.ID, node.PCI)
	}

	// Validate results
	for _, node := range nodes {
		if node.PCI == -1 {
			t.Errorf("Node %v does not have a valid PCI assigned", node.ID)
		}

		// Check for conflicts
		neighbors := graph_coloring.GetNeighborNodes(nodes, node.Neighbors)
		for _, neighbor := range neighbors {
			if node.PCI == neighbor.PCI {
				t.Errorf("Conflict detected between Node %v and Node %v", node.ID, neighbor.ID)
			}
		}
	}

	fmt.Printf("Total adjustments made: %v\n", adjustments)
}

// Test for getNeighborNodes function
func TestGetNeighborNodes(t *testing.T) {
	// Define nodes
	nodes := []nodegenerator.Node{
		{ID: 1, PCI: 2, Neighbors: []int{2, 3}},
		{ID: 2, PCI: 3, Neighbors: []int{1}},
		{ID: 3, PCI: 1, Neighbors: []int{1}},
	}

	// Get neighbors of Node 1
	neighborIDs := []int{2, 3}
	neighbors := graph_coloring.GetNeighborNodes(nodes, neighborIDs)

	// Check if we got the right neighbors
	if len(neighbors) != 2 {
		t.Errorf("Expected 2 neighbors, but got %d", len(neighbors))
	}

	if neighbors[0].ID != 2 || neighbors[1].ID != 3 {
		t.Errorf("Expected neighbors to be Node 2 and Node 3, but got Node %d and Node %d", neighbors[0].ID, neighbors[1].ID)
	}
}

//Test for function HasPCIConflict

func TestHasPCIConflict(t *testing.T) {
	// Define a node with PCI 2
	node := nodegenerator.Node{
		ID:        1,
		PCI:       2,
		Neighbors: []int{2, 3},
	}

	// Define neighbors
	neighbors := []nodegenerator.Node{
		{ID: 2, PCI: 2, Neighbors: []int{1}}, // Conflict here
		{ID: 3, PCI: 3, Neighbors: []int{1}},
	}

	// Check if there is a PCI conflict
	conflict := graph_coloring.HasPCIConflict(node, neighbors)

	// Expect conflict to be true
	if !conflict {
		t.Errorf("Expected conflict to be true, but got false")
	}

	// Modify node to have no conflict
	node.PCI = 1

	// Re-check for conflict
	conflict = graph_coloring.HasPCIConflict(node, neighbors)

	// Expect conflict to be false
	if conflict {
		t.Errorf("Expected conflict to be false, but got true")
	}

}
