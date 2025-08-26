// 获取DOM元素
const resumeTextarea = document.getElementById('resume');
const jobDescriptionTextarea = document.getElementById('jobDescription');
const analyzeBtn = document.getElementById('analyzeBtn');
const resultsSection = document.getElementById('results');
const matchedKeywordsList = document.getElementById('matchedKeywords');
const missingKeywordsList = document.getElementById('missingKeywords');
const matchScoreValue = document.getElementById('matchScoreValue');
const matchedCount = document.getElementById('matchedCount');
const missingCount = document.getElementById('missingCount');
const toggleMatchedBtn = document.getElementById('toggleMatchedBtn');
const toggleMissingBtn = document.getElementById('toggleMissingBtn');
const encouragementSection = document.getElementById('encouragementSection');
const encouragementText = document.getElementById('encouragementText');

// 添加事件监听器
analyzeBtn.addEventListener('click', analyzeKeywords);
toggleMatchedBtn.addEventListener('click', toggleMatchedKeywords);
toggleMissingBtn.addEventListener('click', toggleMissingKeywords);

// 关键字显示状态
let matchedKeywordsExpanded = false;
let missingKeywordsExpanded = false;

// 分析关键字函数
async function analyzeKeywords() {
    const resume = resumeTextarea.value.trim();
    const jobDescription = jobDescriptionTextarea.value.trim();
    
    // 验证输入
    if (!resume || !jobDescription) {
        alert('请填写简历和职位描述');
        return;
    }
    
    // 禁用按钮并显示加载状态
    analyzeBtn.disabled = true;
    analyzeBtn.innerHTML = '<span class="loading"></span> 分析中...';
    
    try {
        // 发送API请求
        const response = await fetch('/api/analyze', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/x-www-form-urlencoded',
            },
            body: `resume=${encodeURIComponent(resume)}&jobDescription=${encodeURIComponent(jobDescription)}`
        });
        
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        
        const result = await response.json();
        
        // 保存结果
        lastAnalysisResult = result;
        
        // 显示结果
        displayResults(result);
    } catch (error) {
        console.error('Error:', error);
        alert('分析过程中出现错误，请重试');
    } finally {
        // 恢复按钮状态
        analyzeBtn.disabled = false;
        analyzeBtn.textContent = '分析关键字';
    }
}

// 显示结果函数
function displayResults(result) {
	// 清空之前的結果
	matchedKeywordsList.innerHTML = '';
	missingKeywordsList.innerHTML = '';
	
	// 显示匹配度
	if (result.match_score !== undefined) {
		matchScoreValue.textContent = result.match_score;
		drawMatchScoreChart(result.match_score);
	}
	
	// 显示匹配的关键字数量
	if (result.matched_keywords) {
		matchedCount.textContent = result.matched_keywords.length;
	}
	
	// 显示缺失的关键字数量
	if (result.missing_keywords) {
		missingCount.textContent = result.missing_keywords.length;
	}
	
	// 显示匹配的关键字（限制显示前10个）
	if (result.matched_keywords && result.matched_keywords.length > 0) {
		displayKeywords(result.matched_keywords, matchedKeywordsList, 10);
		updateToggleButton(toggleMatchedBtn, result.matched_keywords.length > 10);
	} else {
		matchedKeywordsList.innerHTML = '<div class="keyword-tag">没有匹配的关键字</div>';
		updateToggleButton(toggleMatchedBtn, false);
	}
	
	// 显示缺失的关键字（限制显示前10个）
	if (result.missing_keywords && result.missing_keywords.length > 0) {
		displayKeywords(result.missing_keywords, missingKeywordsList, 10);
		updateToggleButton(toggleMissingBtn, result.missing_keywords.length > 10);
	} else {
		missingKeywordsList.innerHTML = '<div class="keyword-tag">没有缺失的关键字</div>';
		updateToggleButton(toggleMissingBtn, false);
	}
	
	// 显示加油鼓励
	showEncouragement(result.match_score);
	
	// 显示结果区域
	resultsSection.style.display = 'block';
	
	// 重置折叠状态
	matchedKeywordsExpanded = false;
	missingKeywordsExpanded = false;
	
	// 滚动到结果区域
	resultsSection.scrollIntoView({ behavior: 'smooth' });
}

// 显示关键字函数
function displayKeywords(keywords, container, limit) {
	container.innerHTML = '';
	const displayKeywords = matchedKeywordsExpanded ? keywords : keywords.slice(0, limit);
	
	displayKeywords.forEach(keyword => {
		const tag = document.createElement('div');
		tag.className = 'keyword-tag';
		tag.textContent = keyword;
		container.appendChild(tag);
	});
}

// 更新折叠按钮
function updateToggleButton(button, hasMore) {
	if (hasMore) {
		button.style.display = 'block';
		button.textContent = '展开全部';
	} else {
		button.style.display = 'none';
	}
}

