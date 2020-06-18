package src

import (
	"fmt"
)

type Path struct {
	Route []string
}

func (p Path) index(str string) int {
	for i, j := range p.Route {
		if j == str {
			return i
		}
	}
	return -1
}

func (p *Path) append(pa Pair) {
	p.Route = append(p.Route, pa.Name)
}

func (p *Path) remove(pa Pair) {
	index := p.index(pa.Name)
	p.Route = append(p.Route[:index], p.Route[index+1:]...)
}

func (p Path) copy() Path {
	var tmp Path
	tmp.Route = make([]string, len(p.Route))
	copy(tmp.Route, p.Route)
	return tmp
}

func (p Path) print() {
	fmt.Println(p.Route)
}

type Paths struct {
	Arr []Path
}

func (ps Paths) Print() {
	for _, i := range ps.Arr {
		i.print()
	}
}

func (ps *Paths) Sort() {
	for i := 0; i < len(ps.Arr); i++ {
		for j := i + 1; j < len(ps.Arr); j++ {
			if len(ps.Arr[i].Route) > len(ps.Arr[j].Route) {
				(*ps).Arr[i], (*ps).Arr[j] = (*ps).Arr[j], (*ps).Arr[i]
			}
		}
	}
}

func (ps Paths) toSlice() [][]string {
	var arr [][]string
	for _, i := range ps.Arr {
		arr = append(arr, i.Route)
	}
	return arr
}

func (ps *Paths) append(arr [][]string) {

}

func (ps Paths) copy() Paths {
	var c Paths
	for i := range ps.Arr {
		c.Arr = append(c.Arr, ps.Arr[i].copy())
	}
	return c
}

func (g *Graph) InitPaths(start, end string) Paths {
	var ps Paths
	visited := Path{[]string{start}}
	g.helper(&ps, Pair{start}, visited, end)
	return ps
}

func (p Paths) isUnique(arr []string) bool {
	for _, i := range p.Arr {
		for _, j := range i.Route {
			for _, k := range arr {
				if j == k {
					return false
				}
			}
		}
	}
	return true
}

func (g *Graph) helper(ps *Paths, cur Pair, visited Path, end string) {
	if cur.Name == end {
		(*ps).Arr = append((*ps).Arr, visited.copy())
		return
	}
	for _, next := range g.Rel[cur.Name] {
		if visited.index(next.Name) == -1 {
			visited.append(next)
			g.helper(ps, next, visited, end)
			visited.remove(next)
		}
	}
}

func isUnique(arr1, arr2 []string) bool {
	arr := append(arr1, arr2...)
	for i := range arr {
		for j := range arr {
			if i != j && arr[i] == arr[j] {
				return false
			}
		}
	}
	return true
}

func contains2(arr []string, str string) bool {
	for _, i := range arr {
		if i == str {
			return true
		}
	}
	return false
}

func isSubstring(visited, mid []string) bool {
	for _, i := range mid {
		if !contains2(visited, i) {
			return false
		}
	}
	return true
}

func maxIndex(arr []int) int {
	max, mindex := 0, 0
	for i, j := range arr {
		if max < j {
			max = j
			mindex = i
		}
	}
	return mindex
}

func DeleteCross(ps Paths) [][]string {
	best := [][]string{}
	bestInt := []int{}
	indexes := [][]int{}
	for i := range ps.Arr {
		visited := make([]string, len(ps.Arr[i].Route)-2)
		indexes = append(indexes, []int{})
		tmp := 1
		copy(visited, ps.Arr[i].Route[1:len(ps.Arr[i].Route)-1])
		for j := i + 1; j < len(ps.Arr); j++ {
			if isUnique(visited, ps.Arr[j].Route) {
				indexes[i] = append(indexes[i], len(visited))
				visited = append(visited, ps.Arr[j].Route[1:len(ps.Arr[j].Route)-1]...)
				tmp++
			}
		}
		indexes[i] = append(indexes[i], len(visited))
		best = append(best, visited)
		bestInt = append(bestInt, tmp)
	}
	arr := [][]string{}
	start := 0
	for i := range indexes[maxIndex(bestInt)] {
		n := indexes[maxIndex(bestInt)][i]
		arr = append(arr, best[maxIndex(bestInt)][start:n])
		start = n
	}
	if len(arr[0]) == 0 {
		arr = [][]string{arr[0]}
	}
	return arr
}
