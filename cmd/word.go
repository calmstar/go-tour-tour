package cmd

import (
	"github.com/go-programming-tour-book/tour/internal/word"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

const (
	ModeUpper                      = iota + 1 // 转成大写
	ModeLower                                 // 转成小写
	ModeUnderscoreToUpperCamelCase            // 下划线转成大写驼峰
	ModeUnderscoreToLowerCamelCase            // 下划线转成小写驼峰
	ModeCamelCaseToUnderscore                 // 驼峰转成下划线
)

var desc = strings.Join([]string{
	"该子命令支持各种单词格式转换，模式如下：",
	"1：全部转大写",
	"2：全部转小写",
	"3：下划线转大写驼峰",
	"4：下划线转小写驼峰",
	"5：驼峰转下划线",
}, "\n")

var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "单词格式转换",
	Long:  "支持多种单词格式转换",
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case ModeUpper:
			content = word.ToUpper(str)
		case ModeLower:
			content = word.ToLower(str)
		case ModeUnderscoreToUpperCamelCase:
			content = word.UnderscoreToUpperCamelCase(str)
		case ModeUnderscoreToLowerCamelCase:
			content = word.UnderscoreToLowerCamelCase(str)
		case ModeCamelCaseToUnderscore:
			content = word.CamelCaseToUnderscore(str)
		default:
			log.Fatalf("暂不支持该转换模式，请执行 help word 查看帮助文档")
		}
	},
}

func init() {}

var str string
var mode int8
