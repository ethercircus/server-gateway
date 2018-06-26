package main
import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethercircus/server-gateway/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)
func main() {
	// Create an IPC based RPC connection to a remote node
	conn, err := ethclient.Dial("/home/cameron/.rinkeby_new/geth.ipc")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	// Instantiate the contract and display its name
	userContentRegister, err := contracts.NewUserContentRegister(common.HexToAddress("0x8979BDf9A55C8B5a8ce4b30c4D80D678b8a91800"), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate the UserContentRegister contract: %v", err)
	}
	ret, err := userContentRegister.UserIndex(nil, common.HexToAddress("0xe2eb4e5418e8d1f90b474318b83034a15fae409f"))
	if err != nil {
		log.Fatalf("Failed to retrieve username: %v", err)
	}
	fmt.Println("User name: ", ret.UserName)
	var blockNumber uint64 = 2530622
	//s := []uint32{}

	ch := make(chan *contracts.UserContentRegisterStoreData)
	opts := &bind.WatchOpts{}
	opts.Start = &blockNumber
	storeDataSub, err := userContentRegister.WatchStoreData(opts, ch)
	if err != nil {
		log.Fatalf("Failed WatchStoreData: %v", err)
	}

	go func() {
		defer storeDataSub.Unsubscribe()

		for {
			var newEvent *contracts.UserContentRegisterStoreData = <-ch
			fmt.Println(newEvent.Data)
		}
	}()

	//Monitor for interrupts and terminate cleanly
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt)
	defer signal.Stop(sigc)
	<-sigc
}
