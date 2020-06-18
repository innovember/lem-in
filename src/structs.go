package src

import (
	"fmt"
	"strings"
)

type Pair struct {
	Name string
}

func (p Pair) Print() {
	fmt.Print(p.Name)
}

type Graph struct {
	Rel map[string][]Pair
}

func (g *Graph) Init(loc, rel []string) {
	g.Rel = make(map[string][]Pair)
	g.InitLocations(loc, rel)
}

func (g Graph) Print() {
	for key, val := range g.Rel {
		fmt.Printf("%-10s:", key)
		for _, i := range val {
			fmt.Print("  ")
			i.Print()
		}
		fmt.Println()
	}
}

func contains(arr []Pair, str string) bool {
	for _, p := range arr {
		if p.Name == str {
			return true
		}
	}
	return false
}

func (g *Graph) InitLocations(loc, rel []string) {
	for _, i := range loc {
		arr := strings.Split(i, " ")
		for _, j := range rel {
			arr2 := strings.Split(j, "-")
			if arr[0] == arr2[0] || arr[0] == arr2[1] {
				if !contains(g.Rel[arr2[0]], arr2[1]) {
					g.Rel[arr2[0]] = append(g.Rel[arr2[0]], Pair{arr2[1]})
				}
				if !contains(g.Rel[arr2[1]], arr2[0]) {
					g.Rel[arr2[1]] = append(g.Rel[arr2[1]], Pair{arr2[0]})
				}
			}
		}
	}
}

func findLocation(arr []string, str string) string {
	for _, s := range arr {
		a := strings.Split(s, " ")
		if a[0] == str {
			return s
		}
	}
	return ""
}
