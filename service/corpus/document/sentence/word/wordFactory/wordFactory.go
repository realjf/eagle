package wordFactory

import (
	. "eagle/service/common"
	"eagle/service/corpus/document/sentence/word"
)

func Create(param string) word.IWord {
	if param == "" {
		return nil
	}
	if String(param).StartsWith("[") && !String(param).StartsWith("[/") {
		return
	}else{
		return word.Create(param)
	}
}
