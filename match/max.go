package match

import (
	"fmt"
	"unicode/utf8"
)

type Max struct {
	Limit int
}

func (self Max) Match(s string) bool {
	var l int
	for range s {
		l += 1
		if l > self.Limit {
			return false
		}
	}

	return true
}

func (self Max) Index(s string) (int, []int) {
	if !self.Match(s) {
		return -1, nil
	}

	segments := make([]int, 0, self.Limit+1)
	segments = append(segments, 0)
	var count int
	for i, r := range s {
		count++
		if count > self.Limit {
			break
		}
		segments = append(segments, i+utf8.RuneLen(r))
	}

	return 0, segments
}

func (self Max) Len() int {
	return lenNo
}

func (self Max) Search(s string) (int, int, bool) {
	return 0, 0, false
}

func (self Max) Kind() Kind {
	return KindMax
}

func (self Max) String() string {
	return fmt.Sprintf("<max:%d>", self.Limit)
}
