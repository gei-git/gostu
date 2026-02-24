package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	// 全局变量，后续可以加 config 等
	version = "v0.1.0"
)

var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "一个强大的命令行 Todo 任务管理系统",
	Long:  `todo-cli 是企业内部使用的命令行 Todo 管理系统\n支持优先级、截止日期、标签、BoltDB 持久化等企业级特性。`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("欢迎使用 todo-cli！输入 todo --help 查看所有命令")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Version = version
	rootCmd.Flags().BoolP("version", "v", false, "显示版本信息")
}
