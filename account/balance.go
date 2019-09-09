package account

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
)

func main() {
	account := common.HexToAddress("0x3AD041BF00Bd3b806F5Bd4aB96F9dcF1f05ab4EC")
	balance, errr := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance)
}
