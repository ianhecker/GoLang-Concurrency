package read

import (
	"encoding/json"
	"io/ioutil"
)

func ReadFile(filepath string) *[]int {
	var numbers []int
	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(bytes, &numbers)
	if err != nil {
		panic(err)
	}
	return &numbers
}