// 保存最后一次分析结果
let lastAnalysisResult = null;

// 切换匹配关键字显示
function toggleMatchedKeywords() {
	if (!lastAnalysisResult) return;
	
	matchedKeywordsExpanded = !matchedKeywordsExpanded;
	displayKeywords(lastAnalysisResult.matched_keywords, matchedKeywordsList, matchedKeywordsExpanded ? Infinity : 10);
	toggleMatchedBtn.textContent = matchedKeywordsExpanded ? '收起' : '展开全部';
}

// 切换缺失关键字显示
function toggleMissingKeywords() {
	if (!lastAnalysisResult) return;
	
	missingKeywordsExpanded = !missingKeywordsExpanded;
	displayKeywords(lastAnalysisResult.missing_keywords, missingKeywordsList, missingKeywordsExpanded ? Infinity : 10);
	toggleMissingBtn.textContent = missingKeywordsExpanded ? '收起' : '展开全部';
}

// 绘制匹配度圆饼图
function drawMatchScoreChart(score) {
	// 检查Chart.js库是否已加载
	if (typeof Chart === 'undefined') {
		console.error('Chart.js library not loaded');
		return;
	}
	
	const ctx = document.getElementById('matchScoreChart').getContext('2d');
	
	// 清除之前的图表
	if (window.matchScoreChart && typeof window.matchScoreChart.destroy === 'function') {
		window.matchScoreChart.destroy();
	}
	
	// 根据分数确定颜色
	let scoreColor;
	if (score >= 80) {
		scoreColor = '#4caf50'; // 绿色 - 优秀
	} else if (score >= 60) {
		scoreColor = '#2196f3'; // 蓝色 - 良好
	} else if (score >= 40) {
		scoreColor = '#ff9800'; // 橙色 - 一般
	} else {
		scoreColor = '#f44336'; // 红色 - 需改进
	}
	
	// 创建新的图表
	window.matchScoreChart = new Chart(ctx, {
		type: 'doughnut',
		data: {
			labels: ['匹配度', '未匹配'],
			datasets: [{
				data: [score, 100 - score],
				backgroundColor: [
					scoreColor,
					'#e0e0e0' // 浅灰色
				],
				borderColor: [
					'white',
					'white'
				],
				borderWidth: 3,
				// 添加阴影效果
				hoverOffset: 10,
				// 添加边框圆角
				borderRadius: 3,
				// 添加间距
				spacing: 2
			}]
		},
		options: {
			cutout: '75%', // 增加中心空白区域
			responsive: true,
			maintainAspectRatio: false,
			plugins: {
				legend: {
					display: false
				},
				tooltip: {
					enabled: true,
					// 自定义提示框
					callbacks: {
						label: function(context) {
							return `${context.label}: ${context.raw}%`;
						}
					}
				}
			},
			animation: {
				// 优化动画效果
				animateRotate: true,
				animateScale: true,
				// 增加动画持续时间
				duration: 2000,
				// 添加弹性效果
				easing: 'easeOutBounce'
			},
			// 添加旋转动画
			rotation: -90,
			// 添加圆周率
			circumference: 360
		}
	});
}

// 显示加油鼓励
function showEncouragement(score) {
	const encouragements = [
		"太棒了！你的简历已经非常接近完美了！继续加油⛽️",
		"不错的表现！你已经掌握了大部分关键技能，再努力一下就更好了！💪",
		"有进步空间！多关注一下缺失的关键字，你的简历会更出色！🌟",
		"继续努力！你的简历还有提升空间，相信你可以做到更好！🚀",
		"不要气馁！每个人都有起点，关键是要持续改进！🌈"
	];
	
	let text = "";
	if (score >= 80) {
		text = encouragements[0];
	} else if (score >= 60) {
		text = encouragements[1];
	} else if (score >= 40) {
		text = encouragements[2];
	} else if (score >= 20) {
		text = encouragements[3];
	} else {
		text = encouragements[4];
	}
	
	encouragementText.textContent = text;
	encouragementSection.style.display = 'block';
}

// 添加一些交互效果
document.addEventListener('DOMContentLoaded', function() {
    // 为文本区域添加焦点效果
    const textareas = document.querySelectorAll('textarea');
    textareas.forEach(textarea => {
        textarea.addEventListener('focus', function() {
            this.parentElement.classList.add('focused');
        });
        
        textarea.addEventListener('blur', function() {
            this.parentElement.classList.remove('focused');
        });
    });
    
    // 添加键盘快捷键 (Ctrl+Enter)
    document.addEventListener('keydown', function(event) {
        if (event.ctrlKey && event.key === 'Enter') {
            event.preventDefault();
            analyzeKeywords();
        }
    });
});
