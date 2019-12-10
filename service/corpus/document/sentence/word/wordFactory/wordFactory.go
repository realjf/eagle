package wordFactory

import (
	. "eagle/service/common"
	"eagle/service/corpus/document/sentence/word"
	"eagle/service/corpus/document/sentence/word/compoundWord"
)

func Create(param string) word.IWord {
	if param == "" {
		return nil
	}
	if String(param).StartsWith("[") && !String(param).StartsWith("[/") {
		return compoundWord.Create(param)
	}else{
		return word.Create(param)
	}
}
