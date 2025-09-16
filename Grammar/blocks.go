package grammar

type Block struct {
	Data      Token
	BlockType string
	Children  []*Block
}
