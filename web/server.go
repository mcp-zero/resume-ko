package web

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"resume_keyword_optimizer/analyzer"
)

// Server represents the web server
type Server struct {
	port string
}

// NewServer creates a new web server
func NewServer(port string) *Server {
	return &Server{
		port: port,
	}
}

// Start starts the web server
func (s *Server) Start() error {
	// Serve static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static/"))))

	// Serve API endpoints
	http.HandleFunc("/api/analyze", s.handleAnalyze)

	// Serve the main page
	http.HandleFunc("/", s.handleMain)

	log.Printf("Server starting on port %s", s.port)
	return http.ListenAndServe(":"+s.port, nil)
}

// handleMain serves the main page
func (s *Server) handleMain(w http.ResponseWriter, r *http.Request) {
	// Simple template for the main page
	html := `
<!DOCTYPE html>
<html lang="zh-CN">
<head>
	   <meta charset="UTF-8">
	   <meta name="viewport" content="width=device-width, initial-scale=1.0">
	   <title>简历关键字优化工具</title>
	   <link rel="stylesheet" href="/static/style.css">
</head>
<body>
	   <div class="container">
	       <header>
	           <h1>简历关键字优化工具</h1>
	           <p>帮助您的简历通过ATS筛选系统</p>
	       </header>
	       
	       <main>
	           <div class="input-section">
	               <div class="input-group">
	                   <label for="resume">您的简历</label>
	                   <textarea id="resume" placeholder="在此粘贴您的简历内容..."></textarea>
	               </div>
	               
	               <div class="input-group">
	                   <label for="jobDescription">职位描述</label>
	                   <textarea id="jobDescription" placeholder="在此粘贴职位描述..."></textarea>
	               </div>
	               
	               <button id="analyzeBtn">分析关键字</button>
	           </div>
	           
	           <div class="results-section" id="results" style="display: none;">
	               <h2>分析结果</h2>
	               
	               <!-- 匹配度圆饼图 -->
	               <div class="match-score-section">
	                   <h3>简历匹配度</h3>
	                   <div class="chart-container">
	                       <canvas id="matchScoreChart" width="200" height="200"></canvas>
	                       <div class="chart-center">
	                           <span id="matchScoreValue">0</span>%
	                       </div>
	                   </div>
	               </div>
	               
	               <!-- 匹配的关键字 -->
	               <div class="result-item">
	                   <div class="result-header">
	                       <h3>匹配的关键字</h3>
	                       <span class="keyword-count" id="matchedCount">0</span>
	                   </div>
	                   <div class="keywords-container" id="matchedKeywordsContainer">
	                       <div class="keywords-list" id="matchedKeywords"></div>
	                       <button class="toggle-btn" id="toggleMatchedBtn">展开全部</button>
	                   </div>
	               </div>
	               
	               <!-- 缺失的关键字 -->
	               <div class="result-item">
	                   <div class="result-header">
	                       <h3>缺失的关键字</h3>
	                       <span class="keyword-count" id="missingCount">0</span>
	                   </div>
	                   <div class="keywords-container" id="missingKeywordsContainer">
	                       <div class="keywords-list" id="missingKeywords"></div>
	                       <button class="toggle-btn" id="toggleMissingBtn">展开全部</button>
	                   </div>
	               </div>
	               
	               <!-- 加油鼓励 -->
	               <div class="encouragement-section" id="encouragementSection" style="display: none;">
	                   <h3>我们相信，凡事发生必有利于你</h3>
	                   <p id="encouragementText"></p>
	               </div>
	           </div>
	       </main>
	       
	       <footer>
	           <p>© 2025 简历关键字优化工具</p>
	       </footer>
	   </div>
	   
	   <script src="/static/chart.umd.min.js"></script>
	   <script src="/static/script.js"></script>
</body>
</html>
`

	tmpl, err := template.New("main").Parse(html)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// handleAnalyze handles the keyword analysis API request
func (s *Server) handleAnalyze(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse form data
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	resume := r.FormValue("resume")
	jobDescription := r.FormValue("jobDescription")

	if resume == "" || jobDescription == "" {
		http.Error(w, "Both resume and job description are required", http.StatusBadRequest)
		return
	}

	// Analyze keywords
	result, err := analyzer.Analyze(resume, jobDescription)
	if err != nil {
		http.Error(w, "Error analyzing keywords: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Send JSON response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}
