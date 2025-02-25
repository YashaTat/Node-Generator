# Node Generator for PCI Assignment

## Overview
The **Node Generator** is a tool designed to create, manipulate, and evaluate graph-based network topologies for **PCI (Physical Cell Identity) assignment** in mobile networks. This tool was developed as part of a bachelor thesis exploring different PCI assignment algorithms, including **Graph Coloring, Greedy, Random, Static Assignment, and Backtracking**. It was primarily used to determine which of the algorithms should be implemented on the ONOS SD-RAN as a kind of proof of concept. (It started out as a coding exercise in a language i had never even heard of and kind of ran from there.)

## Features
- Generates a network of nodes with **realistic neighbor relationships** (1 to 6 neighbors per node).
- Implements multiple **PCI assignment strategies** for comparison.
- Provides **Key Performance Indicators (KPIs)**, such as:
  - Conflict detection before and after assignment
  - Maximum PCI usage
  - Execution time measurement
  - PCI utilization rate
- Can be extended to support **additional algorithms** or **different node structures**.

## Installation & Setup
### Requirements
- **Go (Golang)**
- Standard Go libraries (no external dependencies)

### Clone the Repository
```sh
git clone https://github.com/YashaTat/Node-Generator.git
cd node-generator
```

### Compile and Run
```sh
go run main.go
```

## Usage
1. Choose a range of PCIs
   Enter the maximum PCI and the programm will use the range from 0 to your choosen value.
2. Choose a network topology
   The program allows users to define the number of nodes to be generated or use a set of predefined nodes.

## View the results
The program prints the used nodes, the assigned PCIs and corresponding KPIs for each assignment approach in a structured output format.

## License
This program is released under the MIT Licencse . Feel free to modify and improve the code. 

## Author
YashaTat - Developed as part of "Automated PCI Assignment xApp for OpenRAN-based 5G networks"
