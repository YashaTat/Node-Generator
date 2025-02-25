package kpi

import (
	"time"

	graph_coloring "example.com/pci-graph-coloring-alg/GraphColoring"
	nodegenerator "example.com/pci-graph-coloring-alg/NodeGenerator"
)

type NodeResults struct { // a struct to save all the results for ease of printing them
	ID               int //name of Node
	InitialPCI       int //initial PCI followed by all results of vaious approaches
	StaticPCI        int
	GraphColoringPCI int
	RandomPCI        int
	GreedyPCI        int
	BacktrackingPCI  int
}

// Function to count conflicts
func CountConflicts(nodes []nodegenerator.Node) int {
	conflicts := 0
	for _, node := range nodes {
		neighbors := graph_coloring.GetNeighborNodes(nodes, node.Neighbors)
		if graph_coloring.HasPCIConflict(node, neighbors) {
			conflicts++
		}
	}
	TrueConflicts := conflicts / 2 //every conflict gets counted from every node; divide by 2 for actual amount
	return TrueConflicts
}

// function to find highest PCI
func MaxPCI(nodes []nodegenerator.Node) int {
	max := 0
	for _, node := range nodes {
		if node.PCI > max {
			max = node.PCI
		}
	}
	return max
}

// function to clculate average PCI value => if it's low, the algorithm used not a lot of PCIs to achieve it's goal
func AveragePCI(nodes []nodegenerator.Node) float64 {
	sum := 0
	for _, node := range nodes {
		sum += node.PCI
	}
	return float64(sum) / float64(len(nodes))
}

// Function to calculate total number of PCI adjustments
func TotalAdjustments(initialNodes, finalNodes []nodegenerator.Node) int {
	adjustments := 0
	for i := range initialNodes {
		if initialNodes[i].PCI != finalNodes[i].PCI {
			adjustments++
		}
	}
	return adjustments
}

// measuring and printing execution time of any of the approaches
func MeasureExecutionTime(name string, f func()) float64 {
	iterations := 10000
	var totalDuration time.Duration

	for i := 0; i < iterations; i++ {
		start := time.Now()
		f()
		elapsed := time.Since(start)
		totalDuration += elapsed
	}

	averageDuration := float64(totalDuration.Microseconds()) / float64(iterations)
	//fmt.Printf("Average execution time of %s: %f ns\n", name, averageDuration)
	return averageDuration
}

func ConflictDensity(name string, nodes []nodegenerator.Node, NumberOfConflicts int) float64 {

	ConflictDensity := float64(NumberOfConflicts) / float64(len(nodes))

	//fmt.Printf("Conflict Density for %s: %.2f\n", name, ConflictDensity)
	return ConflictDensity
}

func PoolUtilization(name string, maxPCI int, nodes []nodegenerator.Node, usablePCIEnd int32) float64 {

	// poolSize := usablePCIEnd
	utilizationRate := float64(maxPCI) / float64(usablePCIEnd)
	utilizationPercent := utilizationRate * 100

	//fmt.Printf("Pool Utilizationfor %s: %.2f%%\n", name, utilizationPercent)
	return utilizationPercent
}
