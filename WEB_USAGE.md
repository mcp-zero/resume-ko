<!--
 * @Author: lucklidi@126.com
 * @Date: 2025-08-26 16:05:40
 * @LastEditTime: 2025-08-26 22:45:59
 * @LastEditors: lucklidi@126.com
 * @Description: 
 * Copyright (c) 2025 by lucklidi, All Rights Reserved. 
-->
# Web 界面使用说明

## 启动服务
```bash
./resume_optimizer -web 8168
```

## 访问界面
打开浏览器访问: http://localhost:8168

## 使用步骤
1. 在"您的简历"文本框中粘贴您的简历内容
2. 在"职位描述"文本框中粘贴目标职位的描述
3. 点击"分析关键字"按钮或按 Ctrl+Enter
4. 查看分析结果：
   - 匹配的关键字（您的简历中已包含的技能）
   - 缺失的关键字（您的简历中缺少的技能）

## 测试数据
项目中提供了示例文件供测试：
- 简历示例: resource/web_resume.txt
- 职位描述示例: resource/web_job.txt

示例数据包含英文内容，可产生较好的匹配效果。您也可以使用自己的中英文简历和职位描述进行测试。

## 快捷键
- Ctrl+Enter: 快速触发分析
