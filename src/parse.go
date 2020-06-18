package src

import (
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func index(arr []byte, b byte) int {
	for i, j := range arr {
		if j == b {
			return i
		}
	}
	return -1
}

func countCharacters(str string, r rune) int {
	count := 0
	for _, i := range str {
		if i == r {
			count++
		}
	}
	return count
}

func Parse(filename string) (int, string, string, []string, []string, bool) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	output := string(file)

	arr := []string{}
	for strings.Contains(output, "\n") {
		arr = append(arr, output[:strings.Index(output, "\n")])
		output = output[strings.Index(output, "\n")+1:]
	}
	if len(output) != 0 {
		arr = append(arr, output)
	}

	number, start, end, locations, relations, check := isCorrect(arr)
	return number, start, end, locations, relations, check
}

func trimSpaces(str string) string {
	for i := 0; i < len(str); i++ {
		if str[i] != ' ' && str[i] != '\t' {
			str = str[i:]
			break
		}
	}
	for i := len(str) - 1; i >= 0; i-- {
		if str[i] != ' ' && str[i] != '\t' {
			str = str[:i+1]
			break
		}
	}
	return str
}

func removeDuplicateSpaces(arr []string) []string {
	for i := range arr {
		space := regexp.MustCompile(`\s+`)
		arr[i] = space.ReplaceAllString(trimSpaces(arr[i]), " ")
	}
	return arr
}

func isCorrectAntNumber(str string) bool {
	_, err := strconv.Atoi(str)
	return err == nil
}

func isCorrectLocation(str string) bool {
	arr := strings.Split(str, " ")
	if len(arr) != 3 {
		return false
	}
	if str[0] == 'L' {
		return false
	}
	_, ex := strconv.Atoi(arr[1])
	_, ey := strconv.Atoi(arr[2])
	if ex != nil || ey != nil {
		return false
	}
	return true
}

func isCorrectRelation(str string) bool {
	arr := strings.Split(str, "-")
	if len(arr) != 2 {
		return false
	}

	return true
}

func isCorrectStartEnd(arr *[]string) (string, string, bool) {
	isStart, isEnd := false, false
	start, end := "", ""
	for i := 0; i < len((*arr)); i++ {
		if (*arr)[i] == "##start" && i+1 < len((*arr)) && isCorrectLocation((*arr)[i+1]) {
			isStart = true
			start = strings.Split((*arr)[i+1], " ")[0]
			(*arr) = append((*arr)[:i], (*arr)[i+1:]...)
			i--
		} else if (*arr)[i] == "##end" && i+1 < len((*arr)) && isCorrectLocation((*arr)[i+1]) {
			isEnd = true
			end = strings.Split((*arr)[i+1], " ")[0]
			(*arr) = append((*arr)[:i], (*arr)[i+1:]...)
			i--
		}
	}
	return start, end, isStart && isEnd
}

func isComment(str string) bool {
	return strings.Index(str, "#") == 0 && strings.Index(str, "##") != 0
}

func removeComments(arr []string) []string {
	for i := 0; i < len(arr); i++ {
		if isComment(arr[i]) {
			arr = append(arr[:i], arr[i+1:]...)
			i--
		} else if arr[i] == "" {
			arr = append(arr[:i], arr[i+1:]...)
			i--
		}
	}
	return arr
}

func firstRelationIndex(arr []string) int {
	for i, j := range arr {
		tmp := strings.Split(j, "-")
		if len(tmp) == 2 {
			return i
		}
	}
	return -1
}

func areUnique(arr []string) bool {
	for i := range arr {
		for j := range arr {
			if i != j && arr[i] == arr[j] {
				return false
			}
		}
	}
	return true
}

func isCorrectLocations(arr []string) ([]string, bool) {
	names := []string{}
	locations := []string{}
	for _, j := range arr {
		if !isCorrectLocation(j) {
			return nil, false
		}
		tmp := strings.Split(j, " ")
		names = append(names, tmp[0])
		locations = append(locations, tmp[1]+" "+tmp[2])
		if !areUnique(names) || !areUnique(locations) {
			return nil, false
		}
	}
	return names, true
}

func containsArr(arr []string, str string) bool {
	for _, i := range arr {
		if i == str {
			return true
		}
	}
	return false
}

func isCorrectRelations(arr, names []string) bool {
	for _, i := range arr {
		if !isCorrectRelation(i) {
			return false
		}
		tmp := strings.Split(i, "-")
		if !containsArr(names, tmp[0]) || !containsArr(names, tmp[1]) {
			return false
		}
	}
	for i, j := range arr {
		for l, m := range arr {
			if i != l {
				tmp1 := strings.Split(j, "-")
				tmp2 := strings.Split(m, "-")
				if (tmp1[0] == tmp1[1]) || (tmp2[0] == tmp2[1]) || (tmp1[0] == tmp2[0] && tmp1[1] == tmp2[1]) || (tmp1[0] == tmp2[1] && tmp1[1] == tmp2[0]) {
					return false
				}
			}
		}
	}

	return true
}

func isCorrect(arr []string) (int, string, string, []string, []string, bool) {
	arr = removeDuplicateSpaces(arr)
	arr = removeComments(arr)
	if len(arr) < 5 {
		return 0, "", "", nil, nil, false
	}
	number, err := strconv.Atoi(arr[0])
	if err != nil || number < 1 {
		return 0, "", "", nil, nil, false
	}
	locations := arr[1 : firstRelationIndex(arr[1:])+1]
	relations := arr[firstRelationIndex(arr):]

	start, end, checkStartEnd := isCorrectStartEnd(&locations)
	if !checkStartEnd {
		return 0, "", "", nil, nil, false
	}

	names, checkLocations := isCorrectLocations(locations)
	if !checkLocations {
		return 0, "", "", nil, nil, false
	}

	if !isCorrectRelations(relations, names) {
		return 0, "", "", nil, nil, false
	}

	return number, start, end, locations, relations, true
}
