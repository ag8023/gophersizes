package main // Do not change this.

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//fmt.Println(acronymMaker("red ventures"))
	fmt.Println(stringToInt("95"))
}

func acronymMaker(sentence string) string {
	newSentence := strings.ToUpper(sentence)
	res := strings.Split(newSentence, " ")
	var acronym string
	//fmt.Println(res)
	for _, word := range res {

		acronym += string(word[0])
		//fmt.Println(acronym)
	}
	return acronym
}

func stringToInt(myString string) int {
	i, err := strconv.Atoi(myString)
	if err != nil {
		return 0
	}
	return i
}
