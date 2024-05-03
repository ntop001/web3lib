package web3

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"testing"
)

func TestERC1155_New(t *testing.T) {
	address := common.HexToAddress("0x6811f2f20c42f42656a3c8623ad5e9461b83f719")
	c, err := NewErc1155(address, getContractBackend(t))
	if err != nil {
		t.Fatal(err)
	}
	name, err := c.Name(nil)
	fmt.Println("get name:", name)
}
