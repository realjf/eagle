package algorithm

import (
	. "gifs/service/common"
	"testing"
)

func TestLongestCommonSubstring_ComputeString(t *testing.T) {
	str1 := String("aejrierojewor")
	str2 := String("jrierooafew")
	t.Fatal((&LongestCommonSubstring{}).ComputeString(str2, str1))
}


func Benchmark_LongestCommonSubstring_ComputeString(b *testing.B) {
	str1 := String("aejrierojewor")
	str2 := String("jrierooafew")
	for i := 0; i < b.N; i++ {
		(&LongestCommonSubstring{}).ComputeString(str2, str1)
	}
}

