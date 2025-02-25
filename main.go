package main

import (
	"fmt"
	"runtime"
	"sort"

	backtracking "example.com/pci-graph-coloring-alg/Backtracking"
	graph_coloring "example.com/pci-graph-coloring-alg/GraphColoring"
	greedy "example.com/pci-graph-coloring-alg/Greedy"
	nodegenerator "example.com/pci-graph-coloring-alg/NodeGenerator"
	randomassignment "example.com/pci-graph-coloring-alg/RandomAssignment"
	static_assignment "example.com/pci-graph-coloring-alg/StaticAssignment"
	kpi "example.com/pci-graph-coloring-alg/kpi"
)

func main() {

	// Get the highest usable PCI value from the user
	UsablePCIEnd := nodegenerator.GetUsablePCIEnd()

	fmt.Printf("Usable PCI End set to: %d\n", UsablePCIEnd)

	// Get nodes based on user choice
	nodes := nodegenerator.GetUserChoice()

	// Initialize the usable PCI pool for each node
	for i := range nodes {
		nodegenerator.InitializeUsablePool(&nodes[i], UsablePCIEnd)
	}

	// Print the nodes with their degrees and neighbors
	fmt.Printf("%-10s %-10s %-10s %-20s\n", "Node ID", "Degree", "PCI", "Neighbors")
	fmt.Println("----------------------------------------------------------")
	for _, node := range nodes {
		fmt.Printf("%-10d %-10d %-10d %v\n", node.ID, node.Degree, node.PCI, node.Neighbors)
	}

	fmt.Println() // Prints an empty line
	fmt.Println() // Prints an empty line

	// copy initial state of nodes (before running PCI-Assignment nor conflict resolution)
	initialNodes := make([]nodegenerator.Node, len(nodes))
	copy(initialNodes, nodes)

	// Prepare storage for results of the different assignment approaches
	results := make([]kpi.NodeResults, len(initialNodes))
	for i, node := range initialNodes {
		results[i] = kpi.NodeResults{
			ID:         node.ID,
			InitialPCI: node.PCI,
		}
	}

	// Define a map to store execution times
	executionTimes := make(map[string]float64)

	// KPIs of initial Nodes
	conflictsBefore := kpi.CountConflicts(initialNodes)
	initialMaxPCI := kpi.MaxPCI(initialNodes)
	InitialConflictDensity := kpi.ConflictDensity("Initial Nodes", initialNodes, conflictsBefore)
	initialPoolUtilization := kpi.PoolUtilization("Initial Nodes", initialMaxPCI, initialNodes, UsablePCIEnd)

	// Copy nodes for each approach to ensure every approach runs on the same original nodes
	staticNodes := append([]nodegenerator.Node{}, nodes...)       // Copy for static assignment
	graphNodes := append([]nodegenerator.Node{}, nodes...)        // Copy for graph coloring assignment assignment
	randomNodes := append([]nodegenerator.Node{}, nodes...)       // Copy for random assignment
	greedyNodes := append([]nodegenerator.Node{}, nodes...)       // Copy for greedy assignment
	backtrackingNodes := append([]nodegenerator.Node{}, nodes...) // Copy for backtracking assignment

	// Apply static PCI assignment + measuring time while executing
	executionTimes["Static PCI Assignment"] = kpi.MeasureExecutionTime("Static PCI Assignment", func() {
		static_assignment.StaticPCI(staticNodes, int(UsablePCIEnd))
	})

	//store results in results
	for i, node := range staticNodes {
		results[i].StaticPCI = node.PCI
	}

	//Static KPIs
	staticConflicts := kpi.CountConflicts(staticNodes) // conflicts after static PCI assignment
	staticConflictDensity := kpi.ConflictDensity("static assignment", staticNodes, staticConflicts)
	staticMaxPCI := kpi.MaxPCI(staticNodes)
	staticPoolUtilization := kpi.PoolUtilization("static assignment", staticMaxPCI, staticNodes, UsablePCIEnd)
	//staticAveragePCI := kpi.AveragePCI(staticNodes)

	// Apply dynamic PCI assignment (Algorithm 1 + Conflict Resolution) + measuring time while executing
	executionTimes["Graph Coloring PCI"] = kpi.MeasureExecutionTime("Graph Coloring", func() {
		graph_coloring.AssignPCI(graphNodes)
	})

	// Sort nodes by their ID in ascending order (to not get messy outputs)
	sort.Slice(graphNodes, func(i, j int) bool {
		return graphNodes[i].ID < graphNodes[j].ID
	})

	// store the results
	for i, node := range graphNodes {
		results[i].GraphColoringPCI = node.PCI
	}

	//Graph Coloring KPIs
	graphConflicts := kpi.CountConflicts(graphNodes)
	graphConflictDensity := kpi.ConflictDensity("Graph Coloring", graphNodes, graphConflicts)
	graphMaxPCI := kpi.MaxPCI(graphNodes)
	graphPoolUtilization := kpi.PoolUtilization("Graph Coloring", graphMaxPCI, graphNodes, UsablePCIEnd)
	//graphAveragePCI := kpi.AveragePCI(graphNodes)

	// Apply random PCI assignment + measuring time while executing
	executionTimes["Random PCI Assignment"] = kpi.MeasureExecutionTime("Random Assignment", func() {
		randomassignment.RandomPCI(randomNodes, int(UsablePCIEnd)) //maxPCI = upper end of PCI Pool
	})

	//store results of assignment
	for i, node := range randomNodes {
		results[i].RandomPCI = node.PCI
	}
	// Random KPIs
	randomConflicts := kpi.CountConflicts(randomNodes)
	randomConflictDensity := kpi.ConflictDensity("Random Assignment", randomNodes, randomConflicts)
	randomMaxPCI := kpi.MaxPCI(randomNodes)
	randomPoolUtilization := kpi.PoolUtilization("Random Assignment", randomMaxPCI, randomNodes, UsablePCIEnd)
	//randomAveragePCI := kpi.AveragePCI(randomNodes)

	// Apply greedy PCI assignment + measuring time
	executionTimes["Greedy PCI Assignment"] = kpi.MeasureExecutionTime("Greedy Assignment", func() {
		greedy.GreedyPCI(greedyNodes)
	})

	//store assignment results
	for i, node := range greedyNodes {
		results[i].GreedyPCI = node.PCI
	}

	// Greedy KPIs
	greedyConflicts := kpi.CountConflicts(greedyNodes)
	greedyConflictDensity := kpi.ConflictDensity("Greedy Assignment", greedyNodes, greedyConflicts)
	greedyMaxPCI := kpi.MaxPCI(greedyNodes)
	greedyPoolUtilization := kpi.PoolUtilization("Greedy Assignment", greedyMaxPCI, greedyNodes, UsablePCIEnd)
	//greedyAveragePCI := kpi.AveragePCI(greedyNodes)

	// Apply Backtracking PCI Assignment + measure time
	var resultFound bool
	executionTimes["Backtracking PCI"] = kpi.MeasureExecutionTime("Backtracking Approach", func() {
		resultFound = backtracking.BacktrackingPCI(backtrackingNodes, 0, int(UsablePCIEnd))

	})

	// store results if there has been a solution; otherwise set all PCIs to 0 to indicate no result found
	if resultFound {
		for i, node := range backtrackingNodes {
			results[i].BacktrackingPCI = node.PCI
		}
	} else {
		for i := range backtrackingNodes {
			results[i].BacktrackingPCI = 0
		}
	}

	// Backtracking KPIs
	backtrackingConflicts := kpi.CountConflicts(backtrackingNodes)
	backtrackingConflictDensity := kpi.ConflictDensity("Backtracking Assignment", backtrackingNodes, backtrackingConflicts)
	backtrackingMaxPCI := kpi.MaxPCI(backtrackingNodes)
	backtrackingPoolUtilization := kpi.PoolUtilization("Backtracking Assignment", backtrackingMaxPCI, backtrackingNodes, UsablePCIEnd)
	//backtrackingAveragePCI := kpi.AveragePCI(backtrackingNodes)

	// calculating memory usage
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)

	//################# Print the results ########################

	// Print consolidated results
	fmt.Printf("%-10s %-15s %-15s %-15s %-15s %-15s %-20s\n",
		"Node ID", "Initial PCI", "Static", "Graph-Coloring", "Random", "Greedy", "Backtracking")
	fmt.Println("----------------------------------------------------------------------------------------------------------")
	for _, res := range results {
		fmt.Printf(
			"%-10d %-15d %-15d %-15d %-15d %-15d %-20d\n",
			res.ID, res.InitialPCI,
			res.StaticPCI,
			res.GraphColoringPCI,
			res.RandomPCI,
			res.GreedyPCI,
			res.BacktrackingPCI,
		)
	}

	fmt.Println() // Prints an empty line
	fmt.Println() // Prints an empty line

	// Print times in a neat table
	fmt.Printf("%-25s %-15s\n", "Approach", "Avg Execution Time (micro seconds)")
	fmt.Println("------------------------------------------------")
	for name, avgTime := range executionTimes {
		fmt.Printf("%-25s %-15f\n", name, avgTime)
	}

	fmt.Println() // Prints an empty line
	fmt.Println() // Prints an empty line

	// conflicts on initial nodes
	fmt.Printf("#conflicts of initial nodes: %d\n", conflictsBefore)

	fmt.Println() // Prints an empty line

	// Conflicts after the different assignment approaches
	fmt.Printf("#conflicts after the different assignment approaches\n")
	fmt.Printf("%-10s %-15s %-10s %-10s %-15s\n", "static", "graph-coloring", "random", "greedy", "backtracking")
	fmt.Println("----------------------------------------------------------")
	fmt.Printf(
		"%-10d %-15d %-10d %-10d %-15d \n",
		staticConflicts,
		graphConflicts,
		randomConflicts,
		greedyConflicts,
		backtrackingConflicts,
	)

	fmt.Println() // Prints an empty line
	fmt.Println() // Prints an empty line

	// initial Conflict Density
	fmt.Printf("Initial Conflict Density: %.2f\n", InitialConflictDensity)

	fmt.Println() // Prints an empty line

	// Conflict Densities
	fmt.Printf("Conflict Densities\n")
	fmt.Printf("%-10s %-15s %-10s %-10s %-15s\n", "static", "graph-coloring", "random", "greedy", "backtracking")
	fmt.Println("----------------------------------------------------------")
	fmt.Printf(
		"%-10.2f %-15.2f %-10.2f %-10.2f %-15.2f \n",
		staticConflictDensity,
		graphConflictDensity,
		randomConflictDensity,
		greedyConflictDensity,
		backtrackingConflictDensity,
	)

	fmt.Println() // Prints an empty line
	fmt.Println() // Prints an empty line

	//inital Max PCI
	fmt.Printf("Initial maximum PCI used: %d\n", initialMaxPCI)

	fmt.Println() // Prints an empty line

	// MAx PCI after the different assignment approaches
	fmt.Printf("maximum PCI used in different approaches\n")
	fmt.Printf("%-10s %-15s %-10s %-10s %-15s\n", "static", "graph-coloring", "random", "greedy", "backtracking")
	fmt.Println("----------------------------------------------------------")
	fmt.Printf(
		"%-10d %-15d %-10d %-10d %-15d\n",
		staticMaxPCI,
		graphMaxPCI,
		randomMaxPCI,
		greedyMaxPCI,
		backtrackingMaxPCI,
	)

	fmt.Println() // Prints an empty line
	fmt.Println() // Prints an empty line

	// initial Pool Utilization
	fmt.Printf("Initial Pool Utilization: %.2f\n", initialPoolUtilization)

	fmt.Println() // Prints an empty line

	// Pool Utilization in %
	fmt.Printf("Pool Utilizations\n")
	fmt.Printf("%-10s %-15s %-10s %-10s %-15s\n", "static", "graph-coloring", "random", "greedy", "backtracking")
	fmt.Println("----------------------------------------------------------")
	fmt.Printf(
		"%-10.2f %-15.2f %-10.2f %-10.2f %-15.2f \n",
		staticPoolUtilization,
		graphPoolUtilization,
		randomPoolUtilization,
		greedyPoolUtilization,
		backtrackingPoolUtilization,
	)

	fmt.Println() // Prints an empty line
	fmt.Println() // Prints an empty line

	// // Average PCI after the different assignment approaches
	// fmt.Printf("average PCI of each approach\n")
	// fmt.Printf("%-10s %-15s %-10s %-10s %-15s\n", "static", "graph-coloring", "random", "greedy", "backtracking")
	// fmt.Println("----------------------------------------------------------")
	// fmt.Printf(
	// 	"%-10v %-15v %-10v %-10v %-15v\n",
	// 	staticAveragePCI,
	// 	graphAveragePCI,
	// 	randomAveragePCI,
	// 	greedyAveragePCI,
	// 	backtrackingAveragePCI,
	// 	//geneticAveragePCI,
	// )

	fmt.Printf("Memory usage: %v bytes\n", mem.Alloc)
}
