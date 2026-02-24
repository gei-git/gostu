package cmd

import (
	"fmt"
	"os"

	"github.com/gei-git/todo-cli/internal/service"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "列出所有任务",
	Long:  `列出所有任务（后续会支持过滤）。`,
	Run: func(cmd *cobra.Command, args []string) {
		svc, err := service.NewTaskService()
		if err != nil {
			fmt.Printf("❌ 初始化失败: %v\n", err)
			os.Exit(1)
		}
		defer svc.Close()

		tasks, err := svc.ListTasks()
		if err != nil {
			fmt.Printf("❌ 查询失败: %v\n", err)
			os.Exit(1)
		}

		if len(tasks) == 0 {
			fmt.Println("📭 当前没有任何任务，快去添加吧！")
			return
		}

		fmt.Println("📋 任务列表：")
		for _, t := range tasks {
			status := "⭕ 未完成"
			if t.Done {
				status = "✅ 已完成"
			}
			fmt.Printf("%s [%d] %s  (优先级: %s)\n", status, t.ID, t.Title, t.Priority)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
