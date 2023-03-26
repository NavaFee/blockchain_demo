package pbcc

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
	"time"
)

// Blcok 结构体
type Block struct {
	Height        int64          //区块高度
	PrevBlockHash []byte         //上一个区块的哈希值
	Txs           []*Transaction //交易数据data
	TimeStamp     int64          //时间戳
	Hash          []byte         //区块哈希
	Nonce         int64          //随机数
}

//func (block *Block) SetHash() {
//	//将高度转为字节数组
//	heightBytes := utils.IntToHex(block.Height)
//	//先将时间戳按二进制转化为字符串
//	timeString := strconv.FormatInt(block.TimeStamp, 2)
//	//强转为[]byte
//	timeBytes := []byte(timeString)
//	//拼接所有的属性， 把几个属性按照下面的空字节来分割拼接
//	blockBytes := bytes.Join([][]byte{
//		heightBytes,
//		block.PrevBlockHash,
//		block.Txs,
//		timeBytes,
//	}, []byte{})
//
//	//生成哈希值 返回一个32位的字节数组 256位
//	hash := sha256.Sum256(blockBytes)
//	block.Hash = hash[:]
//}

// 新建区块
func NewBlock(txs []*Transaction, prevBlockHash []byte, height int64) *Block {
	block := &Block{
		height,
		prevBlockHash,
		txs,
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
func CreateGenesisBlock(txs []*Transaction) *Block {
	return NewBlock(txs, make([]byte, 32), 0)
}

// 将区块序列化，得到一个字节数组 ---- 区块的行为
func (block *Block) Serialize() []byte {
	//创建一个Buffer
	var result bytes.Buffer
	//创建一个编码器  NewEncoder返回一个将编码后数据写入w的*Encoder。
	encoder := gob.NewEncoder(&result)
	//编码----> 打包  Encode方法将e编码后发送，并且会保证所有的类型信息都先发送。
	err := encoder.Encode(block)
	if err != nil {
		log.Panic(err)
	}
	return result.Bytes()
}

// 反序列化，得到以个区块
func DeserializeBlock(blockBytes []byte) *Block {
	var block Block
	var reader = bytes.NewReader(blockBytes)
	//创建一个解码器
	decoder := gob.NewDecoder(reader)
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}
	return &block
}

// 将Txs转化为[]byte
func (block *Block) HashTransactions() []byte {
	var txHashes [][]byte
	var txHash [32]byte
	for _, tx := range block.Txs {
		txHashes = append(txHashes, tx.TxID)
	}
	txHash = sha256.Sum256(bytes.Join(txHashes, []byte{}))
	return txHash[:]
}
