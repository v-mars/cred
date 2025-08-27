# 使用说明文档

## 环境准备

### 系统要求
- Go 1.15 或更高版本
- Git (用于版本控制)

### 开发工具推荐
- VS Code 或 GoLand
- Postman (用于API测试)

## 安装部署步骤

### 1. 获取源代码
```bash
git clone <项目地址>
cd cred
```

### 2. 安装依赖
```bash
go mod tidy
```

### 3. 配置环境
在项目根目录创建 `.env` 文件，并添加必要的配置项：
```env
# 数据库配置
DB_HOST=localhost
DB_PORT=5432
DB_USER=user
DB_PASSWORD=password
DB_NAME=cred

# 服务端口
PORT=8080
```

### 4. 编译运行
```bash
# 编译
go build -o cred .

# 运行
./cred
```

或者直接运行:
```bash
go run main.go
```

## 配置说明

### 应用配置
| 配置项 | 说明 | 默认值 |
|--------|------|--------|
| PORT | 服务监听端口 | 8080 |
| LOG_LEVEL | 日志级别 | info |

### 数据库配置
| 配置项 | 说明 | 默认值 |
|--------|------|--------|
| DB_HOST | 数据库主机地址 | localhost |
| DB_PORT | 数据库端口 | 5432 |
| DB_USER | 数据库用户名 | - |
| DB_PASSWORD | 数据库密码 | - |
| DB_NAME | 数据库名称 | cred |

## 使用示例

### 启动服务
```bash
./cred
```
服务将在配置的端口上启动，默认为 `8080`。

### API调用示例
使用 curl 调用API:
```bash
curl -X GET http://localhost:8080/api/health
```

## 常见问题解答

### 1. 如何更改服务端口?
修改 `.env` 文件中的 `PORT` 配置项，或设置环境变量:
```bash
export PORT=9090
```

### 2. 如何查看应用日志?
日志将输出到标准输出，可以重定向到文件:
```bash
./cred > app.log 2>&1
```

### 3. 如何进行单元测试?
运行以下命令执行单元测试:
```bash
go test ./...
```

## 最佳实践

### 开发规范
- 遵循 Go 语言编码规范
- 编写单元测试
- 使用有意义的变量和函数命名
- 添加必要的注释

### 安全建议
- 不要在代码中硬编码敏感信息
- 定期更新依赖包
- 验证所有输入数据
- 使用HTTPS进行生产部署

### 性能优化
- 合理使用缓存
- 避免不必要的数据库查询
- 使用连接池
- 监控内存使用情况