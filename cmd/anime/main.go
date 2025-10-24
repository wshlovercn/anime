package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	inputFile  = flag.String("input", "", "输入的小说文件路径")
	outputDir  = flag.String("output", "output", "输出目录")
	configFile = flag.String("config", "", "配置文件路径")
)

func main() {
	flag.Parse()

	if *inputFile == "" {
		fmt.Println("使用方法: anime -input <小说文件> [-output <输出目录>]")
		flag.PrintDefaults()
		os.Exit(1)
	}

	log.Printf("智能动漫生成系统启动...")
	log.Printf("输入文件: %s", *inputFile)
	log.Printf("输出目录: %s", *outputDir)

	// TODO: 实现核心生成逻辑
	log.Println("系统初始化完成，等待实现...")
}
