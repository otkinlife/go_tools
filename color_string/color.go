package color_string

import "github.com/fatih/color"

var (
	// Green 绿色
	Green = color.New(color.FgGreen).SprintFunc()
	// Red 红色
	Red = color.New(color.FgRed).SprintFunc()
	// Yellow 黄色
	Yellow = color.New(color.FgYellow).SprintFunc()
	// Blue 蓝色
	Blue = color.New(color.FgBlue).SprintFunc()
	// Cyan 青色
	Cyan = color.New(color.FgCyan).SprintFunc()
	// Magenta 洋红色
	Magenta = color.New(color.FgMagenta).SprintFunc()
	// White 白色
	White = color.New(color.FgWhite).SprintFunc()
	// Black 黑色
	Black = color.New(color.FgBlack).SprintFunc()
	// HiGreen 亮绿色
	HiGreen = color.New(color.FgHiGreen).SprintFunc()
	// HiRed 亮红色
	HiRed = color.New(color.FgHiRed).SprintFunc()
	// HiYellow 亮黄色
	HiYellow = color.New(color.FgHiYellow).SprintFunc()
)
