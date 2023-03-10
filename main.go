package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

func soal1() {
	x := []int{1, 4, 1, 2, 4, 2, 7, 9, 4, 5}
	yAxis := ""
	xAxis := "  "

	for i := 9; i >= 1; i-- {
		yAxis = fmt.Sprintf("%d ", i)

		for j := 0; j < len(x); j++ {
			if x[j] >= i {
				yAxis += "# "
			} else {
				yAxis += "- "
			}
			if i == 9 {
				xAxis += fmt.Sprintf("%d ", j)
			}
		}
		fmt.Println(yAxis)
	}

	fmt.Println("  " + strings.Repeat("- ", len(x)))
	fmt.Println(xAxis)
}

func hitungTipeData(jsonData interface{}, count map[string]int) {
	switch tipeKey := jsonData.(type) {
	case string:
		count["string"]++
	case bool:
		count["boolean"]++
	case int, float64:
		count["number"]++
	case map[string]interface{}:
		for _, key := range tipeKey {
			hitungTipeData(key, count)
		}
	}
}

func soal2() {
	jsonFile, err := ioutil.ReadFile("mardy.json")
	if err != nil {
		fmt.Println(err)
	}

	var jsonData map[string]interface{}

	err = json.Unmarshal(jsonFile, &jsonData)
	if err != nil {
		fmt.Println(err)
	}
	// jsonDataStr, _ := json.Marshal(jsonData)
	// fmt.Println(string(jsonDataStr))

	totalOutput := map[string]int{"string": 0, "boolean": 0, "number": 0}
	for _, key := range jsonData {
		hitungTipeData(key, totalOutput)
	}

	fmt.Println("string:", totalOutput["string"])
	fmt.Println("boolean:", totalOutput["boolean"])
	fmt.Println("number:", totalOutput["number"])
}

func main() {
	fmt.Println("============ Soal 1 ============")
	soal1()
	fmt.Println("============ Soal 2 ============")
	soal2()
}
