package lemin

import (
	"fmt"
	"sort"
	"strconv"
)

type Ant struct {
	ID int
}

/*
RouteTaker will print a 'move' if there are ants available and there is enough 'room' in the length of the route.

It will make moves for the first ant of each path first before allowing the second ants to make their first moves.
While the second ants are making their first move, the first ants will make their second move.

The program finishes when there are no more available turns (nothing has been printed).
*/
func RouteTaker(paths []Path, numOfAnts int) {
	antCounter := 1
	antRows := FindMaxAntNum(paths)
	routeCounter := 1
	maxRouteCounter := 1
	print := false
	for routeCounter <= maxRouteCounter {
		for i := 0; i < antCounter; i++ {
			for j := 0; j < len(paths); j++ {
				//(routeCounter - i) is making sure ants further back in the queue are still making moves.
				if i < len(paths[j].ants) && routeCounter-i < len(paths[j].Route) {
					idstring := strconv.Itoa(paths[j].ants[i].ID)
					fmt.Print("L" + idstring + "-" + paths[j].Route[routeCounter-i] + " ")
					print = true
				}
			}
		}
		if antCounter < antRows {
			antCounter++
		}
		if print {
			maxRouteCounter++
			print = false
			fmt.Println()
		}
		routeCounter++
	}
}

/*
Finds the highest number of ants in any of the given paths.
*/
func FindMaxAntNum(paths []Path) int {
	sort.SliceStable(paths, func(i, j int) bool { return len(paths[i].ants) > len(paths[j].ants) })
	return len(paths[0].ants)
}
