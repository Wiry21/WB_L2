package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	dictionary := []string{"Пятак", "пятак", "пятка", "тяпка", "листок", "слиток", "столик", "ветер"}
	res := AnagramMap(dictionary)
	fmt.Println(res)
	expected := map[string][]string{"листок": {"листок", "слиток", "столик"}, "пятак": {"пятак", "пятка", "тяпка"}}
	fmt.Println(expected)
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("find anagram error")
	}
}
