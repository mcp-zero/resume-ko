/*
 * @Author: lucklidi@126.com
 * @Date: 2025-08-26 15:35:48
 * @LastEditTime: 2025-08-26 17:12:52
 * @LastEditors: lucklidi@126.com
 * @Description:
 * Copyright (c) 2025 by lucklidi, All Rights Reserved.
 */
package analyzer

// AnalysisResult 表示分析结果
type AnalysisResult struct {
	MatchedKeywords []string `json:"matched_keywords"`
	MissingKeywords []string `json:"missing_keywords"`
	MatchScore      int      `json:"match_score"`
}
