package array

import (
	"sort"
	"strconv"
	"strings"
	"testing"
)

func TestUniqueBy(t *testing.T) {
	row := []int{1, 1, 2, 3, 4, 4, 5, 6, 7, 8, 8, 9, 10}
	unique := UniqueBy(row, func(value int, index int) int {
		return value
	})

	if len(unique) != 10 {
		t.Error("Expected 10, got", len(unique))
	}

	sort.Ints(unique)
	if strings.Join(Map(unique, func(value int, index int) string { return strconv.Itoa(value) }), ",") != "1,2,3,4,5,6,7,8,9,10" {
		t.Error("Expected 1,2,3,4,5,6,7,8,9,10, got", strings.Join(Map(unique, func(value int, index int) string { return strconv.Itoa(value) }), ","))
	}
}
