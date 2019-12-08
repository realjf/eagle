package config


var (
	GConfig *Config
)

func init() {
	if GConfig == nil {
		GConfig = NewConfig()
	}
}

// 库的全局配置，既可以用代码修改，也可以通过hanlp.properties配置（按照 变量名=值 的形式）
type Config struct {
	DEBUG                                       bool     // 开发模式
	CoreDictionaryPath                          string   // 核心词典路径
	CoreDictionaryTransformMatrixDictionaryPath string   // 核心词典词性转移矩阵路径
	CustomDictionaryPath                        []string // 用户自定义词典路径
	BiGramDictionaryPath                        string   // 2元语法词典路径
	CoreStopWordDictionaryPath                  string   // 停用词词典路径
	CoreSynonymDictionaryDictionaryPath         string   // 同义词词典路径
	PersonDictionaryPath                        string   // 人名词典路径
	PersonDictionaryTrPath                      string   // 人名词典转移矩阵路径
	PlaceDictionaryPath                         string   // 地名词典路径
	PlaceDictionaryTrPath                       string   // 地名词典转移矩阵路径
	OrganizationDictionaryPath                  string   // 地名词典路径
	OrganizationDictionaryTrPath                string   // 地名词典转移矩阵路径
	TcDictionaryRoot                            string   // 简繁转换词典根目录
	PinyinDictionaryPath                        string   // 拼音词典路径
	TranslatedPersonDictionaryPath              string   // 音译人名词典
	JapanesePersonDictionaryPath                string   // 日本人名词典路径
	CharTypePath                                string   // 字符类型对应表
	CharTablePath                               string   // 字符正规化表（全角转半角，繁体转简体）
	PartOfSpeechTagDictionary                   string   // 词性标注集描述表，用来进行中英映射（对于Nature词性，可直接参考Nature.java中的注释）
	WordNatureModelPath                         string   // 词-词性-依存关系模型
	MaxEntModelPath                             string   // 最大熵-依存关系模型（已废弃，请使用{@link KBeamArcEagerDependencyParser）
	NNParserModelPath                           string   // 神经网络依存模型路径
	PerceptronParserModelPath                   string   // 感知机ArcEager依存模型路径
	CRFSegmentModelPath                         string   // CRF分词模型（已废弃，请使用{@link com.hankcs.hanlp.model.crf.CRFLexicalAnalyzer}）
	HMMSegmentModelPath                         string   // HMM分词模型（已废弃，请使用{@link PerceptronLexicalAnalyzer}）
	CRFCWSModelPath                             string   // CRF分词模型
	CRFPOSModelPath                             string   // CRF词性标注模型
	CRFNERModelPath                             string   // CRF命名实体识别模型
	PerceptronCWSModelPath                      string   // 感知机分词模型
	PerceptronPOSModelPath                      string   // 感知机词性标注模型
	PerceptronNERModelPath                      string   // 感知机命名实体识别模型
	ShowTermNature                              bool     // 分词结果是否展示词性
	Normalization                               bool     // 是否执行字符正规化（繁体->简体，全角->半角，大写->小写），切换配置后必须删CustomDictionary.txt.bin缓存

	// IO适配器（默认null，表示从本地文件系统读取），实现com.hankcs.hanlp.corpus.io.IIOAdapter接口
	//IOAdapter io.IIOAdapter
}

func NewConfig() *Config {
	root := "../"
	return &Config{
		DEBUG:              false,
		CoreDictionaryPath: root + "data/dictionary/CoreNatureDictionary.txt",
		CoreDictionaryTransformMatrixDictionaryPath: root + "data/dictionary/CoreNatureDictionary.tr.txt",
		CustomDictionaryPath:                        []string{root + "data/dictionary/custom/CustomDictionary.txt"},
		BiGramDictionaryPath:                        root + "data/dictionary/CoreNatureDictionary.ngram.txt",
		CoreStopWordDictionaryPath:                  root + "data/dictionary/stopwords.txt",
		CoreSynonymDictionaryDictionaryPath:         root + "data/dictionary/synonym/CoreSynonym.txt",
		PersonDictionaryPath:                        root + "data/dictionary/person/nr.txt",
		PersonDictionaryTrPath:                      root + "data/dictionary/person/nr.tr.txt",
		PlaceDictionaryPath:                         root + "data/dictionary/place/ns.txt",
		PlaceDictionaryTrPath:                       root + "data/dictionary/place/ns.tr.txt",
		OrganizationDictionaryPath:                  root + "data/dictionary/organization/nt.txt",
		OrganizationDictionaryTrPath:                root + "data/dictionary/organization/nt.tr.txt",
		TcDictionaryRoot:                            root + "data/dictionary/tc/",
		PinyinDictionaryPath: root + "data/dictionary/pinyin/pinyin.txt",
		TranslatedPersonDictionaryPath: root + "data/dictionary/person/nrf.txt",
		JapanesePersonDictionaryPath: root + "data/dictionary/person/nrj.txt",
		CharTypePath: root + "data/dictionary/other/CharType.bin",
		CharTablePath: root + "data/dictionary/other/CharTable.txt",
		PartOfSpeechTagDictionary: root + "data/dictionary/other/TagPKU98.csv",
		WordNatureModelPath: root + "data/model/dependency/WordNature.txt",
		MaxEntModelPath: root + "data/model/dependency/MaxEntModel.txt",
		NNParserModelPath: root + "data/model/dependency/NNParserModel.txt",
		PerceptronParserModelPath: root + "data/model/dependency/perceptron.bin",
		CRFSegmentModelPath: root + "data/model/segment/CRFSegmentModel.txt",
		HMMSegmentModelPath: root + "data/model/segment/HMMSegmentModel.bin",
		CRFCWSModelPath: root + "data/model/crf/pku199801/cws.txt",
		CRFPOSModelPath: root + "data/model/crf/pku199801/pos.txt",
		CRFNERModelPath: root + "data/model/crf/pku199801/ner.txt",
		PerceptronCWSModelPath: root + "data/model/perceptron/large/cws.bin",
		PerceptronPOSModelPath: root + "data/model/perceptron/pku1998/pos.bin",
		PerceptronNERModelPath: root + "data/model/perceptron/pku1998/ner.bin",
		ShowTermNature: true,
		Normalization: false,
	}
}

// 开启调试模式(会降低性能)
func (c *Config) EnableDebug() {
	c.EnableDebug2(true)
}

// 开启调试模式(会降低性能)
func (c *Config) EnableDebug2(enable bool) {
	c.DEBUG = enable
	if c.DEBUG {

	} else {

	}
}
