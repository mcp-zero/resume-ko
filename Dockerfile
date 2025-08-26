# 使用Go 1.21作为基础镜像
FROM golang:1.21-alpine

# 设置工作目录
WORKDIR /app

# 复制go mod和sum文件
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制源代码
COPY . .

# 构建应用程序
RUN go build -o resume_optimizer

# 暴露端口8168
EXPOSE 8168

# 设置入口点
ENTRYPOINT ["./resume_optimizer", "-web", "8168"]