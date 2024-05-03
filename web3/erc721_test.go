package web3

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"testing"
)

func TestERC721_New(t *testing.T) {
	address := common.HexToAddress("0x34d85c9cdeb23fa97cb08333b511ac86e1c4e258")
	c, err := NewErc721(address, getContractBackend(t))
	if err != nil {
		t.Fatal(err)
	}
	name, err := c.Name(nil)
	fmt.Println("get name:", name)
}

func getContractBackend(t *testing.T) bind.ContractBackend {
	return getEthClient(t)
}

func getEthClient(t *testing.T) *ethclient.Client {
	client, err := ethclient.Dial("https://eth-mainnet.alchemyapi.io/v2/_F8hUoqQRATbPdPidO7BEkwpAzxxFn0g")
	if err != nil {
		log.Fatal(err)
	}
	return client
}
