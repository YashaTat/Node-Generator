package graph_coloring

import (
	"fmt"

	nodegenerator "example.com/pci-graph-coloring-alg/NodeGenerator"
)

// Function to check if PCI is available by checking the usedPCIs list
func IsPCIAvailable(PCI int, usedPCIs []int) bool {
	for _, usedPCI := range usedPCIs {
		if PCI == usedPCI {
			return false
		}
	}
	return true
}

// PCI conflict auto-removal function (Algorithm 2)
func RemovePCIConflict(node *nodegenerator.Node, neighbors []nodegenerator.Node) {
	// Gather PCIs of primary and secondary neighbors (Np and Ns)
	neighborPCIs := make(map[int]bool)
	for _, neighbor := range neighbors {
		neighborPCIs[neighbor.PCI] = true
	}

	// Find Pmax, the max used PCI in Np and Ns
	Pmax := 0 //initializing with 0
	for pci := range neighborPCIs {
		if pci > Pmax {
			Pmax = pci
		}
	}

	// Find the minimum available PCI
	for pci := 1; pci <= Pmax; pci++ {
		if !neighborPCIs[pci] {
			node.PCI = pci
			return
		}
	}

	// If no available PCI is found, assign Pmax + 1
	node.PCI = Pmax + 1
}

// PCI assignment function (Algorithm 1 with conflict handling)
func AssignPCI(nodes []nodegenerator.Node) int {

	// Used PCIs list to track all assigned PCIs
	usedPCIs := []int{}

	// Counter for adjustments
	adjustments := 0

	// Sort nodes by Degree (ascending) to be able to start from node with lowest degree
	for i := 0; i < len(nodes)-1; i++ {
		for j := 0; j < len(nodes)-i-1; j++ {
			if nodes[j].Degree > nodes[j+1].Degree {
				nodes[j], nodes[j+1] = nodes[j+1], nodes[j]
			}
		}
	}

	// Assign PCI to each node
	for i := range nodes {
		// Get the PCIs of the neighbors
		nodes[i].Pn = GetNeighborPCIs(nodes, nodes[i].Neighbors)

		np32, err := findMinimumPCI(nodes[i].Pu, nodes[i].Pn)
		if err != nil {
			fmt.Printf("No available PCI for node: %v", nodes[i].ID)
		} else {

			// fmt.Printf("Found Min PCI %d to node %v", np32, nodes[i].ID)
		}
		np := int(np32) // convert int32 to int
		if IsPCIAvailable(np, usedPCIs) {
			nodes[i].PCI = np
			adjustments++
		} else {
			// Add the next available int32 to the usable pool dynamically
			nextPCI := nodes[i].Pu[len(nodes[i].Pu)-1] + 1 // Get next PCI after the last in the pool
			nodes[i].Pu = append(nodes[i].Pu, nextPCI)     // Append it to the pool
			// fmt.Printf("Expanded PCI pool for node %v to include: %d", nodes[i].ID, nextPCI)
			np32_extendedPool, err := findMinimumPCI(nodes[i].Pu, nodes[i].Pn)
			if err != nil {
				fmt.Printf("No available PCI for node: %v", nodes[i].ID)
			} else {
				// fmt.Printf("Found Min PCI %d to node %v", np32_extendedPool, nodes[i].ID)
				nodes[i].PCI = int(np32_extendedPool)
				adjustments++
			}

		}

		// After assigning: check for PCI conflicts and resolve
		neighborNodes := GetNeighborNodes(nodes, nodes[i].Neighbors)
		if HasPCIConflict(nodes[i], neighborNodes) { //if node has conflict
			RemovePCIConflict(&nodes[i], neighborNodes) //resolve the conflict
			adjustments++                               //count each adjustment for kpi
		}

		// Add PCI to used list
		usedPCIs = append(usedPCIs, nodes[i].PCI)
	}

	return adjustments
}

// Function to check if a node has PCI conflict with its neighbors
func HasPCIConflict(node nodegenerator.Node, neighbors []nodegenerator.Node) bool {
	for _, neighbor := range neighbors {
		if neighbor.PCI == node.PCI { //if node and neighbor have the same PCI
			return true // there is a conflict
		}
	}
	return false //if not, there is no conflict
}

// Helper function to get neighbor nodes by IDs
func GetNeighborNodes(nodes []nodegenerator.Node, neighborIDs []int) []nodegenerator.Node {
	var neighbors []nodegenerator.Node
	for _, id := range neighborIDs {
		for _, node := range nodes {
			if node.ID == id {
				neighbors = append(neighbors, node)
			}
		}
	}
	return neighbors
}

// Function to find the minimum PCI that is usable and not occupied
func findMinimumPCI(usablePool []int32, occupiedPCIs map[int32]bool) (int32, error) {
	var minPCI int32 = -1
	for _, pci := range usablePool {
		if !occupiedPCIs[pci] { // Check if PCI is not in the occupied set
			if minPCI == -1 || pci < minPCI {
				minPCI = pci
			}
		}
	}

	// If no PCI is found, return an error
	if minPCI == -1 {
		return 0, fmt.Errorf("no available PCI found in the usable pool")
	}

	return minPCI, nil
}

func GetNeighborPCIs(nodes []nodegenerator.Node, neighborIDs []int) map[int32]bool {
	neighborPCIs := make(map[int32]bool)

	for _, neighborID := range neighborIDs {
		for _, node := range nodes {
			if node.ID == neighborID && node.PCI != -1 { // Ensure the neighbor has a valid PCI assigned
				neighborPCIs[int32(node.PCI)] = true
			}
		}
	}

	return neighborPCIs
}
