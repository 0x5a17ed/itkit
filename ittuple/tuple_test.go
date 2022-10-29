package ittuple_test

import (
	"testing"

	assertpkg "github.com/stretchr/testify/assert"

	"github.com/0x5a17ed/itkit/ittuple"
)

func TestTuple2_Len(t *testing.T) {
	tup := ittuple.NewT2("Left", "Right")
	assertpkg.Equalf(t, 2, tup.Len(), "Len()")
}

func TestTuple2_Values(t *testing.T) {
	tup := ittuple.NewT2("Left", "Right")
	got1, got2 := tup.Values()
	assertpkg.Equalf(t, "Left", got1, "Values()")
	assertpkg.Equalf(t, "Right", got2, "Values()")
}

func TestTuple2_Array(t *testing.T) {
	tup := ittuple.NewT2("Left", "Right")
	assertpkg.Equalf(t, [2]any{"Left", "Right"}, tup.Array(), "Array()")
}

func TestTuple2_Slice(t *testing.T) {
	tup := ittuple.NewT2("Left", "Right")
	assertpkg.Equalf(t, []any{"Left", "Right"}, tup.Slice(), "Slice()")
}

func TestTuple2_String(t *testing.T) {
	tup := ittuple.NewT2("Left", "Right")
	assertpkg.Equalf(t, `["Left" "Right"]`, tup.String(), "String()")
}
