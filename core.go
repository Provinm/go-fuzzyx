package fuzzyx

import (
	"strconv"

	"github.com/mozillazg/go-pinyin"
)

// Meta
const (
	Author = "partrick.zhou"
	Date   = "2019.6.1"
)

// AanlyzedWord 单个中文字符的声韵母，声调信息
type AanlyzedWord struct {
	Consonant string
	Final     string
	Tone      int
}

// Word 单个中文字符的信息
type Word struct {
	name string
	info []AanlyzedWord
}

// Words 单个词汇信息
type Words struct {
	name string
	info []Word
}

// ListWords 所有词汇组成的列表信息
type ListWords struct {
	words []string
	info  []Words
}

// InitialArgs 获取声母设置
func InitialArgs() pinyin.Args {
	return pinyin.Args{pinyin.Initials, pinyin.Heteronym, pinyin.Separator, pinyin.Fallback}
}

// FinalArgs 获取韵母设置
func FinalArgs() pinyin.Args {
	return pinyin.Args{pinyin.FinalsTone3, pinyin.Heteronym, pinyin.Separator, pinyin.Fallback}
}

// GetInitial 获取声母的函数
func GetInitial(w string) string {
	args := InitialArgs()
	ret := pinyin.Pinyin(w, args)

	if len(ret) == 0 {
		return ""
	}
	return ret[0][0]
}

// GetFinal 获取韵母的函数
func GetFinal(w string) (string, int) {
	args := FinalArgs()
	ret := pinyin.Pinyin(w, args)
	if len(ret) == 0 {
		return "", 0
	}

	FinalTone := ret[0][0]
	f := string(FinalTone[:len(FinalTone)-1])
	t, _ := strconv.Atoi(FinalTone[len(FinalTone)-1:])
	return f, t
}

// Aanlyze 分析每个单字的信息
func (w *Word) Aanlyze() []AanlyzedWord {

	res := make([]AanlyzedWord, 0)
	ch := w.name
	final, tone := GetFinal(ch)
	w1 := AanlyzedWord{Consonant: GetInitial(ch), Final: final, Tone: tone}
	res = append(res, w1)
	return res
}

// AanlyzeWords 分析单个词汇
func (ws *Words) AanlyzeWords() []Word {
	str := ws.name
	res := make([]Word, 0)
	for _, ch := range str {
		Newword := Word{}
		Newword.name = string(ch)
		Newword.info = Newword.Aanlyze()
		res = append(res, Newword)
	}

	return res
}

// AanlyzeListWord 分析包含列表的字符串
func (wl *ListWords) AanlyzeListWord() []Words {

	res := make([]Words, 0)
	for _, val := range wl.words {
		NewWords := Words{}
		NewWords.name = val
		NewWords.info = NewWords.AanlyzeWords()
		res = append(res, NewWords)
	}

	return res
}
