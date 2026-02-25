# todo-cli

企业内部命令行 Todo 管理系统（Golang 实战项目）

## 功能
- ✅ 添加任务 `todo add "任务内容"`
- ✅ 列出任务 `todo list`
- ✅ 完成任务 `todo done <ID>`
- ✅ 数据持久化（BoltDB，本地 tasks.db）

## 使用
```bash
go run main.go add "学习 Go"
go run main.go list
go run main.go done 1