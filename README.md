# 内容管理系统

本项目是一个基于Go语言开发的内容管理系统，包含多个子模块，用于实现内容管理、流程管理等功能。

## 项目结构

- `content_flow/`: 内容流程管理模块
- `content_manage/`: 内容管理核心模块
- `content_system/`: 内容系统模块
- `content_system_single/`: 单机版内容系统模块

## 主要功能

- 内容创建、更新、删除、查询
- 内容审核流程管理
- 用户权限管理
- 内容统计与分析

## 技术栈

- 编程语言: Go
- Web框架: Gin
- 数据库: MySQL
- 缓存: Redis
- 消息队列: RabbitMQ
- 容器化: Docker

## 快速开始

1. 克隆项目
```bash
git clone https://github.com/your-repo/content-system.git
```

2. 安装依赖
```bash
go mod tidy
```

3. 配置环境变量
复制`.env.example`为`.env`并修改配置

4. 启动服务
```bash
go run cmd/content_manage/main.go
```

## 贡献指南

欢迎提交Pull Request，请遵循以下规范：
- 提交信息使用英文
- 代码风格遵循Go官方规范
- 新功能需附带单元测试

## 许可证

本项目采用MIT许可证，详情请见LICENSE文件。
