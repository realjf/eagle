package tag

import (
	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gutil"
	"strings"
)

var (
	GNature *Nature
)

func init() {
	if GNature == nil {
		GNature = NewNature("")
	}
}

// 词性
type Nature struct {
	Bg *Nature // 区别语素
	Mg *Nature // 数语素
	Nl *Nature // 名词性惯用语
	Nx *Nature // 字母专名
	Qg *Nature // 量词语素
	Ud *Nature // 助词
	Uj *Nature // 助词
	Uz *Nature // 着
	Ug *Nature // 过
	Ul *Nature // 连词
	Uv *Nature // 连词
	Yg *Nature // 语气语素
	Zg *Nature // 状态词
	// 以上标签来自ICT，以下标签来自北大
	N     *Nature // 名词
	Nr    *Nature // 人名
	Nrj   *Nature // 日语人名
	Nrf   *Nature // 音译人名
	Nr1   *Nature // 复姓
	Nr2   *Nature // 蒙古姓名
	Ns    *Nature // 地名
	Nsf   *Nature // 音译地名
	Nt    *Nature // 机构团体名
	Ntc   *Nature // 公司名
	Ntcf  *Nature // 工厂
	Ntcb  *Nature // 银行
	Ntch  *Nature // 酒店宾馆
	Nto   *Nature // 政府机构
	Ntu   *Nature // 大学
	Nts   *Nature // 中小学
	Nth   *Nature // 医院
	Nh    *Nature // 医药疾病等健康相关名词
	Nhm   *Nature // 药品
	Nhd   *Nature // 疾病
	Nn    *Nature // 工作相关名词
	Nnt   *Nature // 职务职称
	Nnd   *Nature // 职业
	Ng    *Nature // 名词性语素
	Nf    *Nature // 食品
	Ni    *Nature // 机构相关（不是独立机构名）
	Nit   *Nature // 教育相关机构
	Nic   *Nature // 下属机构
	Nis   *Nature // 机构后缀
	Nm    *Nature // 物品名
	Nmc   *Nature // 化学品名
	Nb    *Nature // 生物名
	Nba   *Nature // 动物名
	Nbc   *Nature // 动物纲目
	Nbp   *Nature // 植物名
	Nz    *Nature //  其他专名
	G     *Nature // 学术词汇
	Gm    *Nature // 数学相关词汇
	Gp    *Nature // 物理相关词汇
	Gc    *Nature // 化学相关词汇
	Gb    *Nature // 生物相关词汇
	Gbc   *Nature // 生物类别
	Gg    *Nature // 地理地址相关词汇
	Gi    *Nature // 计算机相关词汇
	J     *Nature // 简称略语
	I     *Nature // 成语
	L     *Nature // 习惯语
	T     *Nature // 时间词
	Tg    *Nature // 时间词性语素
	S     *Nature // 处所词
	F     *Nature // 方位词
	V     *Nature // 动词
	Vd    *Nature // 副动词
	Vn    *Nature // 名动词
	Vshi  *Nature // 动词 是
	Vyou  *Nature // 动词 有
	Vf    *Nature // 趋向安动词
	Vx    *Nature // 形式动词
	Vi    *Nature // 不及物动词（内动词）
	Vl    *Nature // 动词性惯用语
	Vg    *Nature // 动词性语素
	A     *Nature // 形容词
	Ad    *Nature // 副形词
	An    *Nature // 名形词
	Ag    *Nature // 形容词性语素
	Al    *Nature // 形容词性惯用语
	B     *Nature // 区别词
	Bl    *Nature // 区别词性惯用语
	Z     *Nature // 状态词
	R     *Nature // 代词
	Rr    *Nature // 人称代词
	Rz    *Nature // 指示代词
	Rzt   *Nature // 时间指示代词
	Rzs   *Nature // 处所指示代词
	Rzv   *Nature // 谓语性指示代词
	Ry    *Nature // 疑问代词
	Ryt   *Nature // 时间疑问代词
	Rys   *Nature // 处所疑问代词
	Ryv   *Nature // 谓词性疑问代词
	Rg    *Nature // 代词性语素
	Rg1   *Nature // 古汉语代词性语素
	M     *Nature // 数词
	Mq    *Nature // 数量词
	Mg1   *Nature // 甲乙丙丁之类的数词
	Q     *Nature // 量词
	Qv    *Nature // 动量词
	Qt    *Nature // 时量词
	D     *Nature // 副词
	Dg    *Nature // 辄,俱,复之类的副词
	Dl    *Nature // 连语
	P     *Nature // 介词
	Pba   *Nature // 介词 把
	Pbei  *Nature // 介词 被
	C     *Nature // 连词
	Cc    *Nature // 并列连词
	U     *Nature // 助词
	Uzhe  *Nature // 着
	Ule   *Nature // 了 喽
	Uguo  *Nature // 过
	Ude1  *Nature // 的 底
	Ude2  *Nature // 地
	Ude3  *Nature // 得
	Usuo  *Nature // 所
	Udeng *Nature // 等 等等 云云
	Uyy   *Nature // 一样 一般 似的 般
	Udh   *Nature // 的话
	Uls   *Nature // 来讲 来说 而言 说来
	Uzhi  *Nature // 之
	Ulian *Nature // 连
	E     *Nature // 叹词
	Y     *Nature // 语气词（delete yg）
	O     *Nature // 拟声词
	H     *Nature // 前缀
	K     *Nature // 后缀
	X     *Nature // 字符串
	Xx    *Nature // 非语素字
	Xu    *Nature // 网址URL
	W     *Nature // 标点符号
	Wkz   *Nature // 左括号，全角：（ 〔  ［  ｛  《 【  〖 〈   半角：( [ { <
	Wky   *Nature // 右括号，全角：） 〕  ］ ｝ 》  】 〗 〉 半角： ) ] { >
	Wyz   *Nature // 左引号，全角：“ ‘ 『
	Wyy   *Nature // 右引号，全角：” ’ 』
	Wj    *Nature // 句号，全角：。
	Ww    *Nature // 问号，全角：？ 半角：?
	Wt    *Nature // 叹号，全角：！ 半角：!
	Wd    *Nature // 逗号，全角：， 半角：,
	Wf    *Nature // 分号，全角：； 半角： ;
	Wn    *Nature // 顿号，全角：、
	Wm    *Nature // 冒号，全角：： 半角： :
	Ws    *Nature // 省略号，全角：……  …
	Wp    *Nature // 破折号，全角：——   －－   ——－   半角：---  ----
	Wb    *Nature // 百分号千分号，全角：％ ‰   半角：%
	Wh    *Nature // 单位符号，全角：￥ ＄ ￡  °  ℃  半角：$
	End   *Nature // 仅用于终##终，不会出现在分词结果中
	Begin *Nature // 仅用于始##始，不会出现在分词结果中

	idMap   *gmap.TreeMap
	values  []Nature
	ordinal int
	name    string
}

