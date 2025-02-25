package backtracking

import (
	nodegenerator "example.com/pci-graph-coloring-alg/NodeGenerator"
)

func BacktrackingPCI(nodes []nodegenerator.Node, current int, maxPCI int) bool {
	if current == len(nodes) {
		return true // All nodes assigned
	}

	for pci := 1; pci <= maxPCI; pci++ {
		if isPCIAvailable(pci, nodes[current].Neighbors, nodes) {
			nodes[current].PCI = pci

			// Recursively assign the next node
			if BacktrackingPCI(nodes, current+1, maxPCI) {
				return true
			}

			// If conflict, reset PCI and backtrack
			nodes[current].PCI = 0
		}
	}
	return false
}

// Helper function to check if PCI is available -- backtracking specific
func isPCIAvailable(pci int, neighborIDs []int, nodes []nodegenerator.Node) bool {
	for _, neighborID := range neighborIDs {
		for _, node := range nodes {
			if node.ID == neighborID && node.PCI == pci {
				return false
			}
		}
	}
	return true
}
