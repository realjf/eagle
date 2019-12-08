package segment

// 分词器配置项
type Config struct {
	IndexMode int // 是否是索引分词（合理地最小分割），indexMode代表全切分词语的最小长度（包含）
	NameRecognize bool // 是否识别中国人名，默认true
	TranslateNameRecognize bool // 是否识别音译人名
	JapaneseNameRecognize bool // 是否识别日本人名
	PlaceRecognize bool // 是否识别地名
	OrganizationRecognize bool // 是否识别机构
	UseCustomDictionary bool // 是否加载用户词典
	ForceCustomDictionary bool // 设置用户词典高优先级, false
	SpeechTagging bool // 词性标注
	Ner bool // 命名实体识别是否至少有一项被激活
	Offset bool // 是否计算偏移量
	NumberQuantifierRecognize bool // 是否识别数字和量词
	ThreadNumber int // 并行分词的线程数, goroutine数目
}

func NewConfig() *Config {
	return &Config{
		IndexMode: 0,
		NameRecognize: true,
		TranslateNameRecognize: true,
		JapaneseNameRecognize: false,
		PlaceRecognize: false,
		OrganizationRecognize: false,
		UseCustomDictionary: true,
		ForceCustomDictionary: false,
		SpeechTagging: false,
		Ner: true,
		Offset: false,
		NumberQuantifierRecognize: false,
		ThreadNumber: 1,
	}
}

// 更新命名实体识别总开关
func (c *Config) UpdateNerConfig() {
	c.Ner = c.NameRecognize || c.TranslateNameRecognize || c.JapaneseNameRecognize || c.PlaceRecognize || c.OrganizationRecognize
}


// 是否是索引模式
func (c *Config) IsIndexMode() bool {
	return c.IndexMode > 0
}