func newNature(name string) *Nature {
	n := Nature{
		name: name,
	}

	n.Init()

	return &n
}

func NewNature(name string) *Nature {
	nature := Nature{
		name: name,
		Bg:   newNature("bg"),
		Mg:   newNature("mg"),
		Nl:   newNature("nl"),
		Nx:   newNature("nx"),
		Qg:   newNature("qg"),
		Ud:   newNature("ud"),
		Uj:   newNature("uj"),
		Uz:   newNature("uz"),
		Ug:   newNature("ug"),
		Ul:   newNature("ul"),
		Uv:   newNature("uv"),
		Yg:   newNature("yg"),
		Zg:   newNature("zg"),

		N:     newNature("n"),
		Nr:    newNature("nr"),
		Nrj:   newNature("nrj"),
		Nr1:   newNature("nr1"),
		Nr2:   newNature("nr2"),
		Ns:    newNature("ns"),
		Nsf:   newNature("nsf"),
		Nt:    newNature("nt"),
		Ntc:   newNature("ntc"),
		Ntcf:  newNature("ntcf"),
		Ntcb:  newNature("ntcb"),
		Ntch:  newNature("ntch"),
		Nto:   newNature("nto"),
		Ntu:   newNature("ntu"),
		Nts:   newNature("nts"),
		Nth:   newNature("nth"),
		Nh:    newNature("nh"),
		Nhm:   newNature("nhm"),
		Nhd:   newNature("nhd"),
		Nn:    newNature("nn"),
		Nnt:   newNature("nnt"),
		Nnd:   newNature("nnd"),
		Ng:    newNature("ng"),
		Nf:    newNature("nf"),
		Ni:    newNature("ni"),
		Nit:   newNature("nit"),
		Nic:   newNature("nic"),
		Nis:   newNature("nis"),
		Nm:    newNature("nm"),
		Nmc:   newNature("nmc"),
		Nb:    newNature("nb"),
		Nba:   newNature("nba"),
		Nbc:   newNature("nbc"),
		Nbp:   newNature("nbp"),
		Nz:    newNature("nz"),
		G:     newNature("g"),
		Gm:    newNature("gm"),
		Gp:    newNature("gp"),
		Gc:    newNature("gc"),
		Gb:    newNature("gb"),
		Gbc:   newNature("gbc"),
		Gg:    newNature("gg"),
		Gi:    newNature("gi"),
		J:     newNature("j"),
		I:     newNature("i"),
		L:     newNature("l"),
		T:     newNature("t"),
		Tg:    newNature("tg"),
		S:     newNature("s"),
		F:     newNature("f"),
		V:     newNature("v"),
		Vd:    newNature("vd"),
		Vn:    newNature("vn"),
		Vshi:  newNature("vshi"),
		Vyou:  newNature("vyou"),
		Vf:    newNature("vf"),
		Vx:    newNature("vx"),
		Vi:    newNature("vi"),
		Vl:    newNature("vl"),
		Vg:    newNature("vg"),
		A:     newNature("a"),
		Ad:    newNature("ad"),
		An:    newNature("an"),
		Ag:    newNature("ag"),
		Al:    newNature("al"),
		B:     newNature("b"),
		Bl:    newNature("bl"),
		Z:     newNature("z"),
		R:     newNature("r"),
		Rr:    newNature("rr"),
		Rz:    newNature("rz"),
		Rzt:   newNature("rzt"),
		Rzs:   newNature("rzs"),
		Rzv:   newNature("rzv"),
		Ry:    newNature("ry"),
		Ryt:   newNature("ryt"),
		Rys:   newNature("rys"),
		Ryv:   newNature("ryv"),
		Rg:    newNature("rg"),
		Rg1:   newNature("rg1"),
		M:     newNature("m"),
		Mg1:   newNature("mg1"),
		Q:     newNature("q"),
		Qv:    newNature("qv"),
		Qt:    newNature("qt"),
		D:     newNature("d"),
		Dg:    newNature("dg"),
		Dl:    newNature("dl"),
		P:     newNature("p"),
		Pba:   newNature("pba"),
		Pbei:  newNature("pbei"),
		C:     newNature("c"),
		Cc:    newNature("cc"),
		U:     newNature("u"),
		Uzhe:  newNature("uzhe"),
		Ule:   newNature("ule"),
		Uguo:  newNature("uguo"),
		Ude1:  newNature("ude1"),
		Ude2:  newNature("ude2"),
		Ude3:  newNature("ude3"),
		Usuo:  newNature("usuo"),
		Udeng: newNature("udeng"),
		Uyy:   newNature("uyy"),
		Udh:   newNature("udh"),
		Uls:   newNature("uls"),
		Uzhi:  newNature("uzhi"),
		Ulian: newNature("ulian"),
		E:     newNature("e"),
		Y:     newNature("y"),
		O:     newNature("o"),
		H:     newNature("h"),
		K:     newNature("k"),
		X:     newNature("x"),
		Xx:    newNature("xx"),
		Xu:    newNature("xu"),
		W:     newNature("w"),
		Wkz:   newNature("wkz"),
		Wky:   newNature("wky"),
		Wyz:   newNature("wyz"),
		Wyy:   newNature("wyy"),
		Wj:    newNature("wj"),
		Ww:    newNature("ww"),
		Wt:    newNature("wt"),
		Wd:    newNature("wd"),
		Wf:    newNature("wf"),
		Wn:    newNature("wn"),
		Wm:    newNature("wm"),
		Ws:    newNature("ws"),
		Wp:    newNature("wp"),
		Wb:    newNature("wb"),
		Wh:    newNature("wh"),
		End:   newNature("end"),
		Begin: newNature("begin"),
	}

	nature.Init()

	return &nature
}

