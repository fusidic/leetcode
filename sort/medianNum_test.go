package sort

import "testing"

func Test_getRow(t *testing.T) {
	mf := Constructor()
	mf.AddNum(1)
	mf.AddNum(2)
	mf.FindMedian()
	mf.AddNum(3)
	mf.FindMedian()
}
