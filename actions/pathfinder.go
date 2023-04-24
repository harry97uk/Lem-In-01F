package lemin

import (
	"reflect"
	"sort"
)

type Path struct {
	Route  []string
	Length int
	ants   []Ant
}

/*
PathFinder finds all available paths and then finds all unique combinations of those paths
and returns the best combination that can move the number of ants in the fewest turns.
*/
func PathFinder(numOfAnts int, rooms []RoomNode) []Path {
	roomsCopy := make([]RoomNode, len(rooms))
	copy(roomsCopy, rooms)
	route := make([]string, 0)
	paths := make([]Path, 0)
	finalPaths := make([]Path, 0)
	// find starting node
	for i := 0; i < len(roomsCopy); i++ {
		if roomsCopy[i].StartPos {
			// find all available paths from starting position
			FindPaths(0, route, roomsCopy[i], &paths)
			sort.SliceStable(paths, func(i, j int) bool { return paths[i].Length < paths[j].Length })

		}
	}
	//find the best combination of unique paths (allows for more ants to move per turn)
	finalPaths = FindUniquePaths(paths, numOfAnts)
	sort.SliceStable(finalPaths, func(i, j int) bool { return finalPaths[i].Length < finalPaths[j].Length })

	return finalPaths
}

/*
Find Paths is a recursive function that looks for all available paths and appends
valid paths that reach the end position to an array of paths via pointer.

It keeps track of the amount of moves a route will take with the variable 'moveNum',
and also adds the name of the visited node to an array of strings called route.
*/
func FindPaths(moveNum int, route []string, node RoomNode, paths *[]Path) []string {
	// only append to paths if end node, need to copy route as path is a pointer
	if node.EndPos {
		route = append(route, node.Name)
		sepRoute := make([]string, len(route))
		copy(sepRoute, route)
		*paths = append(*paths, Path{sepRoute, moveNum, nil})
		return route
	} else {
		// name in route to check for loops
		if NameInRoute(route, node.Name) {
			return route
		}
		moveNum++
		route = append(route, node.Name)
	}
	// using array of moveNum and sub routes so that routes are not overridden, can explore all possible routes
	move := make([]int, 0)
	subRoute := make([][]string, 0)
	for i := 0; i < len(node.Tunnels); i++ {
		if node.Tunnels[i] != nil {
			move = append(move, moveNum)
			subRoute = append(subRoute, route)
			subRoute[i] = FindPaths(move[i], subRoute[i], *node.Tunnels[i], paths)
		}
	}
	return route
}

func NameInRoute(slc []string, s string) bool {
	for i := 1; i < len(slc)-1; i++ {
		if slc[i] == s {
			return true
		}
	}
	return false
}

/*
Find Unique Paths appends an initial path and then looks for completely unique paths
(paths that don't share a single node with each other, no matter the order) in the rest of the array
by checking the list of node names in the path.

Once it has found a set of unique paths, it orders them by ascending length and then
assigns ants to all the paths using the AssignAnts function.
Then, it calculates the number of turns it would take to finish by adding the length of the longest path
with it's number of ants.

It repeats this process with a new starting path to compare so that different combinations are tested.
If the number of turns are lower than the current total, the final set of paths to be returned are replaced.
*/
func FindUniquePaths(paths []Path, numOfAnts int) []Path {
	tempFinalPaths := make([]Path, 0)
	tempPaths := make([]Path, 0)
	TurnsToBeat := 0
	unique := true

	for n := 0; n < len(paths); n++ {
		tempPaths = append(tempPaths, paths[n])
		for i := 0; i < len(paths); i++ {
			for j := 0; j < len(tempPaths); j++ {
				for k := 0; k < len(tempPaths[j].Route); k++ {
					if NameInRoute(paths[i].Route, tempPaths[j].Route[k]) {
						unique = false
					}
				}
			}
			if unique {
				if !reflect.DeepEqual(tempPaths[0], paths[i]) {
					tempPaths = append(tempPaths, paths[i])
				}
			} else {
				unique = true
			}
		}
		sort.SliceStable(tempPaths, func(i, j int) bool { return tempPaths[i].Length < tempPaths[j].Length })
		tempPaths = AssignAnts(tempPaths, numOfAnts)
		numberOfTurns := tempPaths[len(tempPaths)-1].Length + len(tempPaths[len(tempPaths)-1].ants)
		if numberOfTurns < TurnsToBeat || TurnsToBeat == 0 {
			TurnsToBeat = numberOfTurns
			tempFinalPaths = tempPaths
		}
		tempPaths = make([]Path, 0)
	}
	return tempFinalPaths
}
