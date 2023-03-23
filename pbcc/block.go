package pbcc

import (
	"blockchain/utils"
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

// Blcok 结构体
type Block struct {
	Height        int64  //区块高度
	PrevBlockHash []byte //上一个区块的哈希值
	Data          []byte //交易数据data
	TimeStamp     int64  //时间戳
	Hash          []byte //区块哈希
	Nonce         int64  //随机数
}

func (block *Block) SetHash() {
	//将高度转为字节数组
	heightBytes := utils.IntToHex(block.Height)
	//先将时间戳按二进制转化为字符串
	timeString := strconv.FormatInt(block.TimeStamp, 2)
	//强转为[]byte
	timeBytes := []byte(timeString)
	//拼接所有的属性， 把几个属性按照下面的空字节来分割拼接
	blockBytes := bytes.Join([][]byte{
		heightBytes,
		block.PrevBlockHash,
		block.Data,
		timeBytes,
	}, []byte{})

	//生成哈希值 返回一个32位的字节数组 256位
	hash := sha256.Sum256(blockBytes)
	block.Hash = hash[:]
}

// 新建区块
func NewBlock(data string, prevBlockHash []byte, height int64) *Block {
	block := &Block{
		height,
		prevBlockHash,
		[]byte(data),
		time.Now().Unix(),
		nil,
		0,
	}
	//调用工作量证明的方法，并且返回有效的Hash 和 Nonce
	pow := NewProofOfWork(block)
	hash, nonce := pow.Run()
	block.Hash = hash
	block.Nonce = nonce
	return block
}

// 创建创世区块
func CreateGenesisBlock(data string) *Block {
	return NewBlock(data, make([]byte, 32), 0)
}
