package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	query        string
	detail       bool
	toClip       bool
	collectionID string
)

var rootCmd = &cobra.Command{
	Use:   "randpic",
	Short: "获取 Unsplash 随机图片",
	Long: `randpic 是一个命令行工具，用于从 Unsplash 获取随机图片。
可以指定搜索主题，选择特定集合，并选择是否显示详细的图片信息。`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := getRandomPic(); err != nil {
			fmt.Println("错误:", err)
			os.Exit(1)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&query, "query", "q", "", "搜索主题")
	rootCmd.PersistentFlags().BoolVarP(&detail, "detail", "d", false, "显示详细信息")
	rootCmd.PersistentFlags().BoolVarP(&toClip, "clipboard", "c", false, "将结果复制到剪贴板")
	rootCmd.PersistentFlags().StringVarP(&collectionID, "collection", "l", "", "指定集合ID")
}
