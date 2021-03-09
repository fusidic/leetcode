package dp

import (
	"fmt"
	"testing"
)

func Test_LCS(t *testing.T) {
	str1 := "2LQ74WK8Ld0x7d8FP8l61pD7Wsz1E9xOMp920hM948eGjL9Kb5KJt80"
	str2 := "U08U29zzuodz16CBZ8xfpmmn5SKD80smJbK83F2T37JRqYfE76vh6hrE451uFQ100ye9hog1Y52LDk0L52SuD948eGjLz0htzd5YF9J1Y6oI7562z4T2"
	fmt.Println(LCS(str1, str2))
}
