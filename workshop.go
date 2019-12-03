package workshop

import (
	"sort"
	"strconv"
	"unicode"
)

func Sort(input ...string) (result []string) {
	st := newSortable(input...)
	sort.Sort(st)
	result = input
	return
}

type sortable struct {
	input      []string
	partitions [][]segment
}

func newSortable(input ...string) *sortable {
	res := &sortable{input: input}
	for _, s := range input {
		p := partition(s)
		res.partitions = append(res.partitions, p)
	}
	return res
}

func (st *sortable) Len() int { return len(st.input) }
func (st *sortable) Less(i, j int) bool {
	var (
		equal = true
		less  = false
		p1    = st.partitions[i]
		p2    = st.partitions[j]
	)
	for k, v := 0, 0; k < len(p1) && v < len(p2); k, v = k+1, v+1 {
		lr := lessSegment(p1[k], p2[v])
		rl := lessSegment(p2[v], p1[k])
		less = less || lr
		equal = equal && !rl && !lr
		if less {
			break
		}
	}
	if equal {
		return len(p1) < len(p2)
	}
	return less
}
func (st *sortable) Swap(i, j int) {
	st.input[i], st.input[j] = st.input[j], st.input[i]
	st.partitions[i], st.partitions[j] = st.partitions[j], st.partitions[i]
}

func lessSegment(sg1, sg2 segment) bool {
	if sg1.str == nil && sg2.str != nil &&
		sg1.num != nil && sg2.num == nil {
		return true
	}

	if sg1.str != nil && sg2.str != nil &&
		sg1.num == nil && sg2.num == nil {
		return *sg1.str < *sg2.str
	}

	if sg1.str == nil && sg2.str == nil &&
		sg1.num != nil && sg2.num != nil {
		return *sg1.num < *sg2.num
	}

	return false
}

func partition(s string) (res []segment) {
	runes := []rune(s)

	runeType := runeNone
	var part []rune
	for _, r := range runes {
		currentRuneType := typeOf(r)

		if runeType == runeNone {
			part = append(part, r)
			runeType = currentRuneType
			continue
		}

		if runeType == currentRuneType {
			part = append(part, r)
			continue
		}

		res = append(res, makeSegment(runeType, string(part)))

		part = nil
		part = append(part, r)
		runeType = currentRuneType
	}

	if part != nil {
		res = append(res, makeSegment(runeType, string(part)))
	}

	return
}

func makeSegment(runeType int, part string) segment {
	switch runeType {
	case runeDigit:
		n, err := strconv.ParseInt(part, 10, 64)
		if err != nil {
			panic(err)
		}
		return segment{num: newNum(int(n))}
	case runeAlphabet:
		return segment{str: newStr(part)}
	}
	panic("invalid runeType")
}

func typeOf(r rune) int {
	if unicode.IsDigit(r) {
		return runeDigit
	}
	return runeAlphabet
}

func newNum(n int) *int {
	res := new(int)
	*res = n
	return res
}

func newStr(s string) *string {
	res := new(string)
	*res = s
	return res
}

type segment struct {
	str *string
	num *int
}

const (
	runeNone     = -1
	runeAlphabet = 100
	runeDigit    = 200
)
