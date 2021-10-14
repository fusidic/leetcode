package leetcode

type SummaryRanges struct {
	nums      []int
	intervals [][]int
}

func Constructor() SummaryRanges {
	return SummaryRanges{
		nums:      []int{-1 << 16, 1 << 16},
		intervals: [][]int{{-1 << 16, -1 << 16}, {1 << 16, 1 << 16}},
	}
}

func (s *SummaryRanges) AddNum(val int) {
	for i := 0; i < len(s.nums)-1; i++ {
		if val == s.nums[i] {
			return
		}

		if val > s.nums[i] && val < s.nums[i+1] {

			// intervals
			if abs(val-s.nums[i]) == 1 && abs(val-s.nums[i+1]) == 1 {
				idx1 := s.getIntervalsIndex(val - 1)
				idx2 := s.getIntervalsIndex(val + 1)
				// merge
				s.intervals[idx1][1] = s.intervals[idx2][1]
				s.intervals = append(s.intervals[:idx1], s.intervals[idx2+1:]...)
			} else if abs(val-s.nums[i]) == 1 {
				idx1 := s.getIntervalsIndex(val - 1)
				s.intervals[idx1][1] = val
			} else if abs(val-s.nums[i+1]) == 1 {
				idx2 := s.getIntervalsIndex(val + 1)
				s.intervals[idx2][0] = val
			} else {
				idx1 := s.getIntervalsIndex(s.nums[i])
				tmp1 := append([][]int{{val, val}}, s.intervals[:idx1+1]...)
				s.intervals = append(s.intervals[:idx1+1], tmp1...)
			}

			tmp2 := append([]int{val}, s.nums[i+1:]...)
			s.nums = append(s.nums[:i+1], tmp2...)
		}
	}
	return
}

func (s SummaryRanges) getIntervalsIndex(val int) int {
	for i, interval := range s.intervals {
		if interval[0] <= val && interval[1] >= val {
			return i
		}
	}
	return -1
}

// GetIntervals ...
func (s *SummaryRanges) GetIntervals() [][]int {
	return s.intervals[1 : len(s.intervals)-1]
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

/**
 * Your SummaryRanges object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(val);
 * param_2 := obj.GetIntervals();
 */
