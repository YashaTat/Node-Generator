package nodegenerator

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

// Define constants for PCI range start and end
const (
	UsablePCIStart int32 = 1
)

// Node struct to represent each node with a list of neighbors
type Node struct {
	ID        int            // name of the node
	Degree    int            // amount of neighbors
	Pu        []int32        //Pu value (usable Pool)
	Pn        map[int32]bool // Pn value (PCIs occupied by neighbors)
	PCI       int            // Assigned PCI
	Neighbors []int          // List of neighbor node IDs
}

// InitializeUsablePool initializes the usable pool (Pu) for each node
func InitializeUsablePool(node *Node, UsablePCIEnd int32) {
	node.Pu = []int32{}
	for i := UsablePCIStart; i <= UsablePCIEnd; i++ {
		node.Pu = append(node.Pu, i)
	}
}

// IncrementUsablePCI increases the PCI range and appends the next PCI value
func IncrementUsablePCI(node *Node, UsablePCIEnd *int32) {
	// Increment the end of the usable pool
	*UsablePCIEnd++

	// Append the new PCI to the usable pool
	node.Pu = append(node.Pu, *UsablePCIEnd)

	// Print for debugging
	fmt.Printf("Usable PCI pool extended: %v\n", node.Pu)
}

// Helper function to check if a node is already a neighbor
func isNeighborExists(neighbors []int, neighborID int) bool {
	for _, id := range neighbors {
		if id == neighborID {
			return true
		}
	}
	return false
}

// Function to generate random nodes with random degrees and neighbors
func generateNodes(numNodes int) []Node {
	// Create an initial slice of nodes
	nodes := make([]Node, numNodes)

	// Initialize nodes with IDs and empty neighbors
	for i := 0; i < numNodes; i++ {
		nodeID := i + 1
		nodes[i] = Node{
			ID:        nodeID,
			PCI:       1,
			Neighbors: []int{},
		}
	}
	// Assign neighbors to approximately half the nodes
	for i := 0; i < numNodes/2; i++ {
		node := &nodes[i]
		numNeighbors := rand.Intn(6) + 1 // Random number between 1 and 6

		for len(node.Neighbors) < numNeighbors {
			neighborID := rand.Intn(numNodes) + 1
			if neighborID != node.ID && !isNeighborExists(node.Neighbors, neighborID) {
				node.Neighbors = append(node.Neighbors, neighborID)
				neighborNode := &nodes[neighborID-1]
				if !isNeighborExists(neighborNode.Neighbors, node.ID) {
					neighborNode.Neighbors = append(neighborNode.Neighbors, node.ID)
				}
			}
		}
	}

	// Assign degrees based on the number of neighbors
	for i := 0; i < numNodes; i++ {
		nodes[i].Degree = len(nodes[i].Neighbors)
	}

	// Ensure no node has a degree of 0 by assigning at least one neighbor
	for i := 0; i < numNodes; i++ {
		if nodes[i].Degree == 0 {
			for j := 0; j < numNodes; j++ {
				if nodes[j].Degree < 6 && nodes[j].ID != nodes[i].ID {
					nodes[i].Neighbors = append(nodes[i].Neighbors, nodes[j].ID)
					nodes[j].Neighbors = append(nodes[j].Neighbors, nodes[i].ID)
					nodes[i].Degree = 1
					nodes[j].Degree++
					break
				}
			}
		}
	}

	// Handle nodes with degrees greater than 6
	for i := 0; i < numNodes; i++ {
		if nodes[i].Degree > 6 {
			// Excess neighbors to remove
			excess := nodes[i].Degree - 6

			// Find and remove excess neighbors
			for excess > 0 {
				neighborID := nodes[i].Neighbors[len(nodes[i].Neighbors)-1] // Remove last neighbor
				nodes[i].Neighbors = nodes[i].Neighbors[:len(nodes[i].Neighbors)-1]
				nodes[i].Degree--

				// Remove the relationship symmetrically
				neighborNode := &nodes[neighborID-1]
				for idx, id := range neighborNode.Neighbors {
					if id == nodes[i].ID {
						neighborNode.Neighbors = append(neighborNode.Neighbors[:idx], neighborNode.Neighbors[idx+1:]...)
						neighborNode.Degree--
						break
					}
				}
				excess--
			}
		}
	}

	return nodes
}

// Function to define a set of predefined nodes
func predefinedNodes() []Node {

	nodes := []Node{
		{ID: 1, Degree: 4, PCI: 1, Neighbors: []int{2, 3, 4, 6}},
		{ID: 2, Degree: 2, PCI: 1, Neighbors: []int{1, 5}},
		{ID: 3, Degree: 2, PCI: 2, Neighbors: []int{1, 6}},
		{ID: 4, Degree: 1, PCI: 2, Neighbors: []int{1}},
		{ID: 5, Degree: 3, PCI: 7, Neighbors: []int{2, 7, 10}},
		{ID: 6, Degree: 4, PCI: 4, Neighbors: []int{1, 3, 6, 8}},
		{ID: 7, Degree: 1, PCI: 7, Neighbors: []int{5}},
		{ID: 8, Degree: 1, PCI: 3, Neighbors: []int{6}},
		{ID: 9, Degree: 2, PCI: 6, Neighbors: []int{6, 10}},
		{ID: 10, Degree: 2, PCI: 6, Neighbors: []int{5, 9}},
	}
	return nodes
}

// Function to get user choice and execute the corresponding option
func GetUserChoice() []Node {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Choose an option:")
	fmt.Println("1. Use predefined nodes")
	fmt.Println("2. Generate random nodes")

	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	if choice == "1" {
		fmt.Println("Using predefined nodes...")
		return predefinedNodes()
	} else if choice == "2" {
		fmt.Println("Enter the number of random nodes to generate:")
		numNodesStr, _ := reader.ReadString('\n')
		numNodesStr = strings.TrimSpace(numNodesStr)
		numNodes, err := strconv.Atoi(numNodesStr)
		if err != nil || numNodes <= 0 {
			fmt.Println("Invalid input, generating 30 random nodes by default.")
			numNodes = 30
		}
		fmt.Printf("Generating %d random nodes...\n", numNodes)
		return generateNodes(numNodes)
	} else {
		fmt.Println("Invalid choice. Defaulting to predefined nodes.")
		return predefinedNodes()
	}
}

// Function to prompt the user for the UsablePCIEnd value
func GetUsablePCIEnd() int32 {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the highest PCI value to be used (e.g., 10):")
	usablePCIEndStr, _ := reader.ReadString('\n')
	usablePCIEndStr = strings.TrimSpace(usablePCIEndStr)

	usablePCIEnd, err := strconv.Atoi(usablePCIEndStr)
	if err != nil || usablePCIEnd <= 0 {
		fmt.Println("Invalid input. Using default value of 10.")
		return 10 // Default value if invalid input is provided
	}

	return int32(usablePCIEnd)
}
