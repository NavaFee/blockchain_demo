package cli

import (
	"flag"
	"fmt"
	"log"
	"os"
)

// CLI结构体
type CLI struct {
}

// 添加Run 方法
func (cli *CLI) Run() {
	//判断命令行参数的长度
	fmt.Println("--------------", os.Args)
	isValidArgs()
	//创建 flagset 标签对象
	addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
	createBlockChainCmd := flag.NewFlagSet("createblockchain", flag.ExitOnError)
	//设置标签后的参数
	flagAddBlockData := addBlockCmd.String("data", "helloworld", "交易数据")
	flagCreateBlockChainData := createBlockChainCmd.String("data", "Genesis block data", "创世区块交易数据")

	//解析
	switch os.Args[1] {
	case "addblock":
		err := addBlockCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "printchain":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "createblockchaincmd":
		err := createBlockChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		fmt.Println("请检查输入的参数")
		printUsage()
		os.Exit(1)
	}
	if addBlockCmd.Parsed() {
		if *flagAddBlockData == "" {
			printUsage()
			os.Exit(1)
		}
		cli.addBlock(*flagAddBlockData)
	}
	if createBlockChainCmd.Parsed() {
		if *flagCreateBlockChainData == "" {
			printUsage()
			os.Exit(1)
		}
		cli.createGenesisBlockchain(*flagCreateBlockChainData)
	}
	if printChainCmd.Parsed() {
		cli.printChains()
	}

}

// 检查是否有命令行的参数
func isValidArgs() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}
}

// 打印提示
func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("\t createblockchain -data DATA -- 创建创世区块")
	fmt.Println("\t addblock -data Data -- 交易数据")
	fmt.Println("\t printchain --输出信息")
}
