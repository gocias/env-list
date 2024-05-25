package envlist

import (
	"strings"
	"testing"
)

const EnvListMok = "SHELL=/bin/bash,USERNAME=something,HOME=/home/gocias"

func TestVarPrefix(t *testing.T) {
	var mok EnvList = strings.Split(EnvListMok, ",")
	if len(mok.VarPrefix("HOME")) == 0 {
		t.Error("Prefix 'HOME' not found!")
	}
}

func TestVarContains(t *testing.T) {
	var mok EnvList = strings.Split(EnvListMok, ",")
	if len(mok.VarContains("bin")) == 0 {
		t.Error("Substring 'bin' not found!")
	}
}

func TestVarKeyContains(t *testing.T) {
	var mok EnvList = strings.Split(EnvListMok, ",")
	if len(mok.VarKeyContains("ERN")) == 0 {
		t.Error("Key with substring 'ERN' not found!")
	}
}

func TestVarValueContains(t *testing.T) {
	var mok EnvList = strings.Split(EnvListMok, ",")
	if len(mok.VarValueContains("goc")) == 0 {
		t.Error("Value with substring 'goc' not found!")
	}
}
