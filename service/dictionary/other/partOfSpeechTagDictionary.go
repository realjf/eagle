package other

import (
	"eagle/service/config"
	"github.com/gogf/gf/container/gmap"
)

var (
	GPartOfSpeechTagDictionary *PartOfSpeechTagDictionary
)

func init() {
	if GPartOfSpeechTagDictionary == nil {
		GPartOfSpeechTagDictionary = NewPartOfSpeechTagDictionary()
	}
	GPartOfSpeechTagDictionary.Load(config.GConfig.PartOfSpeechTagDictionary)
}

// 词性标注集中英映射表
type PartOfSpeechTagDictionary struct {
	Translator gmap.TreeMap // 词性映射表

}

func NewPartOfSpeechTagDictionary() *PartOfSpeechTagDictionary {
	return &PartOfSpeechTagDictionary{}
}

func (p *PartOfSpeechTagDictionary) Load(path string) {

}

// 翻译词性
func (p *PartOfSpeechTagDictionary) Translate(tag string) string {
	cn := p.Translator.Get(tag)
	if cn.(string) == "" {
		return tag
	}else{
		return cn.(string)
	}
}
