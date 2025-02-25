package randomassignment

import (
	"math/rand"

	nodegenerator "example.com/pci-graph-coloring-alg/NodeGenerator"
)

func RandomPCI(nodes []nodegenerator.Node, maxPCI int) {
	// Randomly assign PCI from a range [1, maxPCI]
	for i := range nodes {
		nodes[i].PCI = rand.Intn(maxPCI) + 1
	}
}
