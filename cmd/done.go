package cmd

import "github.com/spf13/cobra"

var doneCmd = &cobra.Command{
	Use:   "done [任务ID]",
	Short: "标记任务为已完成",
	Long:  `标记指定任务为已完成，例如：todo done 1`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: 后面实现
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
