package greedy

import (
	nodegenerator "example.com/pci-graph-coloring-alg/NodeGenerator"
)

func GreedyPCI(nodes []nodegenerator.Node) {
	for i := range nodes {
		usedPCIs := make(map[int]bool)
		// Check the PCIs used by neighbors
		for _, neighborID := range nodes[i].Neighbors {
			for _, neighbor := range nodes {
				if neighbor.ID == neighborID {
					usedPCIs[neighbor.PCI] = true
				}
			}
		}

		// Assign the lowest available PCI
		pci := 1
		for usedPCIs[pci] {
			pci++
		}
		nodes[i].PCI = pci
	}
}
