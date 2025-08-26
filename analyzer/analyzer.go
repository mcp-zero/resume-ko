/*
 * @Author: lucklidi@126.com
 * @Date: 2025-08-26 15:36:07
 * @LastEditTime: 2025-08-26 17:21:59
 * @LastEditors: lucklidi@126.com
 * @Description:
 * Copyright (c) 2025 by lucklidi, All Rights Reserved.
 */
package analyzer

import (
	"strings"
	"unicode"

	"github.com/go-ego/gse"
)

// 全局gse分词器，避免重复加载字典
var seg gse.Segmenter

func init() {
	seg.LoadDict()
}

// Analyze 分析简历和职位描述，找出缺失的关键字
func Analyze(resume, jobDescription string) (*AnalysisResult, error) {
	// 提取关键字
	resumeKeywords := extractKeywords(resume)
	jobKeywords := extractKeywords(jobDescription)

	// 比较关键字
	matched, missing := compareKeywords(resumeKeywords, jobKeywords)

	// 计算匹配度
	matchScore := 0
	if len(jobKeywords) > 0 {
		matchScore = (len(matched) * 100) / len(jobKeywords)
	}

	return &AnalysisResult{
		MatchedKeywords: matched,
		MissingKeywords: missing,
		MatchScore:      matchScore,
	}, nil
}

// extractKeywords 从文本中提取关键字
// 定义停用词集合
var stopWords = map[string]bool{
	// 中文停用词
	"的": true, "了": true, "在": true, "是": true, "我": true, "有": true, "和": true,
	"就": true, "不": true, "人": true, "都": true, "一": true, "一个": true, "上": true,
	"也": true, "很": true, "到": true, "说": true, "要": true, "去": true, "你": true,
	"会": true, "着": true, "没有": true, "看": true, "好": true, "自己": true, "这": true,
	"那": true, "里": true, "就是": true, "还是": true, "为了": true, "我们": true, "可以": true,
	"需要": true, "要求": true, "一名": true, "优秀": true, "良好": true, "精神": true,

	// 英文停用词
	"the": true, "a": true, "an": true, "and": true, "or": true, "but": true, "in": true,
	"on": true, "at": true, "to": true, "for": true, "of": true, "with": true, "by": true,
	"is": true, "are": true, "was": true, "were": true, "be": true, "been": true, "have": true,
	"has": true, "had": true, "do": true, "does": true, "did": true, "will": true, "would": true,
	"could": true, "should": true, "may": true, "might": true, "must": true, "can": true,
	"i": true, "you": true, "he": true, "she": true, "it": true, "we": true, "they": true,
	"me": true, "him": true, "her": true, "us": true, "them": true,
}

// extractKeywords 从文本中提取关键字
func extractKeywords(text string) map[string]bool {
	keywords := make(map[string]bool)

	// 转换为小写
	text = strings.ToLower(text)

	// 使用全局gse分词器进行中英文混合分词
	// 使用搜索模式分词
	words := seg.CutSearch(text, true)

	// 处理分词结果
	for _, word := range words {
		// 清理标点符号
		cleanWord := strings.TrimFunc(word, func(r rune) bool {
			return !unicode.IsLetter(r) && !unicode.IsNumber(r)
		})

		// 过滤空字符串、停用词和无效词汇
		// 中文必须2个字或者以上，英文必须2个字母或者以上
		if len(cleanWord) > 0 && !stopWords[cleanWord] && isValidKeyword(cleanWord) {
			if isChinese(cleanWord) && countChineseChars(cleanWord) >= 2 {
				keywords[cleanWord] = true
			} else if !isChinese(cleanWord) && len(cleanWord) >= 2 {
				keywords[cleanWord] = true
			}
		}
	}

	return keywords
}

// isValidKeyword 检查词汇是否为有效的关键字
func isValidKeyword(word string) bool {
	// 过滤掉包含数字的词汇（除非是纯数字）
	hasLetter := false
	hasDigit := false
	for _, r := range word {
		if unicode.IsLetter(r) {
			hasLetter = true
		}
		if unicode.IsDigit(r) {
			hasDigit = true
		}
	}

	// 如果同时包含字母和数字，则认为是无效关键字
	if hasLetter && hasDigit {
		return false
	}

	// 过滤掉太长的词汇（超过10个字符）
	if len(word) > 10 {
		return false
	}

	return true
}

// countChineseChars 计算字符串中中文字符的数量
func countChineseChars(s string) int {
	count := 0
	for _, r := range s {
		if unicode.Is(unicode.Scripts["Han"], r) {
			count++
		}
	}
	return count
}

// isChinese 检查字符串是否包含中文字符
func isChinese(s string) bool {
	for _, r := range s {
		if unicode.Is(unicode.Scripts["Han"], r) {
			return true
		}
	}
	return false
}

// compareKeywords 比较两个关键字集合
func compareKeywords(resumeKeywords, jobKeywords map[string]bool) (matched, missing []string) {
	for keyword := range jobKeywords {
		// 检查关键字是否符合要求：中文必须2个字或者以上，英文必须2个字母或者以上
		if (isChinese(keyword) && len(keyword) >= 2) || (!isChinese(keyword) && len(keyword) >= 2) {
			if resumeKeywords[keyword] {
				matched = append(matched, keyword)
			} else {
				missing = append(missing, keyword)
			}
		}
	}

	return matched, missing
}
