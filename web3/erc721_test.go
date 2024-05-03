package web3

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/joho/godotenv"
	"testing"
)

func TestERC721_New(t *testing.T) {
	_ = godotenv.Load("../.env")
	address := common.HexToAddress("0x34d85c9cdeb23fa97cb08333b511ac86e1c4e258")
	c, err := NewErc721(address, getContractBackend(t))
	if err != nil {
		t.Fatal(err)
	}
	name, err := c.Name(nil)
	fmt.Println("get name:", name)
}

func getContractBackend(t *testing.T) bind.ContractBackend {
	return getEthClient()
}
