package utils

import (
	"fmt"
	"testing"
)

func TestUnionMap(t *testing.T) {
	m1 := map[string]string{"a": "a", "b": "b", "c": "c"}
	m2 := map[string]string{"a": "a", "b": "b", "d": "d"}
	unionMap, _ := UnionMap(m1, m2)
	for k, v := range unionMap {
		fmt.Println(k, v)
	}
}

func TestDiffMap(t *testing.T) {
	type M map[string]interface{}
	m1 := M{"a": "a", "b": "b", "c": "c", "e": "e"}
	m2 := M{"a": "a", "b": "b", "d": "d", "e": "f"}
	justM1, justM2, diffM1AndM2, _ := DiffMap(M(m1), M(m2))
	fmt.Println("justM1")
	for k, v := range justM1 {
		fmt.Println(k, v)
	}

	fmt.Println("justM2")
	for k, v := range justM2 {
		fmt.Println(k, v)
	}

	fmt.Println("diffM1AndM2")
	for k, v := range diffM1AndM2 {
		fmt.Println(k, v)
	}
}
