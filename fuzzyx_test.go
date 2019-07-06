package fuzzyx

import (
	"fmt"
	"testing"
)

func TestSingleWord(t *testing.T) {
	NewWord := Word{}
	NewWord.name = "波"
	NewWord.info = NewWord.Aanlyze()
	fmt.Println(NewWord.info)
}

func TestWords(t *testing.T) {
	words := Words{}
	words.name = "张三"
	words.info = words.AanlyzeWords()
	fmt.Println(words.info)
}

func TestListWords(t *testing.T) {
	lw := ListWords{}
	lw.words = []string{"张三", "李四"}
	lw.info = lw.AanlyzeListWord()

	for key, val := range lw.info {
		fmt.Println(key)
		for _k, _v := range val {
			fmt.Println(_k)
			for _, _val := range _v {
				fmt.Println(_val)
			}
		}
	}
}


func TestInitialRank(t *testing.T) {
	fmt.Println(InitialRank)
}


func TestFinalRank(t *testing.T) {
	fmt.Println(FinalRank)
}


func TestFuzzyx(t *testing.T) {

	w1 := "张三"
	wl := []string {"张三", "李四", "章三", "掌山"}

	res := fuzzyx(w1, wl, 10)
	// fmt.Println(res)
	t.Log(res)
}
