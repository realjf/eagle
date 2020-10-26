package charType

import (
	"eagle/service/config"
	TextUtility "eagle/service/utility/textUtility"
	"eagle/utils"
	"errors"
	. "eagle/service/common"
	"eagle/service/corpus/io"
	"eagle/service/corpus/io/ByteArray"
	"github.com/gogf/gf/container/glist"
	"github.com/gogf/gf/util/gconv"
	"time"
)

// 字符类型

var (
	GCharType *CharType
)

func init()  {
	if GCharType == nil {
		GCharType = NewCharType()
	}
}

type CharType struct {
	CT_SINGLE Char // 单字节
	CT_DELIMITER Char // 分隔符"!,.?()[]{}+=
	CT_CHINESE Char // 中文字符
	CT_LETTER Char // 字母
	CT_NUM Char // 数字
	CT_INDEX Char // 序号
	CT_CNUM Char // 中文数字
	CT_OTHER Char // 其他

	Type []Char
}

func NewCharType() *CharType {
	charType := &CharType{
		CT_SINGLE: TextUtility.CT_SINGLE,
		CT_DELIMITER:TextUtility.CT_DELIMITER,
		CT_CHINESE:TextUtility.CT_CHINESE,
		CT_LETTER:TextUtility.CT_LETTER,
		CT_NUM:TextUtility.CT_NUM,
		CT_INDEX:TextUtility.CT_INDEX,
		CT_CNUM:TextUtility.CT_CNUM,
		CT_OTHER:TextUtility.CT_OTHER,
	}

	charType.Type = make([]Char, 65536)
	utils.Logger.Info("字符类型对应表开始加载 " + config.GConfig.CharTypePath)
	start := time.Now().UnixNano()
	byteArray := ByteArray.CreateByteArray(config.GConfig.CharTypePath)
	if byteArray == nil {
		var err error
		byteArray, err = charType.Generate()
		if err != nil {
			utils.Logger.Warning("字符类型对应表 " + config.GConfig.CharTypePath + " 加载失败：" + err.Error())
			return nil
		}
	}

	for byteArray.HasMore() {
		b := byteArray.NextChar()
		e := byteArray.NextChar()
		t := byteArray.NextByte()
		for i := b; i <= e; i++ {
			charType.Type[i] = Char(gconv.Rune(t))
		}
	}

	utils.Logger.Info("字符类型对应表加载成功，耗时", time.Now().UnixNano() - start, " ms")

	return charType
}

//
func (ct *CharType) Generate() (*ByteArray.ByteArray, error) {
	preType := 5
	preChar := 0
	typeList := glist.New(true)
	for i := 0; i <= Character_MAX_VALUE; i++ {
		t := TextUtility.CharType(Char(i))
		if t != preType {
			array := make([]int, 3)
			array[0] = preChar
			array[1] = i - 1
			array[2] = preType
			typeList.PushBack(array)
			preChar = i
		}
		preType = t
	}
	{
		array := make([]int, 3)
		array[0] = preChar
		array[1] = Character_MAX_VALUE
		array[2] = preType
		typeList.PushBack(array)
	}
	out := io.NewDataOutputStream(config.GConfig.CharTypePath)
	if out == nil {
		return nil, errors.New("open file error: " + config.GConfig.CharTypePath)
	}
	defer out.Close()

	for _, array := range typeList.FrontAll() {
		arr := array.([]int)
		out.WriteChar(arr[0])
		out.WriteChar(arr[1])
		out.WriteChar(arr[2])
	}

	bArray := ByteArray.CreateByteArray(config.GConfig.CharTypePath)
	return bArray, nil
}

// 获取字符的类型
func (ct *CharType) Get(c Char) Char {
	return ct.Type[c]
}

// 设置字符类型
// c 字符
// t 类型
func (ct *CharType) Set(c Char, t Char) {
	ct.Type[c] = t
}


