package workshop

import (
	"sort"
	"strconv"
)

func Sort(input ...string) {
	var segments [][]interface{}
	for _, str := range input {
		runes := []rune(str)

		rawSegments := splitSegments(runes...)
		typedSegments := convertSegments(rawSegments...)

		segments = append(segments, typedSegments)
	}

	st := &sorter{
		len: func() int { return len(input) },
		less: func(i, j int) bool {
			seg1 := segments[i]
			seg2 := segments[j]

			return less(seg1, seg2)
		},
		swap: func(i, j int) {
			input[i], input[j] = input[j], input[i]
			segments[i], segments[j] = segments[j], segments[i]
		},
	}
	sort.Sort(st)
}

func less(seg1, seg2 []interface{}) bool {
	for k, l := 0, 0; k < len(seg1) && l < len(seg2); k, l = k+1, l+1 {
		switch v1 := seg1[k].(type) {
		case string:
			switch v2 := seg2[l].(type) {
			case string:
				if v1 == v2 {
					continue
				}
				return v1 < v2
			case int64:
				return false
			default:
				panic("invalid type")
			}
		case int64:
			switch v2 := seg2[l].(type) {
			case string:
				return true
			case int64:
				if v1 == v2 {
					continue
				}
				return v1 < v2
			default:
				panic("invalid type")
			}
		default:
			panic("invalid type")
		}
	}
	return false
}

func convertSegments(rawSegments ...[]rune) []interface{} {
	var (
		segments []interface{}
	)
	for _, raw := range rawSegments {
		str := string(raw)

		if '0' <= raw[0] && raw[0] <= '9' {
			n, err := strconv.ParseInt(str, 10, 64)
			if err != nil {
				panic(err)
			}
			segments = append(segments, n)
			continue
		}

		segments = append(segments, str)
	}
	return segments
}

func splitSegments(runes ...rune) [][]rune {
	var (
		lastCond          bool
		currentRawSegment []rune
		rawSegments       [][]rune
	)
	for _, r := range runes {
		cond := '0' <= r && r <= '9'

		if lastCond == cond {
			currentRawSegment = append(currentRawSegment, r)
		} else {
			lastCond = cond
			if len(currentRawSegment) > 0 {
				rawSegments = append(rawSegments, currentRawSegment)
				currentRawSegment = []rune{r}
			}
		}
	}

	if len(currentRawSegment) > 0 {
		rawSegments = append(rawSegments, currentRawSegment)
	}

	return rawSegments
}

type sorter struct {
	len  func() int
	less func(i, j int) bool
	swap func(i, j int)
}

func (st *sorter) Len() int           { return st.len() }
func (st *sorter) Less(i, j int) bool { return st.less(i, j) }
func (st *sorter) Swap(i, j int)      { st.swap(i, j) }
