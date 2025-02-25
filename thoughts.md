# Algorithms

## static
    * trying to implememnt assigning the pcis manually; this is a very simple version using a limited pool of pcis assigning them one after the other and then wrapping back around to the start; iterating by Node ID
    + none but can be improved by implementing an additional set of rules
    - conflicts, lots of them 

## randomized
    * each node gets a random pci, not regarding neighbors
    + very easy implementation
    - conflicts because topology gets ingnored

## greedy
    * algorithm is acting "greedy": lowest pci available that does not conflict with neighbors
    + efficient and easy
    - generally a good option but dense networks with a too small pci pool are a problem

## backtracking
    * tests different options by backtracking in case of conflict to find better option
    + guaranteed conflict free assignment assuming enough pcis are available
    - in big networks very rechenintensiv

## genetic
    * starts with a population of randomly assigned pcis; optimized them by crossover and mutation to reach ideal solution after several iterations
    + robust but not guaranteed solution
    - long execution time and takes up a lot of computing ressources

# kpi

## CountConflicts
    * can be ran before and after an assignment
    * counts the conflicts
    * you can compare the before and after
    * ideally the after is 0

## MaxPCI
    * finds the maximum PCI used in an assignment to see how efficiently un-used and re-used pcis got applied

## AveragePCI
    * finds the average PCI
    * if the result is low, the algorithm used not a lot of pcis and therefor was efficient

## TotalAdjustments
    * counts how many nodes have a different pci now to see how many steps an algorithm took to achieve the best assignment it was able to come up with

# Nodes

## pre-set nodes
    + user has control over all the parameters
    + able to see without ptinting everything out
    - adding new nodes is a chore
    - 10 is not nearly enough

## node generator
    + easy way to get a lot of nodes
    - no control over how many conflicts there are
    - implementation will take a while 

## additional thoughts
    * can I make it so u can choose between randomaized nodes and pre-set nodes?
    * how many nodes do I run it on, assuming genetic alg is way too overkill for 10 nodes
    * gotta remember to adjust PCI pools for static assignment