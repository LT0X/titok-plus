package test

import (
	"errors"
	"fmt"
	"testing"
)

func TestCode(t *testing.T) {

	xx := errors.New("dsfsdfs")
	fmt.Print(xx)

}

func funcTest() func(string) string {
	return func(s string) string {
		return s
	}
}

func judge(mid int, num int, x []int, m int) bool {

	now := 0
	next := 0
	for next <= len(x) {
		if next == len(x) {
			if (m - x[next-1]) < mid {
				num--
			}
			break
		}
		if x[next]-now < mid {
			num--
			next++
		} else {
			now = x[next]
			next++
		}
	}
	if num < 0 {
		return false
	}
	return true

}
