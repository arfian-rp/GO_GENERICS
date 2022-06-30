package utils

import "fmt"

type any interface{}

func Length[T any](param T) T { //di golang harus pakai constraint => aturan yang digunakan untuk menentukan tipe data yang diperbolehkan pada Type Parameter
	fmt.Println(param)
	return param
}

func MultipleParameter[T1 any, T2 any](p1 T1, p2 T2) {
	fmt.Println(p1)
	fmt.Println(p2)
}

func IsSame[T comparable](v1, v2 T) bool {
	return v1 == v2
}

type Number interface {
	int | int8 | int16 | int32 | int64 |
		float32 | float64
}

func Min[T Number](first T, last T) T {
	if first < last {
		return first
	}
	return last
}
