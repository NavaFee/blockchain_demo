package pbcc

// 输入结构体
type TXInput struct {
	TxID      []byte //交易的ID
	Vout      int    // 存储Txoutput 的vout里面的索引
	ScriptSig string //用户名
}
