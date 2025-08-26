# 🚀 简历关键字优化工具 (Resume Keyword Optimizer)

[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/mcp-zero/resume-ko)](https://goreportcard.com/report/github.com/mcp-zero/resume-ko)
[![GitHub issues](https://img.shields.io/github/issues/mcp-zero/resume-ko)](https://github.com/mcp-zero/resume-ko/issues)
[![GitHub stars](https://img.shields.io/github/stars/mcp-zero/resume-ko)](https://github.com/mcp-zero/resume-ko/stargazers)
[![Docker](https://img.shields.io/badge/docker-%E2%9C%93-blue)](#docker部署)

## 📋 目录

- [🌟 简介](#-简介)
- [✨ 功能特点](#-功能特点)
- [⚡ 快速开始](#-快速开始)
  - [🐳 Docker一键部署](#-docker一键部署)
  - [💻 本地部署](#-本地部署)
- [📘 使用指南](#-使用指南)
  - [🖥️ Web界面模式](#-web界面模式)
  - [⌨️ 命令行模式](#-命令行模式)
- [📚 示例](#-示例)
- [🛠️ 技术架构](#-技术架构)
- [🤝 贡献](#-贡献)
- [📄 许可证](#-许可证)
- [📬 联系方式](#-联系方式)

## 🌟 简介

> **5% 的简历从未到达人类手中。由于缺少关键字，它们被申请人跟踪系统（ATS）过滤掉。**

你是否曾疑惑为什么你的简历总是石沉大海？简历关键字优化工具可以帮助你：

- 🔍 **扫描职位描述**：智能提取职位要求中的关键技能
- 🔄 **对比分析**：将你的简历与职位要求进行智能匹配
- 📋 **缺失关键字**：清晰列出你简历中遗漏的重要技能
- 🎯 **提高通过率**：优化简历关键字，提高通过ATS筛选系统的几率

无论你是编程新手还是经验丰富的开发者，这个工具都能帮助你轻松优化简历，让你的才华不再被埋没！

## ✨ 功能特点

- 🎯 **智能关键字分析**：自动提取职位描述中的核心技能
- 🔄 **实时对比功能**：即时查看简历与职位的匹配度
- 🌐 **美观Web界面**：直观的可视化界面，操作简单
- ⌨️ **命令行工具**：支持命令行操作，满足不同用户需求
- ⚡ **快速分析**：秒级获得分析结果
- 🌍 **跨平台支持**：支持Windows、macOS和Linux
- 🐳 **Docker支持**：提供Docker部署方案，一键启动
- 📁 **示例数据**：内置测试文件，快速上手体验

## ⚡ 快速开始

### 🐳 Docker一键部署（推荐）

只需安装Docker，即可一键启动应用：

```bash
# 克隆项目
git clone https://github.com/mcp-zero/resume-ko.git
cd resume-ko

# 一键启动
docker-compose up -d

# 访问应用
# 浏览器打开 http://localhost:8168

# 停止服务
docker-compose down
```

就是这么简单！🚀

### 💻 本地部署

#### 前提条件

- Go 1.21 或更高版本

#### 安装步骤

```bash
# 1. 克隆仓库
git clone https://github.com/mcp-zero/resume-ko.git

# 2. 进入项目目录
cd resume-ko

# 3. 构建项目
go build -o resume_optimizer

# 4. 启动Web服务
./resume_optimizer -web 8168

# 5. 访问应用
# 浏览器打开 http://localhost:8168
```

## 📘 使用指南

### 🖥️ Web界面模式

1. 启动服务后，打开浏览器访问 `http://localhost:8168`
2. 在"您的简历"文本框中粘贴您的简历内容
3. 在"职位描述"文本框中粘贴目标职位的描述
4. 点击"分析关键字"按钮或按 `Ctrl+Enter`
5. 查看分析结果：
   - ✅ 匹配的关键字（您的简历中已包含的技能）
   - ❌ 缺失的关键字（您的简历中缺少的技能）

### ⌨️ 命令行模式

```bash
# 分析简历和职位描述
./resume_optimizer -resume resource/test_resume.txt -job resource/test_job.txt
```

### 参数说明

- `-resume <file>`: 指定简历文件路径
- `-job <file>`: 指定职位描述文件路径
- `-web <port>`: 启动Web服务器并指定端口

## 📚 示例

### Web界面示例

1. 启动Web服务器：
   ```bash
   ./resume_optimizer -web 8168
   ```

2. 在浏览器中访问 `http://localhost:8168`

3. 使用内置示例文件：
   - 简历示例: `resource/web_resume.txt`
   - 职位描述示例: `resource/web_job.txt`

更详细的Web界面使用说明请查看 [WEB_USAGE.md](WEB_USAGE.md) 文件。

### 命令行示例

```bash
# 创建示例简历文件
echo "软件工程师，具有Go、Python和JavaScript开发经验" > resume.txt

# 创建示例职位描述文件
echo "我们需要一名软件工程师，精通Go语言和云计算技术" > job.txt

# 分析关键字
./resume_optimizer -resume resume.txt -job job.txt
```

## 🛠️ 技术架构

- **后端**：Go语言
- **前端**：原生HTML/CSS/JavaScript（无框架）
- **算法**：智能关键字提取和匹配算法

## 🤝 贡献

欢迎提交Issue和Pull Request来改进这个工具！

### 开发步骤

1. Fork仓库
2. 创建功能分支
3. 提交更改
4. 推送到分支
5. 创建Pull Request

请查看 [CONTRIBUTING.md](CONTRIBUTING.md) 文件了解更多贡献指南。

## 📄 许可证

本项目采用MIT许可证。详情请见[LICENSE](LICENSE)文件。

## 📬 联系方式

如果您有任何问题或建议，请通过以下方式联系：

- 提交Issue
- 发送邮件至：your-email@example.com

---

🎉 **让每一份优秀的简历都能被看见！** 🎉