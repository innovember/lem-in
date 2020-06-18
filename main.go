package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"./src"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("ERROR: invalid arguments")
		return
	}
	antNum, start, end, locations, relations, check := src.Parse(os.Args[1])

	if !check {
		fmt.Println("ERROR: invalid data format")
		os.Exit(1)
	}

	var g src.Graph
	g.Init(locations, relations)
	// g.Print()

	// fmt.Println()

	ps := g.InitPaths(start, end)
	ps.Sort()
	// ps.Print()
	// fmt.Println()
	arr := src.DeleteCross(ps)
	for i := range arr {
		arr[i] = append([]string{}, arr[i]...)
		arr[i] = append(arr[i], end)
	}
	fmt.Println(antNum)

	pathsNum := len(arr)
	ants := make([]int, pathsNum)
	cnt := 0
	ants[cnt]++
	antNum--
	for antNum != 0 {
		if cnt == pathsNum-1 {
			if len(arr[cnt])+ants[cnt] <= len(arr[0])+ants[0] {
				ants[cnt]++
			} else {
				ants[0]++
			}
			antNum--
			cnt = 0
		} else {
			if len(arr[cnt])+ants[cnt] <= len(arr[cnt+1])+ants[cnt+1] {
				ants[cnt]++
			} else {
				ants[cnt+1]++
			}
			antNum--
			cnt++
		}
	}
	//path := make([]string, antNum)
	antIndex := 1
	var steps [][]string
	for i, v := range ants {
		var path [][]string
		cnt = 0
		for j := 0; j < v; j++ {
			var ans []string
			for _, room := range arr[i] {
				ans = append(ans, "L"+strconv.Itoa(antIndex)+"-"+room)
			}
			cnt++
			antIndex++
			path = append(path, ans)
		}
		//fmt.Printf("%#v\n", path)
		res := make([][]string, len(path)-1+len(path[0]))
		inc := 0
		for _, p := range path {
			for ind, r := range p {
				res[ind+inc] = append(res[ind+inc], r)
			}
			inc++
		}
		for ind, step := range res {
			if ind > len(steps)-1 {
				steps = append(steps, step)
			} else {
				steps[ind] = append(steps[ind], step...)
			}
		}

	}
	for _, v := range locations {
		tmp := strings.Split(v, " ")[0]
		if start == tmp {
			fmt.Println("##start")
		}
		if end == tmp {
			fmt.Println("##end")
		}
		fmt.Println(v)
	}
	fmt.Println(strings.Join(relations, "\n"))
	fmt.Println()
	for _, v := range steps {
		fmt.Println(strings.Join(v, " "))
	}

}
