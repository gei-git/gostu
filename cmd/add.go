package cmd

import (
	"fmt"
	"os"

	"github.com/gei-git/todo-cli/internal/service"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [任务内容]",
	Short: "添加一个新任务",
	Long: `添加一个新任务，支持优先级、截止日期、标签。
示例：
  todo add "写周报" --priority=high --due=2025-03-01`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("❌ 错误：请提供任务内容，例如：todo add \"学习 Go\"")
			os.Exit(1)
		}

		title := args[0]

		svc, err := service.NewTaskService()
		if err != nil {
			fmt.Printf("❌ 初始化失败: %v\n", err)
			os.Exit(1)
		}
		defer svc.Close()

		task, err := svc.AddTask(title)
		if err != nil {
			fmt.Printf("❌ 添加失败: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("✅ 添加成功！ID: %d\n", task.ID)
		fmt.Printf("   标题: %s\n", task.Title)
		fmt.Printf("   优先级: %s\n", task.Priority)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
