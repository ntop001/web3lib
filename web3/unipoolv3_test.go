package web3

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"testing"
)

func TestUniV3_New(t *testing.T) {
	address := common.HexToAddress("0xE819013997B4186e1813590C51837Edb491001Dc")
	c, err := NewUnipoolv3(address, getContractBackend(t))
	if err != nil {
		t.Fatal(err)
	}
	result, err := c.Slot0(nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("get slot0:", result)
}
