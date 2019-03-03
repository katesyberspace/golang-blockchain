package blockchain

type BlockChain struct {
	Blocks []*Block
}

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

func CreateBlock(data string, prevHash []byte) *Block { //* to say use this block
	block := &Block{[]byte{}, []byte(data), prevHash, 0}
	pow := NewProof(block)

	//fmt.Printf("pow inside createBlock is: %v\n", pow)

	nonce, hash := pow.Run()

	//fmt.Printf("nonce inside createBlock is: %v\n", nonce)

	//fmt.Printf("hash inside createBlock is: %x\n", hash)
	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	newBlock := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, newBlock)
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}