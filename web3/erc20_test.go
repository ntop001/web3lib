package web3

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"testing"
)

func TestERC20_New(t *testing.T) {
	address := common.HexToAddress("0xf4d2888d29d722226fafa5d9b24f9164c092421e")
	c, err := NewErc20(address, getContractBackend(t))
	if err != nil {
		t.Fatal(err)
	}
	name, err := c.Name(nil)
	fmt.Println("get name:", name)
}
