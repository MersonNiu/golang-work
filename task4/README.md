# GoBlogProject
GoBlogProject 是一个使用 Golang 编写的博客后端服务，支持用户注册、登录、发布文章、评论等基本功能，采用 RESTful API 设计，使用 Gin 框架 + GORM + MySQL + JWT 实现。
---
## 📁 项目结构
```
GoBlogProject/
├── config/                # 数据库连接配置
│   └── db.go
├── handlers/              # 路由处理逻辑
│   ├── auth.go
│   ├── comment.go
│   ├── post.go
│   └── services.go
├── middleware/            # 中间件（JWT验证、日志记录）
│   ├── loggerMiddleware.go
│   └── middleware.go
├── models/                # 数据模型定义
│   └── models.go
├── utils/                 # 工具类（错误处理、token生成等）
│   ├── error.go
│   └── token.go
├── GoBlogProject.postman_collection.json  # Postman 接口测试集合
├── main.go                # 项目入口
├── go.mod / go.sum        # Go 依赖管理
└── README.md              # 项目说明文件
```
---
## 🛠️ 环境要求
- Go 版本：go1.24.4 windows/amd64
- MySQL数据库：`8.0`
## 🚀 快速开始

### 1. 克隆项目

```bash
git clone https://github.com/yourname/GoBlogProject.git
cd GoBlogProject
```

### 2. 安装依赖

```bash
go mod tidy
```

### 3. 配置数据库

修改 `config/db.go` 中的数据库连接字符串：

```go
dsn := "root:yourpassword@tcp(127.0.0.1:3306)/goblog?charset=utf8mb4&parseTime=True&loc=Local"
```

请确保你已经在本地创建了数据库：

```sql
CREATE DATABASE goblog CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
```

### 4. 启动项目

```bash
go run main.go
```

服务默认运行在 `http://localhost:8080`。

---