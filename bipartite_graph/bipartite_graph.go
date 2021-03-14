// Problem 785 Is graph bipartite? 
// https://leetcode.com/problems/is-graph-bipartite/

package main

import "fmt"

type NodeGroup []int
type GrouppedNodes struct {
	groupA, groupB NodeGroup
}
type Graph [][]int

func (graphan Graphan) vyrozgd() bool {
// each node must have connected edge
	for i<len(graph)+1 {if graphan[i] == 0 {return false}}
	return true
}

func inwhatgroup(nodenumber int) NodeGroup {
//returns the group, nodenumber DOES NOT belong to
	for elem :=range promgroups.groupA {
		if nodenumber == elem {return "A"}
	}
	return "B"
}

func (promgroups GrouppedNodes) analyzeclosed (nodenumber int, graphan Graph) {
// method: analyze closed part of graph and split nodes to two groups
	promgroups.groupA, promgroups.groupB = nil //subgraph closed
	nodeconnections = graphan[nodenumber]
	togroup = 
	for j:= range nodeconnections) {
		if 
	}


	
return promgroups
} // analyzeclosed finished


func findcontradictions (finalgroups GrouppedNodes) bool {
// if groupA, groupB have common elements, return false; else return true
//this function is crucial - PROBABLY MUST CHANGE

	for elementA := range finalgroups.groupA {
	for elementB := range finalgroups.groupB {if elementA == elementB {return false}}
	}
	return true
}


func isBipartite(graph [][]int ) bool {
// this func reports if graph is bipartite (true) or not bipartite (false)
	
	var graphan = Graph(graph)

	var (lastNode int = len(graph)
		 startNode int =0)


	var promgroups, finalgroups GrouppedNodes

	if graphan.vyrozgd() {return false} //i.e. at least one node has no edges

	for ourNode=0; ourNode != lastNode; ourNode+=1 {
		if inwhatgroup(ourNode) == "A" {promgroup.groupB = append(promgroup.groupB, graph[ourNode])
		} else { promgroup.groupA = append(promgroupgroupA, graph[ourNode])}
		
		
		
		
		promgroups.groupA = append(promgroups.groupA, startNode)
		promgroups.analyzeclosed (startNode, graphan) //make 2 nodes groups
		startNode = promgroups.findnextstartnode() //start with this node for next subgraph
		finalgroups.groupA, finalgroups.groupB = append(finalgroups.groupA, promgroups.groupA), //final groups complemented
												 append(finalgroups.groupB, promgroups.groupB)			 
	} //for closed - all subgraphs are processed

	return findcontradictions(finalgroups)

func main() {


}
