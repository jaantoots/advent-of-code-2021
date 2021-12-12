package main

import (
	"fmt"
	"io"
	"strings"
)

func check_visited(history []string, cave string) bool {
	for _, visited := range history {
		if cave == visited {
			return true
		}
	}
	return false
}

func find_paths(caves map[string]bool, connections map[string][]string, history []string, extra_visit bool) [][]string {
	current := history[len(history)-1]
	if current == "end" {
		return [][]string{history}
	}
	paths := [][]string{}
	for _, cave := range connections[current] {
		next_extra_visit := extra_visit
		if !caves[cave] && check_visited(history, cave) {
			if extra_visit && cave != "start" {
				next_extra_visit = false
			} else {
				continue
			}
		}
		next_history := make([]string, len(history)+1)
		copy(next_history, history)
		next_history[len(history)] = cave
		for _, path := range find_paths(caves, connections, next_history, next_extra_visit) {
			paths = append(paths, path)
		}
	}
	return paths
}

func main() {
	var line string
	caves := make(map[string]bool, 64)
	connections := make(map[string][]string, 64)
	for {
		_, err := fmt.Scanln(&line)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		s := strings.Split(line, "-")
		if len(s) != 2 {
			panic("bad connection")
		}
		caves[s[0]] = (s[0] == strings.ToUpper(s[0]))
		caves[s[1]] = (s[1] == strings.ToUpper(s[1]))
		connections[s[0]] = append(connections[s[0]], s[1])
		connections[s[1]] = append(connections[s[1]], s[0])
	}
	//fmt.Println(caves)
	//fmt.Println(connections)
	paths := find_paths(caves, connections, []string{"start"}, false)
	//fmt.Println(paths)
	fmt.Println(len(paths))

	extra_paths := find_paths(caves, connections, []string{"start"}, true)
	fmt.Println(len(extra_paths))
}
