package lemin

import "sort"

/*
AssignAnts adopts a algorithm described here (https://medium.com/@jamierobertdawson/lem-in-finding-all-the-paths-and-deciding-which-are-worth-it-2503dffb893).
Ants are assigned to paths and then the paths are returned.
*/
func AssignAnts(paths []Path, numOfAnts int) []Path {
	paths[0].ants = append(paths[0].ants, Ant{1})
	for i := 1; i < numOfAnts; i++ {
		pathNum := AntPathDecider(paths)
		paths[pathNum].ants = append(paths[pathNum].ants, Ant{i + 1})
	}
	sort.SliceStable(paths, func(i, j int) bool { return paths[i].Length+len(paths[i].ants) < paths[j].Length+len(paths[j].ants) })
	return paths
}

func AntPathDecider(paths []Path) int {
	routeNum := 0
	length := len(paths)
	for i := 0; i < length-1; i++ {
		if paths[routeNum].Length+len(paths[routeNum].ants) <= paths[routeNum+1].Length+len(paths[routeNum+1].ants) {
			return routeNum
		} else {
			routeNum++
		}
	}
	return routeNum
}
