# blog_system
使用 Go 语言结合 Gin 框架和 GORM 库开发一个个人博客系统的后端，实现博客文章的基本管理功能，包括文章的创建、读取、更新和删除（CRUD）操作，同时支持用户认证和简单的评论功能


## 运行环境
- Go 1.18+
- MySQL 5.7+
- Gin 框架
- GORM

## 依赖安装步骤
### 安装gorm
```bash
go get -u gorm.io/gorm
```
### 安装gorm-mysql
```bash
go get -u gorm.io/driver/mysql
```
### 安装gin
```bash
go get -u github.com/gin-gonic/gin
```
### 安装godotenv
```bash
go get -u github.com/joho/godotenv
```
### 安装jwt
```bash
go get -u github.com/dgrijalva/jwt-go
```

## 启动方式
在根目录下执行
```bash
go run main.go
```
## 功能列表
- 用户注册、登录
- 创建文章（只有已认证的用户才能创建）
- 获取所有文章列表和单个文章的详细信息
- 文章的更新和删除（只有文章的作者才能操作）
- 创建评论（已认证的用户才能对文章评论）
- 获取某篇文章的所有评论列表

## API 文档
| 路由                 | 方法 | 描述 |
|--------------------|------|------|
| /user/register     | POST | 用户注册 |
| /user/login        | POST | 用户登录 |
| /posts             | GET | 获取文章列表 |
| /posts/:id         | GET | 获取单篇文章 |
| /posts             | POST | 创建文章 |
| /posts/:id         | PUT | 更新文章 |
| /posts/:id         | DELETE | 删除文章 |
| /comments/post/:id | GET | 获取文章评论 |
| /comments          | POST | 创建评论 |

## 接口列表
### 1. 用户注册
**请求方式**：`POST`  
**URL**：`http://127.0.0.1:8080/user/register`  
**请求参数**：
```json
{
  "username": "liuhaitao",
  "password": "123456",
  "email": "liuht@qq.com"
}
```
**响应示例**：
```json
{
  "message": "用户注册成功"
}
```
### 2. 用户登录
**请求方式**：`POST`  
**URL**：`http://127.0.0.1:8080/user/login  
**请求参数**：
```json
{
  "username": "liuhaitao2",
  "password": "123456"
}
```
**响应示例**：
```json
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJoYWl6aGl0YW8tYmxvZy1zeXN0ZW0iLCJleHAiOjE3NTQ4Mzc2NzgsIm5iZiI6MTc1NDc1MTI3OCwiaWF0IjoxNzU0NzUxMjc4LCJVc2VybmFtZSI6ImxpdWhhaXRhbzIiLCJVc2VySUQiOjN9.DWPaVmjlySw-KccnM8Ejts01gjQPrQtTa6eruleMhXA"
}
```
### 3. 创建文章
**请求方式**：`POST`  
**URL**：`http://127.0.0.1:8080/posts`  
**Headers**：
```json
{
  "Authorization": "Bearer <token>",
  "Content-Type": "application/json"
}
```
**请求参数**：
```json
{
    "title": "我爱中国4",
    "content": "我爱中国我爱中国我爱中国我爱中国我爱中国我爱中国我爱中国我爱中国我爱中国"
}
```
**响应示例**：
```json
{
    "message": "创建成功"
}
```
### 4. 获取文章列表
**请求方式**：`GET`  
**URL**：`http://127.0.0.1:8080/posts`  
**请求参数**：

**响应示例**：
```json
[
    {
        "CreatedAt": "2025-08-09T00:32:34.887+08:00",
        "UpdatedAt": "2025-08-09T10:02:32.554+08:00",
        "DeletedAt": null,
        "ID": 1,
        "Title": "I love China",
        "Content": "我爱中国我爱中国我爱中国我爱中国我爱中国我爱中国我爱中国我爱中国我爱中国",
        "UserID": 1
    },
    {
        "CreatedAt": "2025-08-09T00:45:25.977+08:00",
        "UpdatedAt": "2025-08-09T00:45:25.977+08:00",
        "DeletedAt": null,
        "ID": 2,
        "Title": "我爱中国2",
        "Content": "我爱中国我爱中国我爱中国我爱中国我爱中国我爱中国我爱中国我爱中国我爱中国",
        "UserID": 1
    },
    {
        "CreatedAt": "2025-08-09T22:56:20.426+08:00",
        "UpdatedAt": "2025-08-09T22:56:20.426+08:00",
        "DeletedAt": null,
        "ID": 4,
        "Title": "我爱中国4",
        "Content": "我爱中国我爱中国我爱中国我爱中国我爱中国我爱中国我爱中国我爱中国我爱中国",
        "UserID": 1
    }
]
```
### 5. 获取单篇文章信息
**请求方式**：`GET`  
**URL**：`http://127.0.0.1:8080/posts/1`  
**请求参数**：

**响应示例**：
```json
{
  "post": {
    "CreatedAt": "2025-08-09T00:32:34.887+08:00",
    "UpdatedAt": "2025-08-09T23:02:34.533+08:00",
    "DeletedAt": null,
    "ID": 1,
    "Title": "I love China",
    "Content": "我爱中国我爱中国我爱中国我爱中国我爱中国我爱中国我爱中国我爱中国我爱中国",
    "UserID": 1
  }
}
```
### 6. 更新文章
**请求方式**：`PUT`  
**URL**：`http://127.0.0.1:8080/posts/1`  
**Headers**：
```json
{
  "Authorization": "Bearer <token>",
  "Content-Type": "application/json"
}
```
**请求参数**：
```json
{
    "title":"I love China"
}
```
**响应示例**：
```json
{
    "message": "更新成功"
}
```
### 7. 删除文章
**请求方式**：`DELETE`  
**URL**：`http://127.0.0.1:8080/posts/3`  
**Headers**：
```json
{
  "Authorization": "Bearer <token>",
  "Content-Type": "application/json"
}
```
**请求参数**：

**响应示例**：
```json
{
    "message": "删除成功"
}
```
### 8. 获取单篇文章的评论列表
**请求方式**：`GET`  
**URL**：`http://127.0.0.1:8080/comments/post/1`  
**请求参数**：

**响应示例**：
```json
{
  "comments": [
    {
      "UserID": 1,
      "PostID": 1,
      "Content": "不错的文章，点赞"
    },
    {
      "UserID": 1,
      "PostID": 1,
      "Content": "棒的很"
    }
  ]
}
```
### 9. 创建文章评论
**请求方式**：`POST`  
**URL**：`http://127.0.0.1:8080/comments`  
**Headers**：
```json
{
  "Authorization": "Bearer <token>",
  "Content-Type": "application/json"
}
```
**请求参数**：
```json
{
    "post_id":"1",
    "content":"棒的很"
}
```
**响应示例**：
```json
{
  "message": "创建成功"
}
```