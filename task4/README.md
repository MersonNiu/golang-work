1.项目介绍GoBlogProject
GoBlogProject 是一个使用 Golang 编写的博客后端服务，支持用户注册、登录、发布文章、评论等基本功能，采用 RESTful API 设计，使用 Gin 框架 + GORM + MySQL + JWT 实现。
2.项目结构
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
3.环境要求
- Go 版本：go1.24.4 windows/amd64
- MySQL数据库：`8.0`
4.启动项目
go run main.go
服务默认运行在 `http://localhost:8080`。