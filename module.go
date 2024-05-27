package envlist

import (
	"fmt"
	"sort"
	"strings"
)

type EnvList []string

// VarWithPrefix get environment variables with a specific prefix
func (el EnvList) VarWithPrefix(pfx string) EnvList {
	results := EnvList{}
	for _, row := range el {
		if strings.HasPrefix(row, pfx) {
			results = append(results, row)
		}
	}
	return results
}

// VarContains get environment variables containing substring
func (el EnvList) VarContains(pfx string) EnvList {
	results := EnvList{}
	for _, row := range el {
		if strings.Contains(row, pfx) {
			results = append(results, row)
		}
	}
	return results
}

// VarKeyContains get environment variables where the key contains a substring
func (el EnvList) VarKeyContains(sub string) EnvList {
	results := EnvList{}
	for _, row := range el {
		key, _, _ := strings.Cut(row, "=")
		if strings.Contains(key, sub) {
			results = append(results, row)
		}
	}
	return results
}

// VarValueContains get environment variables where the value contains a substring
func (el EnvList) VarValueContains(sub string) EnvList {
	results := EnvList{}
	for _, row := range el {
		_, value, _ := strings.Cut(row, "=")
		if strings.Contains(value, sub) {
			results = append(results, row)
		}
	}
	return results
}

// SortAsc environment variables sort ascending
func (el EnvList) SortAsc() EnvList {
	sort.Sort(sort.StringSlice(el))
	return el
}

// SortDesc environment variables sort descending
func (el EnvList) SortDesc() EnvList {
	sort.Sort(sort.Reverse(sort.StringSlice(el)))
	return el
}

// Print print all/filtered environment variables via 'fmt.Println'
func (el EnvList) Print() {
	for _, val := range el {
		fmt.Println(val)
	}
}
