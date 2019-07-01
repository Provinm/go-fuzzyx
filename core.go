package fuzzyx

import (
	"sort"
	"math"
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

// TypeFinal mapping 表 {"ou": [1,2,3]}
type TypeFinal map[string] [][]int 
// TypeListWords 包含列表字符串的结构
type TypeListWords map[string] TypeFinal

// ListWords 所有词汇组成的列表信息
type ListWords struct {
	words []string
	info  TypeListWords
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
func (wl *ListWords) AanlyzeListWord() TypeListWords {

	res := TypeListWords{}
	for outidx, val := range wl.words {
		NewWords := Words{}
		NewWords.name = val
		NewWords.info = NewWords.AanlyzeWords()
		for idx, item := range NewWords.info {
			word := item.info[0]
			newItem := []int{outidx, idx, word.Tone}
			if subCon, ok := res[word.Consonant]; ok {
				 if subFinal, iok := subCon[word.Final]; iok {
					 subFinal = append(subFinal, newItem)
				 } else {
					 temp := make([][]int, 0)
					 temp = append(temp, newItem)
					 subCon[word.Final] = temp
				 }
			} else {
				temp := make([][]int, 0)
				temp = append(temp, newItem)
				newF := map[string] [][]int {word.Final:temp}
				res[word.Consonant] = newF
			}
		}
	}
	return res
}


// TypeRes 结果
type TypeRes []float64

func revertIndexRank(w Words, wl TypeListWords, top int) []int {

	// res := make([]int, 0)

	// outterIndex: innerIndex: rank
	// mapping := make(map[int]map[int]float64) 
	mapping := make(map[int]int)
	for _, word := range w.info {
		for _, wd := range word.info {
			c, f := wd.Consonant, wd.Final
			totalranks, allres := initialMatch(c, f, wl)
			for idx, r := range allres {
				rank := float64(totalranks[idx])
				for _, temp := range r {

					outIndex, tone := temp[0], temp[2]
					if tone != wd.Tone {
						rank *= 0.9
					}
					if innerMap, ok := mapping[outIndex]; ok {
						mapping[outIndex] += innerMap
					} else {
						mapping[outIndex] = 1
					}
				}
			}
		}
	}

	forSort := make([]int, 0)
	valueKeyMapping := make(map[int]int)

	for out, in := range mapping{
		forSort = append(forSort, in)
		valueKeyMapping[in] = out
	}

	sort.Ints(forSort)
	forSort = forSort[:top]

	res := make([]int, 0)

	for _, item := range forSort {
		res = append(res, valueKeyMapping[item])
	}

	return res
}


// initialMatch 计算声母韵母与倒排索引结构结果
func initialMatch(c string, f string, wl TypeListWords) ([]int, [][][]int) {
	relatedC := InitialRank[c]
	res := make([][][]int, 0)
	ranks := make([]int, 0)
	for ini, rank := range relatedC {
		if fin, ok := wl[ini]; ok {
			franks, fres := finalMatch(f, fin)
			for i, rankval := range franks {
				fr := fres[i]
				ranks = append(ranks, rankval*rank)
				res = append(res, fr)
			}
		}
	}

	return ranks, res
} 


// finalMatch 计算韵母与倒排索引内结构结果
func finalMatch(f string, fin TypeFinal) ([]int, [][][]int) {
	relatedF := FinalRank[f]
	res := make([][][]int, 0)
	ranks := make([]int, 0)
	for f, rank := range relatedF {
		if r, ok := fin[f]; ok {
			ranks = append(ranks, rank)
			res = append(res, r)
		}
	}
	return ranks, res
}


// livenstein
func livenstein(w1 Words, w2 Words) float64 {
	// var res float64
	var length, width int

	if len(w1.info) >= len(w2.info) {
		length, width = len(w1.info), len(w2.info)
	} else {
		length, width = len(w2.info), len(w1.info)
	}

	matrix := make([][]float64, width+1)
	for i:=0; i<=width; i++ {
		matrix[i] = make([]float64, length+1)
	}


	for rowIdx, row := range matrix {
		for colIdx, _ := range row {
			if rowIdx == 0 {
				matrix[rowIdx][rowIdx] = float64(colIdx)
			} else if colIdx == 0 {
				matrix[colIdx][colIdx] = float64(rowIdx)
			} else {
				matrix[colIdx][colIdx] = math.Min(
					math.Min(
						matrix[rowIdx-1][rowIdx]+1.0,
						matrix[rowIdx][rowIdx-1]+1.0,
					),
					getEditSpace(w1.info[colIdx].info[0], w2.info[rowIdx].info[0]),
				)
			}
		}
	}

	return matrix[width][length]
}


func getEditSpace(w1 AanlyzedWord, w2 AanlyzedWord) float64 {
	
	iniList := InitialRank[w1.Consonant]
	iniRank := 0.0
	for v, rank := range iniList {
		if v == w2.Consonant {
			iniRank = float64(rank)
			break
		}
	}

	finList := FinalRank[w1.Final]
	finRank := 0.0
	for v, rank := range finList {
		if v == w2.Final {
			finRank = float64(rank)
			break
		}
	}

	typeRank := 1.0
	if w1.Tone != w2.Tone {
		typeRank = 0.9
	}

	res := (iniRank * finRank * typeRank) / 10000
	return res

}


func fuzzyx(w1 string, wl []string, top int) map[string]float64 {

	words := Words{}
	words.name = w1
	words.info = words.AanlyzeWords()

	lw := ListWords{}
	lw.words = wl
	lw.info = lw.AanlyzeListWord()

	reversIndexRes := revertIndexRank(words, lw.info, top)

	ranks := make(map[float64]int)
	forSort := make([]float64, 0)

	w2 := Words{}
	for _, idx := range reversIndexRes {
		w2.name = wl[idx]
		w2.info = w2.AanlyzeWords()
		editSpaceRes := livenstein(words, w2)

		ranks[editSpaceRes] = idx
		forSort = append(forSort, editSpaceRes)
	}

	sort.Float64s(forSort)

	res := make(map[string] float64)
	for _, f := range forSort {
		idx := ranks[f]
		res[wl[idx]] = f
	}

	return res
}