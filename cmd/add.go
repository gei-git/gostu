package cmd

import "github.com/spf13/cobra"

var addCmd = &cobra.Command{
	Use:   "add [任务内容]",
	Short: "添加一个新任务",
	Long: `添加一个新任务，支持优先级、截止日期、标签。
示例：
  todo add "写周报" --priority=high --due=2025-03-01
  todo add "健身" --priority=low`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: 后面实现业务逻辑
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
