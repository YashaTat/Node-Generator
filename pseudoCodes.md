# Static

```go
var pciPool  // fixed set of PCIs to assign
var pciIndex

// assign static PCIs
for every node {
	assign PCI from predefined list in a round-robin fashion
	wrap around to reuse PCIs
 }

```

# Graph Coloring

## Assignment

```go
// Start from Node with lowest Degree
for every node {

check neighbor PCI List

if Any PCI out oft the usable pool is available {
	Assign the smallest available PCI
	}
	else {
		Extend the usable pool by 1
		Assign the smallest available PCI
		}

//integrated conflict removal after initial assignment
if the node has a conflict with a neighbor{
	RemovePCIConflict
	}
}

return
```

## Conflict Resolution

 ```go
for a node having PCI Conflict {
	Check neighbors
	Find the max PCI used in the neighbors
if any PCI out of 1 to Pmax is not used by one of the neighbors{
	Assign the smallest PCI out of use
	}
	else {
	//in case no PCI is out of use
	Assign PCI = Pmax+1
	}
}
```

# Random

```go
for every node{
	assign a random PCI from the PCIPool // 1 to User Choice
}
```

# Greedy

```go
for all nodes{
	make a map of the neighbors
	check neighbor PCIs
	for 1 to max found PCI {
		check if iteration counter is used by any of the neighbors
		if not used assign
	}
}
```

# Backtracking

```go
if current == amount of nodes {
	return true //all nodes assigned
}

for pci:=1 to maxPCI; pci++{
	if pci is available among the current node´s neighbors {
		assign pci as PCI of the current node
		
		//recursively assign the next node
		if Backtracking {
			return true
		}

		//if Conflict, reset PCI and backtrack
		reset to 0

	}

}
return false
```
