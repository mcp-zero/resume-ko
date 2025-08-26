# 简历关键字优化工具 (Resume Keyword Optimizer)

[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/mcp-zero/resume-ko)](https://goreportcard.com/report/github.com/mcp-zero/resume-ko)
[![GitHub issues](https://img.shields.io/github/issues/mcp-zero/resume-ko)](https://github.com/mcp-zero/resume-ko/issues)
[![GitHub stars](https://img.shields.io/github/stars/mcp-zero/resume-ko)](https://github.com/mcp-zero/resume-ko/stargazers)

## 目录

- [简介](#简介)
- [功能特点](#功能特点)
- [安装](#安装)
  - [前提条件](#前提条件)
  - [安装步骤](#安装步骤)
- [使用方法](#使用方法)
  - [命令行模式](#命令行模式)
  - [Web界面模式](#web界面模式)
- [示例](#示例)
  - [命令行示例](#命令行示例)
  - [Web界面示例](#web界面示例)
- [技术架构](#技术架构)
- [贡献](#贡献)
  - [开发步骤](#开发步骤)
- [许可证](#许可证)
- [联系方式](#联系方式)

## 简介

5% 的简历从未到达人类手中。由于缺少关键字，它们被申请人跟踪系统（ATS）过滤掉。

简历关键字优化工具可以帮助您：
- 扫描职位描述
- 将其与您的简历进行比较
- 输出您遗漏的关键字
- 提高通过ATS筛选系统的几率

这个工具既可以作为命令行工具使用，也提供了美观的Web界面，方便不同技术水平的用户使用。

## 功能特点

- **关键字分析**：智能提取职位描述中的关键字
- **对比功能**：比较简历与职位描述的匹配度
- **Web界面**：提供美观的Web界面，方便使用
- **命令行工具**：支持命令行操作
- **实时分析**：快速获得分析结果
- **跨平台支持**：支持Windows、macOS和Linux

## 安装

### 前提条件

- Go 1.21 或更高版本

### 安装步骤

1. 克隆仓库：
   ```bash
   git clone https://github.com/mcp-zero/resume-ko.git
   ```

2. 进入项目目录：
   ```bash
   cd resume-ko
   ```

3. 构建项目：
   ```bash
   go build -o resume_optimizer
   ```

## 使用方法

### 命令行模式

```bash
# 分析简历和职位描述
./resume_optimizer -resume your_resume.txt -job job_description.txt
```

### Web界面模式

```bash
# 启动Web服务器
./resume_optimizer -web 8080
```

然后在浏览器中访问 `http://localhost:8080`

### 参数说明

- `-resume <file>`: 指定简历文件路径
- `-job <file>`: 指定职位描述文件路径
- `-web <port>`: 启动Web服务器并指定端口

## 示例

### 命令行示例

```bash
# 创建示例简历文件
echo "软件工程师，具有Go、Python和JavaScript开发经验" > resume.txt

# 创建示例职位描述文件
echo "我们需要一名软件工程师，精通Go语言和云计算技术" > job.txt

# 分析关键字
./resume_optimizer -resume resume.txt -job job.txt
```

### Web界面示例

1. 启动Web服务器：
   ```bash
   ./resume_optimizer -web 8080
   ```

2. 在浏览器中访问 `http://localhost:8080`

3. 在文本框中粘贴简历和职位描述

4. 点击"分析关键字"按钮

更详细的Web界面使用说明请查看 [WEB_USAGE.md](WEB_USAGE.md) 文件。

## 技术架构

- **后端**：Go语言
- **前端**：原生HTML/CSS/JavaScript（无框架）
- **算法**：智能关键字提取和匹配算法

## 贡献

欢迎提交Issue和Pull Request来改进这个工具。

### 开发步骤

1. Fork仓库
2. 创建功能分支
3. 提交更改
4. 推送到分支
5. 创建Pull Request

请查看 [CONTRIBUTING.md](CONTRIBUTING.md) 文件了解更多贡献指南。

## 许可证

本项目采用MIT许可证。详情请见[LICENSE](LICENSE)文件。

## 联系方式

如果您有任何问题或建议，请通过以下方式联系：

- 提交Issue
- 发送邮件至：your-email@example.com