// 默认需要调用初始化
func (n *Nature) Init() {
	n.idMap = gmap.NewTreeMap(gutil.ComparatorInt, true)
	n.ordinal = n.idMap.Size()
	n.idMap.Set(n.name, n.ordinal)
	var extended []Nature = make([]Nature, n.idMap.Size())
	if n.values != nil {
		copy(extended, n.values)
	}
	extended[n.ordinal] = *n
	n.values = extended
}

// 词性是否以该前缀开头
// 词性根据开头的几个字母可以判断大的类别
func (n *Nature) StartsWith(prefix string) bool {
	return strings.HasPrefix(n.name, prefix)
}

func (n *Nature) StartsWithFirstChar(prefix string) bool {
	return strings.HasPrefix(n.name, prefix)
}

// 词性的首字母
// 词性根据开头的几个字母可以判断大的类别
func (n *Nature) FirstChar() string {
	return gstr.SubStr(n.name, 0, 1)
}

// 安全地将字符串类型的词性转为Enum类型，如果未定义该词性，则返回空
// @name 字符串词性
func (n *Nature) FromString(name string) Nature {
	var nature Nature
	if id, ok := n.idMap.Get(name).(int); ok {
		nature = n.values[id]
	}
	return nature
}

// 创建自定义词性,如果已有该对应词性,则直接返回已有的词性
func (n *Nature) Create(name string) Nature {
	var nature Nature = n.FromString(name)
	if (&nature) == nil {
		return *NewNature(name)
	}
	return nature
}

func (n *Nature) ToString() string {
	return n.name
}

func (n *Nature) Ordinal() int {
	return n.ordinal
}

func (n *Nature) Values() []Nature {
	return n.values
}
