package main

import (
	"fmt"
	lemin "lemin/actions"
	"os"
	"strings"
)

func main() {
	// Read file and split into readable data
	file, _ := os.ReadFile(os.Args[1])
	fileStrings := strings.Split(string(file), "\n")
	//Build nodes with file data
	numOfAnts, rooms := lemin.MapBuilder(fileStrings)
	if numOfAnts < 1 {
		fmt.Println("ERROR: invalid data format, invalid number of ants")
	} else {
		//Find optimised paths to take
		paths := lemin.PathFinder(numOfAnts, rooms)
		if len(paths) < 1 {
			fmt.Println("ERROR: invalid data format")
		} else {
			//Print data information
			for _, v := range fileStrings {
				fmt.Println(v)
			}
			fmt.Println()
			//Simulate ants moving through the best routes
			lemin.RouteTaker(paths, numOfAnts)
		}
	}
	//DOESN'T WORK YET - Testing to make picture of map
	lemin.MapImage()
}
