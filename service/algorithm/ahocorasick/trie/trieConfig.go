package trie

// 配置
type TrieConfig struct {
	allowOverlaps bool // 允许重叠
	RemainLongest bool // 只保留最长匹配
}

// 是否允许重叠
func (tc *TrieConfig) IsAllowOverlaps() bool {
	return tc.allowOverlaps
}

// 设置是否允许重叠
func (tc *TrieConfig) SetAllowOverlaps(allowOverlaps bool) {
	tc.allowOverlaps = allowOverlaps
}
