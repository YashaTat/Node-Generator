package static_assignment

import (
	nodegenerator "example.com/pci-graph-coloring-alg/NodeGenerator"
)

// Static PCI Assignment (No Conflict Resolution)
func StaticPCI(nodes []nodegenerator.Node, maxUsablePCI int) {
	// Predefined static PCI values assigned to each node sequentially
	pciPool := make([]int, maxUsablePCI)
	for i := 0; i < maxUsablePCI; i++ {
		pciPool[i] = i + 1 // Fixed set of PCIs to assign
	}
	pciIndex := 0

	// Assign static PCIs
	for i := range nodes {
		// Assign PCI from predefined list in a round-robin fashion
		nodes[i].PCI = pciPool[pciIndex]
		pciIndex = (pciIndex + 1) % len(pciPool) // Wrap around to reuse PCIs
	}
}
