package main

import "blockchain/cli"

func main() {
	//创建命令行工具
	cli := cli.CLI{}
	//激活cli
	cli.Run()

}
