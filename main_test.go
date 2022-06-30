package main

import (
	"generics/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTypeParameter(t *testing.T) {
	utils.Length[int32](1)
	utils.Length("Judika")
}

func TestMultipleParameter(t *testing.T) {
	utils.MultipleParameter[string, uint16]("hello", 12)
}

func TestIsSmae(t *testing.T) {
	assert.Equal(t, true, utils.IsSame("yudi", "yudi"))
	assert.Equal(t, false, utils.IsSame(121, 12))
}

func TestTypeParameterInheritance(t *testing.T) {
	assert.Equal(t, "Dewi", utils.GetName[utils.Manager](&utils.MyManager{Name: "Dewi"}))
	assert.Equal(t, "Devano", utils.GetName[utils.VicePresident](&utils.MyVicePresident{Name: "Devano"}))
}

func TestTypeSets(t *testing.T) {
	//menentukan lebih dari satu tipe constraint yang diperbolehkan pada type parameter

	assert.Equal(t, 12.3, utils.Min(12.3, 77.9))
}
