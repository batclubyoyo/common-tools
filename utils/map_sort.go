package utils

import (
	"fmt"
	"sort"
)

/*
Pair A data structure to hold a key/value pair.
要对map按照value进行排序，思路是直接不用map，用struct存放key和value，实现sort接口，就可以调用sort.Sort进行排序了。
*/
type Pair struct {
	Key   string
	Value int
}

/*
PairList A slice of Pairs that implements sort.Interface to sort by Value.
*/
type PairList []Pair

func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }

/*
SortMapByValueAsc A function to turn a map into a PairList, then sort and return it.
*/
func SortMapByValueAsc(m map[string]int) PairList {
	p := make(PairList, len(m))
	i := 0
	for k, v := range m {
		p[i] = Pair{k, v}
	}
	sort.Sort(p)
	return p
}

/*
SortMapByValueDesc A function to turn a map into a PairList, then sort and return it.
*/
func SortMapByValueDesc(m map[string]int) PairList {
	p := make(PairList, len(m))
	i := 0
	for k, v := range m {
		p[i] = Pair{k, v}
	}
	// sort.Sort(sort.Reverse(p))
	_ = sort.Reverse(p)
	return p
}

/*
SortMap Sort Map
*/
func SortMap(mp map[string]int) map[string]int {
	var newMp = make([]int, 0)
	var newMpKey = make([]string, 0)

	var newMap map[string]int = make(map[string]int)

	for oldk, v := range mp {
		newMp = append(newMp, v)
		newMpKey = append(newMpKey, oldk)
	}
	sort.Ints(newMp)
	fmt.Println(newMp)
	for k, v := range newMp {
		// fmt.Printf("根据value排序后的新集合为:%v: [%v]=%v \n", k, newMpKey[k], v)
		newMap[newMpKey[k]] = v
	}

	return newMap
}
