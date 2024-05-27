package envlist

import (
	"bytes"
	"strings"
	"testing"
)

const EnvListMock = "SHELL=/bin/bash,USERNAME=something,HOME=/home/gocias"

func TestVarWithPrefix(t *testing.T) {
	var mock EnvList = strings.Split(EnvListMock, ",")
	if len(mock.VarWithPrefix("HOME")) == 0 {
		t.Error("Prefix 'HOME' not found!")
	}
}

func TestVarContains(t *testing.T) {
	var mock EnvList = strings.Split(EnvListMock, ",")
	if len(mock.VarContains("bin")) == 0 {
		t.Error("Substring 'bin' not found!")
	}
}

func TestVarKeyContains(t *testing.T) {
	var mock EnvList = strings.Split(EnvListMock, ",")
	if len(mock.VarKeyContains("ERN")) == 0 {
		t.Error("Key with substring 'ERN' not found!")
	}
}

func TestVarValueContains(t *testing.T) {
	var mock EnvList = strings.Split(EnvListMock, ",")
	if len(mock.VarValueContains("goc")) == 0 {
		t.Error("Value with substring 'goc' not found!")
	}
}

func TestSortAsc(t *testing.T) {
	var mock EnvList = strings.Split(EnvListMock, ",")
	exp := []string{"HOME=/home/gocias", "SHELL=/bin/bash", "USERNAME=something"}
	res := bytes.Equal([]byte(strings.Join(mock.SortAsc(), " ")), []byte(strings.Join(exp, " ")))
	if !res {
		t.Error("Sort by 'asc' is failed!")
	}
}

func TestSortDesc(t *testing.T) {
	var mock EnvList = strings.Split(EnvListMock, ",")
	exp := []string{"USERNAME=something", "SHELL=/bin/bash", "HOME=/home/gocias"}
	res := bytes.Equal([]byte(strings.Join(mock.SortDesc(), " ")), []byte(strings.Join(exp, " ")))
	if !res {
		t.Error("Sort by 'desc' is failed!")
	}
}
