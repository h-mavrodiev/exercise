package main

import "fmt"

// func printInOrder(m map[string][]string, key string) {
// 	if len(m[key]) == 0 {
// 		fmt.Println(key)
// 	}
// }

// func print(m map[string][]string) {
// 	var visited []string

// 	for k, v := range m {
// 		if len(v) == 0 && !contains(k, visited) {
// 			visited = append(visited, k)
// 		} else {
// 			for _, i := range v {
// 				if contains(i, visited) {
// 					continue
// 				}
// 				visited = append(visited, i)
// 			}
// 		}
// 	}

// 	for i, v := range visited {
// 		fmt.Printf("%d - %s\n", i, v)
// 	}
// }

// func contains(val string, vals []string) bool {
// 	var contains bool
// 	for _, i := range vals {
// 		if i == val {
// 			contains = true
// 		}
// 	}
// 	return contains
// }

func getBuildOrder(packageName string, dependencies map[string][]string) []string {
	var buildOrder []string
	visited := make(map[string]bool)
	var visit func(string)
	visit = func(packageName string) {
		if !visited[packageName] {
			visited[packageName] = true
			for _, dependency := range dependencies[packageName] {
				visit(dependency)
			}
			buildOrder = append(buildOrder, packageName)
		}
	}
	visit(packageName)
	return buildOrder
}

func main() {
	ex1 := map[string][]string{
		"A": {"F", "C"},
		"B": {"D"},
		"C": {},
		"D": {},
		"E": {},
		"F": {"A"},
		"G": {"B"},
	}

	order := getBuildOrder("A", ex1)

	for i, v := range order {
		fmt.Printf("%d - %s\n", i, v)
	}

}

// Print the "packages" in order, taking the dependencies in consideration
