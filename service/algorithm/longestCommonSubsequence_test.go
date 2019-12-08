package algorithm

import (
	. "eagle/service/common"
	"testing"
)

func TestLongestCommonSubsequence_ComputeString(t *testing.T) {
	str1 := String("aejrierojewor")
	str2 := String("jrierooafew")
	t.Fatal((&LongestCommonSubsequence{}).ComputeString(str2, str1))
}

