package match

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

type List struct {
	List string
	Not  bool
}

func (self List) Kind() Kind {
	return KindList
}

func (self List) Match(s string) bool {
	inList := strings.Index(self.List, s) != -1
	return inList == !self.Not
}

func (self List) Len() int {
	return lenOne
}

func (self List) Index(s string) (int, []int) {
	for i, r := range s {
		if self.Not == (strings.IndexRune(self.List, r) == -1) {
			return i, []int{utf8.RuneLen(r)}
		}
	}

	return -1, nil
}

func (self List) String() string {
	var not string
	if self.Not {
		not = "!"
	}

	return fmt.Sprintf("<list:%s[%s]>", not, self.List)
}
