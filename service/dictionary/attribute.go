package dictionary

import (
	"eagle/service/corpus/io"
	ByteArray "eagle/service/corpus/io/byteArray"
	. "eagle/service/corpus/tag"
	"eagle/utils"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
	"strings"
)

type Attribute struct {
	Nature []Nature
	Frequency []int
	TotalFrequency int
}

func NewAttribute(size int) *Attribute {
	return &Attribute{
		Nature: make([]Nature, 0, size),
		Frequency: make([]int, 0, size),
	}
}

func NewAttribute2(nature []Nature, frequency []int) *Attribute {
	return &Attribute{
		Nature:nature,
		Frequency:frequency,
	}
}

func NewAttribute3(nature Nature, frequency int) *Attribute {
	a := NewAttribute(1)
	a.Nature[0] = nature
	a.Frequency[0] = frequency
	a.TotalFrequency = frequency
	return a
}

func NewAttribute4(nature []Nature, frequency []int, totalFrequency int) *Attribute {
	return &Attribute{
		Nature:nature,
		Frequency:frequency,
		TotalFrequency:totalFrequency,
	}
}

/**
 * 使用单个词性，默认词频1000构造
 *
 * @param nature
 */
func NewAttribute5(nature Nature) *Attribute {
	return NewAttribute3(nature, 1000)
}


func (a *Attribute) Create(natureWithFrequency string) Attribute {
	var param []string = strings.Split(natureWithFrequency, " ")
	if len(param) % 2 != 0 {
		nature := NewNature("").Create(gstr.TrimStr(natureWithFrequency, " "))
		if &nature == nil {
			utils.Logger.Warning("使用字符串" + natureWithFrequency + "创建词条属性失败！")
			return Attribute{}
		}
		return *NewAttribute3(nature, 1)
	}
	natureCount := len(param) / 2
	attribute := NewAttribute(natureCount)
	for i := 0; i < natureCount; i++ {
		s := param[2*i]
		attribute.Nature[i] = NewNature("").Create(s)
		attribute.Frequency[i] = gconv.Int(param[1 + 2 * i])
		attribute.TotalFrequency += attribute.Frequency[i]
	}
	return *attribute
}

// 从字节流加载
func (a *Attribute) CreateFromByteArray(byteArray ByteArray.ByteArray, natureIndexArray []Nature) Attribute {
	currentTotalFrequency := byteArray.NextInt()
	length := byteArray.NextInt()
	attribute := NewAttribute(length)
	attribute.TotalFrequency = currentTotalFrequency
	for j := 0; j < length; j++ {
		attribute.Nature[j] = natureIndexArray[byteArray.NextInt()]
		attribute.Frequency[j] = byteArray.NextInt()
	}
	return *attribute
}

/**
 * 获取词性的词频
 *
 * @param nature 字符串词性
 * @return 词频
 * @deprecated 推荐使用Nature参数！
 */
func (a *Attribute) GetNatureFrequency(nature string) int {
	pos := NewNature("").Create(nature)
	if &pos == nil {
		return 0
	}
	return a.GetNatureFrequencyFromNature(pos)
}

/**
 * 获取词性的词频
 *
 * @param nature 词性
 * @return 词频
 */
func (a *Attribute) GetNatureFrequencyFromNature(nature Nature) int {
	i := 0
	for _, pos := range a.Nature {
		if pos.ToString() == nature.ToString() {
			return a.Frequency[i]
		}
		i++
	}
	return 0
}

/**
 * 是否有某个词性
 * @param nature
 * @return
 */
func (a *Attribute) HasNature(nature Nature) bool {
	return a.GetNatureFrequencyFromNature(nature) > 0
}

/**
 * 是否有以某个前缀开头的词性
 * @param prefix 词性前缀，比如u会查询是否有ude, uzhe等等
 * @return
 */
func (a *Attribute) HasNatureStartsWith(prefix string) bool {
	for _, n := range a.Nature {
		if n.StartsWith(prefix) {
			return true
		}
	}
	return false
}

func (a *Attribute) ToString() string {
	sb := ""
	for i :=0; i< len(a.Nature); i++ {
		sb += a.Nature[i].ToString()
		sb += " "
		sb += string(a.Frequency[i])
		sb += " "
	}
	return sb
}


func (a *Attribute) Save(out *io.DataOutputStream) {
	out.WriteInt(a.TotalFrequency)
	out.WriteInt(len(a.Nature))
	for i := 0; i< len(a.Nature); i++ {
		out.WriteInt(a.Nature[i].Ordinal())
		out.WriteInt(a.Frequency[i])
	}
}
