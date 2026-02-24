package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "显示当前版本",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("todo-cli %s\n", version)
	},
}
