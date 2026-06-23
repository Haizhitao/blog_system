# 第一阶段：构建阶段
FROM golang:1.24-alpine AS builder

# 设置工作目录
WORKDIR /app

# 设置 Go 代理（国内加速）
ENV GOPROXY=https://goproxy.cn,direct
ENV GO111MODULE=on

# 复制依赖文件（利用 Docker 缓存）
COPY go.mod go.sum ./
RUN go mod download

# 复制所有源代码
COPY . .

# 编译成静态二进制文件
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags="-w -s" \
    -o blog_system

# 第二阶段：运行阶段
FROM alpine:latest

# 安装 CA 证书和时区数据
RUN apk --no-cache add ca-certificates tzdata

# 设置时区
ENV TZ=Asia/Shanghai

# 创建非 root 用户
RUN adduser -D -g '' appuser

# 切换到 appuser
USER appuser

# 创建工作目录
WORKDIR /app

# 从构建阶段复制二进制文件和配置文件
COPY --from=builder /app/blog_system .

# 暴露端口（根据你的项目修改，默认 8080）
EXPOSE 8080

# 健康检查
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:8080/health || exit 1

# 运行程序
CMD ["./blog_system"]