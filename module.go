package envlist

import (
	"fmt"
	"strings"
)

type EnvList []string

// VarPrefix get environment variables with a specific prefix
func (el EnvList) VarPrefix(pfx string) EnvList {
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

// PrintItems print all/filtered environment variables
func (el EnvList) PrintItems() {
	for _, val := range el {
		fmt.Println(val)
	}
}
