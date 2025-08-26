/*
 * @Author: lucklidi@126.com
 * @Date: 2025-08-26 15:34:28
 * @LastEditTime: 2025-08-26 16:12:24
 * @LastEditors: lucklidi@126.com
 * @Description:
 * Copyright (c) 2025 by lucklidi, All Rights Reserved.
 */
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sort"

	"resume_keyword_optimizer/analyzer"
	"resume_keyword_optimizer/web"
)

func main() {
	// 定义命令行参数
	resumeFile := flag.String("resume", "", "Path to resume file")
	jobDescriptionFile := flag.String("job", "", "Path to job description file")
	webPort := flag.String("web", "", "Port to start web server on")

	flag.Parse()

	// 如果提供了web参数，则启动Web服务器
	if *webPort != "" {
		server := web.NewServer(*webPort)
		log.Fatal(server.Start())
	}

	// 检查参数
	if *resumeFile == "" || *jobDescriptionFile == "" {
		fmt.Println("Usage: resume_optimizer -resume <resume_file> -job <job_description_file>")
		fmt.Println("   or: resume_optimizer -web <port>")
		fmt.Println("Examples:")
		fmt.Println("  resume_optimizer -resume my_resume.txt -job job_description.txt")
		fmt.Println("  resume_optimizer -web 8080")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// 读取文件内容
	resumeContent, err := os.ReadFile(*resumeFile)
	if err != nil {
		log.Fatalf("Error reading resume file: %v", err)
	}

	jobDescriptionContent, err := os.ReadFile(*jobDescriptionFile)
	if err != nil {
		log.Fatalf("Error reading job description file: %v", err)
	}

	// 分析关键字
	result, err := analyzer.Analyze(string(resumeContent), string(jobDescriptionContent))
	if err != nil {
		log.Fatalf("Error analyzing keywords: %v", err)
	}

	// 输出结果
	// 按字母顺序排序关键字
	sort.Strings(result.MissingKeywords)
	sort.Strings(result.MatchedKeywords)
	
	fmt.Println("=================== MATCHED KEYWORDS ===================")
	for _, keyword := range result.MatchedKeywords {
		fmt.Printf("✅ %s\n", keyword)
	}
	fmt.Println("========================================================")
	
	fmt.Println("\n=================== MISSING KEYWORDS ===================")
	for _, keyword := range result.MissingKeywords {
		fmt.Printf("❌ %s\n", keyword)
	}
	fmt.Println("========================================================")
	
	fmt.Printf("\nMatched Keywords: %d\n", len(result.MatchedKeywords))
	fmt.Printf("Missing Keywords: %d\n", len(result.MissingKeywords))
	fmt.Println("========================================================")
}
