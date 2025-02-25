package kpi_test

import (
	"testing"

	nodegenerator "example.com/pci-graph-coloring-alg/NodeGenerator"
	kpi "example.com/pci-graph-coloring-alg/kpi"
)

// Test for CountConflicts function
func TestCountConflicts(t *testing.T) {
	nodes := []nodegenerator.Node{
		{ID: 1, Degree: 1, PCI: 1, Neighbors: []int{2}}, // Conflict with Node 2
		{ID: 2, Degree: 1, PCI: 1, Neighbors: []int{1}}, // Conflict with Node 1
		{ID: 3, Degree: 2, PCI: 2, Neighbors: []int{4}},
		{ID: 4, Degree: 1, PCI: 3, Neighbors: []int{3}},
	}

	//expecting 1 conflict between Nodes 1 and 2
	expectedConflicts := 1
	actualConflicts := kpi.CountConflicts(nodes)

	if actualConflicts != expectedConflicts {
		t.Errorf("Expected %d conflicts but got %d conflicts", expectedConflicts, actualConflicts)
	}
}

// Test for MaxPCI function
func TestMaxPCI(t *testing.T) {
	nodes := []nodegenerator.Node{
		{ID: 1, PCI: 5},
		{ID: 2, PCI: 3},
		{ID: 3, PCI: 10}, // Max PCI
		{ID: 4, PCI: 8},
	}

	expectedMaxPCI := 10
	actualMaxPCI := kpi.MaxPCI(nodes)

	if actualMaxPCI != expectedMaxPCI {
		t.Errorf("Expected max PCI to be %d, but got %d", expectedMaxPCI, actualMaxPCI)
	}
}

// Test for AveragePCI function
func TestAveragePCI(t *testing.T) {
	nodes := []nodegenerator.Node{
		{ID: 1, PCI: 4},
		{ID: 2, PCI: 2},
		{ID: 3, PCI: 6},
		{ID: 4, PCI: 8},
	}

	expectedAveragePCI := 5.0
	actualAveragePCI := kpi.AveragePCI(nodes)

	if actualAveragePCI != expectedAveragePCI {
		t.Errorf("Expected average PCI to be %.2f, but got %.2f", expectedAveragePCI, actualAveragePCI)
	}
}

// Test for TotalAdjustments function
func TestTotalAdjustments(t *testing.T) {
	initialNodes := []nodegenerator.Node{
		{ID: 1, PCI: 2},
		{ID: 2, PCI: 3},
		{ID: 3, PCI: 4},
	}

	finalNodes := []nodegenerator.Node{
		{ID: 1, PCI: 1}, //Adjustment made
		{ID: 2, PCI: 3},
		{ID: 3, PCI: 5}, // Adjustment made
	}

	expectedAdjustments := 2
	actualAdjustments := kpi.TotalAdjustments(initialNodes, finalNodes)

	if actualAdjustments != expectedAdjustments {
		t.Errorf("Expected %d adjustments, but got %d", expectedAdjustments, actualAdjustments)
	}
}
