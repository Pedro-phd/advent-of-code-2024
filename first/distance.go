package first

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func Setup() {
	// firstChallenge()
	secondChallenge()
}

type col struct {
	col_01 []int
	col_02 []int
}

func transformInput() col {
	file, err := os.ReadFile("./first/input_final.txt")
	if err != nil {
		panic(err)
	}

	col_01 := []int{}
	col_02 := []int{}

	data := string(file)
	lines := strings.Split(data, "\n")

	for _, line := range lines {
		line = normalizeLine(line)      // normalize line, change to any space to one space
		col := strings.Split(line, " ") // split to columns

		number01, err := strconv.ParseInt(col[0], 10, 64) //convert to int
		if err != nil {
			panic(err)
		}

		number02, err := strconv.ParseInt(col[1], 10, 64) //convert to int
		if err != nil {
			panic(err)
		}

		col_01 = append(col_01, int(number01)) //input to array
		col_02 = append(col_02, int(number02)) //input to array
	}

	sort.Ints(col_01) //sort array
	sort.Ints(col_02) //sort array

	return col{col_01, col_02}
}

func toPositiveNum(num int) int {
	if num < 0 {
		return num * -1
	}

	return num
}

func normalizeLine(line string) string {
	space := regexp.MustCompile(`\s+`)
	return space.ReplaceAllString(line, " ")
}

func firstChallenge() {
	col := transformInput()
	distance := 0

	for index := range col.col_01 {

		temp := col.col_02[index] - col.col_01[index]
		distance += toPositiveNum(temp) // change to positive number

	}
	fmt.Printf("Distance: %d", distance)
}

func secondChallenge() {
	col := transformInput()

	result := 0

	for _, row01 := range col.col_01 {
		for _, row02 := range col.col_02 {
			if row01 == row02 {
				result += row01
			}
		}
	}

	fmt.Println(result)
}
