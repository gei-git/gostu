package cmd

import "github.com/spf13/cobra"

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "列出所有任务",
	Long:  `列出所有任务（后续支持按状态、优先级过滤）。`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: 后面实现
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
