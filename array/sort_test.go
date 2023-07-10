package array

import (
	"fmt"
	"testing"

	"github.com/go-zoox/core-utils/strings"
)

func TestSort(t *testing.T) {
	a := []int{4, 3, 8, 9, 10, 1, 2, 6, 7, 5}

	b := Sort(a, func(v1, v2 int) bool {
		return v1 < v2
	})
	bArr := []string{}
	for _, v := range b {
		bArr = append(bArr, fmt.Sprintf("%d", v))
	}
	if strings.Join(bArr, ",") != "1,2,3,4,5,6,7,8,9,10" {
		t.Error("Expected 1,2,3,4,5,6,7,8,9,10, got", strings.Join(bArr, ","))
	}

	c := Sort(a, func(v1, v2 int) bool {
		return v1 > v2
	})
	cArr := []string{}
	for _, v := range c {
		cArr = append(cArr, fmt.Sprintf("%d", v))
	}
	if strings.Join(cArr, ",") != "10,9,8,7,6,5,4,3,2,1" {
		t.Error("Expected 10,9,8,7,6,5,4,3,2,1, got", strings.Join(cArr, ","))
	}

	aArr := []string{}
	for _, v := range a {
		aArr = append(aArr, fmt.Sprintf("%d", v))
	}
	if strings.Join(aArr, ",") != "4,3,8,9,10,1,2,6,7,5" {
		t.Error("Expected 4,3,8,9,10,1,2,6,7,5, got", strings.Join(aArr, ","))
	}
}
