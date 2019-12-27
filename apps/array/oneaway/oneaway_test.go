package main

import (
	"testing"
)

const n = 1024

func TestOneDelete(t *testing.T) {
	str1 := "pale"
	str2 := "ple"
	if !oneAway(str1, str2) {
		t.Errorf("%v, %v, expected true", str1, str2)
	}
}

func TestOneInsert(t *testing.T) {
	str1 := "pale"
	str2 := "pales"
	if !oneAway(str1, str2) {
		t.Errorf("%v, %v, expected true", str1, str2)
	}
}

func TestOneEdit(t *testing.T) {
	str1 := "pale"
	str2 := "bale"
	if !oneAway(str1, str2) {
		t.Errorf("%v, %v, expected true", str1, str2)
	}
}

func TestMultiEdit(t *testing.T) {
	str1 := "pale"
	str2 := "bake"
	if oneAway(str1, str2) {
		t.Errorf("%v, %v, expected false", str1, str2)
	}
}
