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
	fmt.Println(lw.info)
}
