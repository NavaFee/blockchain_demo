package pbcc

// Transaction结构体
type Transaction struct {
	TxID  []byte     //交易ID
	Vins  []*TXInput // 输入
	Vouts []*TXOuput // 输出
}
