package sentencesUtil

import (
	. "eagle/service/common"
	"github.com/gogf/gf/container/glist"
	"strings"
)

// 文本断句

/**
 * 将文本切割为最细小的句子（逗号也视作分隔符）
 *
 * @param content
 * @return
 */
func ToSentenceList(content string) glist.List {
	return ToSentenceList4(String(content).ToCharArray(), true)
}

/**
 * 文本分句
 *
 * @param content  文本
 * @param shortest 是否切割为最细的单位（将逗号也视作分隔符）
 * @return
 */
func ToSentenceList2(content string, shortest bool) glist.List {
	return ToSentenceList4(String(content).ToCharArray(), shortest)
}

func ToSentenceList3(chars []Char) glist.List {
	return ToSentenceList4(chars, true)
}

func ToSentenceList4(chars []Char, shortest bool) glist.List {
	var sb String
	sentences := glist.New(true)
	for i := 0; i < len(chars); i++ {
		if sb.Length() == 0 && (chars[i].IsWhitespace() || chars[i] == ' ') {
			continue
		}

		sb.Append(chars[i])
		switch chars[i] {
		case '.':
			if i < len(chars) - 1 && chars[i+1].ToInt() > 128 {
				InsertIntoList(sb, sentences)
				sb = NewString()
			}
		case '…':
			if i < len(chars) - 1 && chars[i+1] == '…' {
				sb.Append('…')
				i++
				InsertIntoList(sb, sentences)
				sb = NewString()
			}
		case '，': fallthrough
		case ',': fallthrough
		case ';': fallthrough
		case '；':
			if !shortest {
				continue
			}
		case ' ': fallthrough
		case '	': fallthrough
		case ' ': fallthrough
		case '。': fallthrough
		case '!': fallthrough
		case '！': fallthrough
		case '?': fallthrough
		case '？': fallthrough
		case '\n': fallthrough
		case '\r':
			InsertIntoList(sb, sentences)
			sb = NewString()
		}
	}

	if sb.Length() > 0 {
		InsertIntoList(sb, sentences)
	}

	return *sentences
}

func InsertIntoList(sb String, sentences *glist.List) {
	content := strings.Trim(sb.ToString(), " ")
	if String(content).Length() > 0 {
		sentences.PushBack(content)
	}
}


