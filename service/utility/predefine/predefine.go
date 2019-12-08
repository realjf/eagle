package predefine

import (
	"regexp"
)

// 一些预定义的静态全局变量
const (
	CHINESE_NUMBERS = "零○〇一二两三四五六七八九十廿百千万亿壹贰叁肆伍陆柒捌玖拾佰仟"

	// hanlp.properties的路径，一般情况下位于classpath目录中。
	HANLP_PROPERTIES_PATH = ""

	MIN_PROBABILITY = 1e-10

	POSTFIX_SINGLE = "坝邦堡城池村单岛道堤店洞渡队峰府冈港阁宫沟国海号河湖环集江礁角街井郡坑口矿里岭楼路门盟庙弄牌派坡铺旗桥区渠泉山省市水寺塔台滩坛堂厅亭屯湾屋溪峡县线乡巷洋窑营屿园苑院闸寨站镇州庄族陂庵町"

	SEPERATOR_C_SENTENCE = "。！？：；…"
	SEPERATOR_C_SUB_SENTENCE = "、，（）“”‘’"
	SEPERATOR_E_SENTENCE = "!?:;"
	SEPERATOR_E_SUB_SENTENCE = ",()*'"

	//注释：原来程序为",()\042'"，"\042"为10进制42好ASC字符，为*
	SEPERATOR_LINK = "\n\r 　"

	WORD_SEGMENTER = "@"

	MAX_SEGMENT_NUM = 10

	MAX_FREQUENCY = 25146057 // 现在总词频25146057


	// 平滑参数
	DSmoothingPara = 0.1
	// 地址 ns
	TAG_PLACE = "未##地"
	// 句子的开始 begin
	TAG_BIGIN = "始##始"
	// 其它
	TAG_OTHER = "未##它"
	// 团体名词 nt
	TAG_GROUP = "未##团"
	// 数词 m
	TAG_NUMBER = "未##数"
	// 数量词 mq （现在觉得应该和数词同等处理，比如一个人和一人都是合理的）
	TAG_QUANTIFIER = "未##量"
	// 专有名词 nx
	TAG_PROPER = "未##专"
	// 时间 t
	TAG_TIME = "未##时"
	// 字符串 x
	TAG_CLUSTER = "未##串"
	// 结束 end
	TAG_END = "末##末"
	// 人名 nr
	TAG_PEOPLE = "未##人"
	// trie树文件后缀名
	TRIE_EXT = ".trie.dat"
	// 值文件后缀名
	VALUE_EXT = ".value.dat"
	// 逆转后缀名
	REVERSE_EXT = ".reverse"
	// 二进制文件后缀
	BIN_EXT = ".bin"
)

var (
	// 浮点数正则
	PATTERN_FLOAT_NUMBER = regexp.MustCompile("^(-?\\d+)(\\.\\d+)?$")

	POSTFIX_MUTIPLE = []string{"半岛","草原","城区","大堤","大公国","大桥","地区",
		"帝国","渡槽","港口","高速公路","高原","公路","公园","共和国","谷地","广场",
		"国道","海峡","胡同","机场","集镇","教区","街道","口岸","码头","煤矿",
		"牧场","农场","盆地","平原","丘陵","群岛","沙漠","沙洲","山脉","山丘",
		"水库","隧道","特区","铁路","新村","雪峰","盐场","盐湖","渔场","直辖市",
		"自治区","自治县","自治州"}

	// Smoothing 平滑因子
	DTemp = float64(1.0) / MAX_FREQUENCY + 0.00001

	// 日志组件
)
