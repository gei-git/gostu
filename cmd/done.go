package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gei-git/todo-cli/internal/service"
	"github.com/spf13/cobra"
)

var doneCmd = &cobra.Command{
	Use:   "done [任务ID]",
	Short: "标记任务为已完成",
	Long:  `标记指定任务为已完成，例如：todo done 1`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("❌ 错误：请提供任务ID，例如：todo done 1")
			os.Exit(1)
		}

		id, err := strconv.ParseUint(args[0], 10, 32)
		if err != nil {
			fmt.Println("❌ 错误：任务ID必须是数字")
			os.Exit(1)
		}

		svc, err := service.NewTaskService()
		if err != nil {
			fmt.Printf("❌ 初始化失败: %v\n", err)
			os.Exit(1)
		}
		defer svc.Close()

		if err := svc.MarkTaskDone(uint(id)); err != nil {
			fmt.Printf("❌ 操作失败: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("✅ 任务 %d 已标记为完成！\n", id)
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
