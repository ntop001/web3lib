package web3

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"os"
)

func getEthClient() *ethclient.Client {
	fmt.Println("use rpc:", os.Getenv("RPC_ENDPOINT"))
	client, err := ethclient.Dial(os.Getenv("RPC_ENDPOINT"))
	if err != nil {
		log.Fatal(err)
	}
	return client
}
