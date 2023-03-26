package pbcc

import (
	"blockchain/conf"
	"blockchain/utils"
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"math/big"
	"time"
)

// 创建区块链
type BlockChain struct {
	Tip []byte   //最新区块的哈希值
	DB  *bolt.DB //数据库对象
}

// 创建区块链，带有创世区块
func CreateBlockChainWithGenesisBlock(address string) {
	if utils.DBExists() {
		fmt.Println("数据库已经存在")
		return
	}
	fmt.Println("数据库不存在，创建创世区块：")
	//先创建 coinbase 交易
	txCoinBase := NewCoinBaseTransaction(address)
	//创世区块
	genesisBlock := CreateGenesisBlock([]*Transaction{txCoinBase})
	//打开数据库
	db, err := bolt.Open(conf.DBNAME, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	//存入数据表
	err = db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucket([]byte(conf.BLOCKTABLENAME))
		if err != nil {
			log.Panic(err)
		}
		if b != nil {
			err = b.Put(genesisBlock.Hash, genesisBlock.Serialize())
			if err != nil {
				log.Panic("创世区块链存储有误")
			}
			//存储最新区块的hash
			b.Put([]byte("l"), genesisBlock.Hash)
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	////先判断数据库是否存在，如果存在，从数据库读取信息
	//if utils.DBExists() {
	//	fmt.Println("数据库已经存在")
	//	//打开数据库   0600 表示文件的所有者对该文件有读写的权限
	//	db, err := bolt.Open(conf.DBNAME, 0600, nil)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	var blockchain *BlockChain
	//	//读取数据库
	//	err = db.View(func(tx *bolt.Tx) error {
	//		//打开表
	//		b := tx.Bucket([]byte(conf.BLOCKTABLENAME))
	//		if b != nil {
	//			//读取最后一个hash
	//			hash := b.Get([]byte("l"))
	//			//创建blockchain
	//			blockchain = &BlockChain{hash, db}
	//		}
	//		return nil
	//	})
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	return blockchain
	//}
	////数据库不存在， 说明第一次创建，然后存入到数据库中
	//fmt.Println("数据库不存在，创建中")
	////创建创世区块
	//genesisBlock := CreateGenesisBlock(data)
	//db, err := bolt.Open(conf.DBNAME, 0600, nil)
	//if err != nil {
	//	log.Fatal(err)
	//}
	////存入数据表
	//err = db.Update(func(tx *bolt.Tx) error {
	//	b, err := tx.CreateBucket([]byte(conf.BLOCKTABLENAME))
	//	if err != nil {
	//		log.Panic(err)
	//	}
	//	if b != nil {
	//		err = b.Put(genesisBlock.Hash, genesisBlock.Serialize())
	//		if err != nil {
	//			log.Panic("创世区块存储错误")
	//		}
	//		//存储最新区块的哈希
	//		b.Put([]byte("l"), genesisBlock.Hash)
	//	}
	//	return nil
	//})
	//if err != nil {
	//	log.Panic(err)
	//}
	////返回区块链对象
	//return &BlockChain{genesisBlock.Hash, db}
}
func (bc *BlockChain) AddBlockToBlockChain(txs []*Transaction) {
	err := bc.DB.Update(func(tx *bolt.Tx) error {
		//打开表
		b := tx.Bucket([]byte(conf.BLOCKTABLENAME))
		if b != nil {
			//根据最新的块的hash读取数据，并反序列化最后一个区块
			blockBytes := b.Get(bc.Tip)
			lastBlock := DeserializeBlock(blockBytes)
			//创建新的区块，根据最后一个区块注入prevblockHash 和 height
			newBlock := NewBlock(txs, lastBlock.Hash, lastBlock.Height+1)
			//将新的区块序列化并存储
			err := b.Put(newBlock.Hash, newBlock.Serialize())
			if err != nil {
				log.Panic(err)
			}
			//更新最后一个哈希值以及blockchain的tip
			b.Put([]byte("l"), newBlock.Hash)
			bc.Tip = newBlock.Hash
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
}

// 获取一个迭代器
func (bc *BlockChain) Iterator() *BlockChainIterator {
	return &BlockChainIterator{bc.Tip, bc.DB}
}

// 借助迭代器来打印区块链
func (bc *BlockChain) PrintChains() {
	//获取迭代器对象
	bcIterator := bc.Iterator()
	//循环迭代
	for {
		block := bcIterator.Next()
		fmt.Printf("第%d个区块的信息：\n", block.Height+1)
		//获取当前hash对应的数据，并进行反序列化
		fmt.Printf("\t上一个区块的hash：%x\n", block.PrevBlockHash)
		fmt.Printf("\t当前的hash：%x\n", block.Hash)
		//fmt.Printf("\t数据：%s\n", block.Data)
		fmt.Println("\t交易：")
		for _, tx := range block.Txs {
			fmt.Printf("\t\t 交易ID: %x\n", tx.TxID)
			fmt.Println("\t\tVins:")
			for _, in := range tx.Vins {
				fmt.Printf("\t\t\tTxID: %x\n", in.TxID)
				fmt.Printf("\t\t\tVout: %d\n", in.Vout)
				fmt.Printf("\t\t\tScriptSig: %s\n", in.ScriptSig)
			}
			fmt.Println("\t\tVouts:")
			for _, out := range tx.Vouts {
				fmt.Printf("\t\t\tvalue:%d\n", out.Value)
				fmt.Printf("\t\t\tScriptPubKey: %s\n", out.ScriptPubkey)
			}
		}
		fmt.Printf("\t时间：%s\n", time.Unix(block.TimeStamp, 0).Format("2006-01-02 15:04:05"))
		fmt.Printf("\tNonce值：%d\n", block.Nonce)
		//直到父hash值是0
		hashInt := new(big.Int)
		hashInt.SetBytes(block.PrevBlockHash)
		if big.NewInt(0).Cmp(hashInt) == 0 {
			break
		}
	}
}

// 获取最新的区块链
func GetBlockchainObject() *BlockChain {
	if !utils.DBExists() {
		fmt.Println("数据库不存在，无法获取区块链")
		return nil
	}
	db, err := bolt.Open(conf.DBNAME, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	var blockchain *BlockChain
	// 读取数据库
	err = db.View(func(tx *bolt.Tx) error {
		//打开表
		b := tx.Bucket([]byte(conf.BLOCKTABLENAME))
		if b != nil {
			//读取最后一个hash
			hash := b.Get([]byte("l"))
			//创建blockchain
			blockchain = &BlockChain{hash, db}
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)

	}
	return blockchain
}
