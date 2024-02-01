/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/immafrady/studybuddy/internal/database"
	"github.com/manifoldco/promptui"
	"log"

	"github.com/spf13/cobra"
)

// resetCmd represents the reset command
var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "所有数据重置",
	Long:  `所有数据重置`,
	Run: func(cmd *cobra.Command, args []string) {
		force, _ := cmd.Flags().GetBool("force")
		if !force {
			prompt := promptui.Prompt{
				Label:     "确定要重置所有数据吗？",
				IsConfirm: true,
			}
			str, _ := prompt.Run()
			if str != "y" {
				return
			}
		}
		database.Reset()
		log.Println("数据已重置")
	},
}

func init() {
	rootCmd.AddCommand(resetCmd)

	resetCmd.Flags().BoolP("force", "f", false, "强制执行")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// resetCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// resetCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
