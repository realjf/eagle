package common

import (
	"bufio"
	"eagle/service/segment"
	"fmt"
	"github.com/gogf/gf/container/gset"
	"github.com/gogf/gf/util/gconv"
	"os"
	"strings"
)

type CWSEvaluator struct {
	A_size int
	B_size int
	A_cap_B_size int
	OOV int
	OOV_R int
	IV int
	IV_R int

	dic gset.StrSet
}

func NewCWSEvaluator2(dic gset.StrSet) *CWSEvaluator {
	return &CWSEvaluator{
		dic: dic,
	}
}

func NewCWSEvaluator(dictPath string) *CWSEvaluator {
	cwse := NewCWSEvaluator2(*gset.NewStrSet(true))
	if dictPath == "" {
		return nil
	}

	// 加载

	return cwse
}

// 获取PRF
func (c *CWSEvaluator) GetResult(percentage bool) Result {
	p := gconv.Float32(c.A_cap_B_size) / gconv.Float32(c.B_size)
	r := gconv.Float32(c.A_cap_B_size) / gconv.Float32(c.A_size)
	if percentage {
		p *= 100
		r *= 100
	}
	var oov_r float32
	if c.OOV > 0 {
		oov_r = gconv.Float32(c.OOV_R) / gconv.Float32(c.OOV)
		if percentage {
			oov_r *= 100
		}
	}
	var iv_r float32
	if c.IV > 0 {
		iv_r = gconv.Float32(c.IV_R) / gconv.Float32(c.IV)
		if percentage {
			iv_r *= 100
		}
	}
	return *NewResult(p, r, 2 * p * r / (p+r), oov_r, iv_r)
}


func (c *CWSEvaluator) GetResult2() Result {
	return c.GetResult(true)
}

// 比较标准答案与分词结果
func (c *CWSEvaluator) Compare(gold string, pred string) {
	var wordArray []string = strings.Split(gold, "\\s+")
	c.A_size += len(wordArray)
	var predArray []string = strings.Split(pred, "\\s+")
	c.B_size += len(predArray)

	goldIndex := 0
	predIndex := 0
	goldLen := 0
	predLen := 0

	for goldIndex < len(wordArray) && predIndex < len(predArray) {
		if goldLen == predLen {
			if wordArray[goldIndex] == predArray[predIndex] {
				if c.dic.Size() != 0 {
					if c.dic.Contains(wordArray[goldIndex]) {
						c.IV_R += 1
					} else {
						c.OOV_R += 1
					}
				}
				c.A_cap_B_size++
				goldLen += len(wordArray[goldIndex])
				predLen += len(wordArray[goldIndex])
				goldIndex++
				predIndex++
			}else{
				goldLen += len(wordArray[goldIndex])
				predLen += len(wordArray[predIndex])
				goldIndex++
				predIndex++
			}
		}else if goldLen < predLen {
			goldLen += len(wordArray[goldIndex])
			goldIndex++
		}else{
			predLen += len(predArray[predIndex])
			predIndex++
		}
	}

	if c.dic.Size() != 0 {
		for _, word := range wordArray {
			if c.dic.Contains(word) {
				c.IV += 1
			}else{
				c.OOV += 1
			}
		}
	}
}

// 在标准答案与分词结果上执行评测
//
func (c *CWSEvaluator) Evaluate(goldFile string, predFile string) Result {
	return c.Evaluate4(goldFile, predFile, "")
}

// 标准化评测分词器
// segment    分词器
// outputPath 分词预测输出文件
// goldFile   测试集segmented file
// dictPath   训练集单词列表
func (c *CWSEvaluator) Evaluate2(segment segment.Segment, outputPath string, goldFile string, dictPath string) Result {
	goldr, _ := os.Open(goldFile)
	defer goldr.Close()
	outw, _ := os.OpenFile(outputPath, os.O_RDWR, 600)
	defer outw.Close()
	buf := bufio.NewWriter(outw)


	result := c.Evaluate4(goldFile, outputPath, dictPath)
}

// 标准化评测分词器
// segment    分词器
// testFile   测试集raw text
// outputPath 分词预测输出文件
// goldFile   测试集segmented file
// dictPath   训练集单词列表
func (c *CWSEvaluator) Evaluate3(segment segment.Segment, testFile string, outputPath string, goldFile string, dictPath string) Result {
	return c.Evaluate2(segment, outputPath, goldFile, dictPath)
}

// 在标准答案与分词结果上执行评测
func (c *CWSEvaluator) Evaluate4(goldFile string, predFile string, dictPath string) Result {
	evaluator := NewCWSEvaluator(dictPath)


	return evaluator.GetResult2()
}

type Result struct {
	P float32
	R float32
	F1 float32
	OOV_R float32
	IV_R float32
}

func NewResult(p float32, r float32, f1 float32, OOV_R float32, IV_R float32) *Result {
	return &Result{
		P: p,
		R: r,
		F1: f1,
		OOV_R: OOV_R,
		IV_R: IV_R,
	}
}

func (r *Result) ToString() string {
	return fmt.Sprintf("P:%.2f R:%.2f F1:%.2f OOV-R:%.2f IV-R:%.2f", r.P, r.R, r.F1, r.OOV_R, r.IV_R)
}


