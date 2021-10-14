package leetcode

import "testing"

func TestAddNum(t *testing.T) {
	s := Constructor()
	s.AddNum(1)
	s.AddNum(3)
	s.AddNum(7)
	s.AddNum(2)
	s.AddNum(6)
}
