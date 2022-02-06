package assignment

import (
	"fmt"
	"math"
	"reflect"
	"sort"
	"strings"
)

func AddUint32(x, y uint32) (uint32, bool) {
	sum := x + y
	overflow := sum < x || sum < y
	return sum, overflow
}

func CeilNumber(f float64) float64 {
	return math.Ceil(f*4) / 4
}

func AlphabetSoup(s string) string {
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})

	return string(r)
}

func StringMask(s string, n uint) string {
	slen := len(s)
	m := int(n)
	if slen == 0 {
		return "*"
	}

	rs := []rune(s)

	if m == 0 || m >= slen {
		for i, _ := range rs {
			rs[i] = '*'
		}
	} else {
		for i, r := range rs {
			if m <= 0 && r != '*' {
				rs[i] = '*'
			} else {
				m--
			}
		}
	}

	return string(rs)
}

func WordSplit(arr [2]string) interface{} {
	text := arr[0]
	words := strings.Split(arr[1], ",")
	var result []string

	for _, word := range words {
		if b := strings.Contains(text, word); b {
			result = append(result, word)
		}
	}

	if len(result) > 1 {
		return result
	} else {
		return "not possible"
	}
}

func VariadicSet(i ...interface{}) []interface{} {
	m := make(map[string]interface{})
	sb := strings.Builder{}
	for _, v := range i {
		sb.Reset()
		sb.WriteString(reflect.TypeOf(v).String())
		sb.WriteString("-")
		sb.WriteString(fmt.Sprintf("%v", v))
		m[sb.String()] = v
	}

	values := make([]interface{}, 0, len(m))

	for _, v := range m {
		values = append(values, v)
	}

	return values
}
