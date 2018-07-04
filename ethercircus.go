package main
import (
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"time"
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethercircus/solidity-contracts/contracts/gocontracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/log"
)
func main() {
	//Create and Boot up the IPFS node
	ipfs, err := makeIpfs(filepath.Join(os.Getenv("HOME"), ".ethercircus"))
	if err != nil {
		log.Crit("Failed to create IPFS node", "err", err)
	}
	defer ipfs.Close()


	// Create an IPC based RPC connection to a remote node
	conn, err := ethclient.Dial("/home/cameron/.rinkeby_new/geth.ipc")
	if err != nil {
		log.Crit("Failed to connect to the Ethereum client: %v", err)
	}
	// Instantiate the contract and display its name
	publicationRegister, err := gocontracts.NewPublicationRegister(common.HexToAddress("0x18BC6842b400b09D5CC06aBfB7522a12d8F23579"), conn)
	if err != nil {
		log.Crit("Failed to instantiate the UserContentRegister contract: %v", err)
	}
	//ret, err := userContentRegister.UserIndex(nil, common.HexToAddress("0xe2eb4e5418e8d1f90b474318b83034a15fae409f"))
	//if err != nil {
	//	log.Crit("Failed to retrieve username: %v", err)
	//}
	//fmt.Println("User name: ", ret.UserName)
	var blockNumber uint64 = 2530622

	ch := make(chan *gocontracts.PublicationRegisterStoreData)
	opts := &bind.WatchOpts{}
	opts.Start = &blockNumber
	storeDataSub, err := publicationRegister.WatchStoreData(opts, ch)
	if err != nil {
		log.Crit("Failed WatchStoreData: %v", err)
	}

	go func() {
		defer storeDataSub.Unsubscribe()

		for {
			//received an Ethereum Contract Event
			var newEvent *gocontracts.PublicationRegisterStoreData = <-ch
			fmt.Println(newEvent.Data)
			go func() {
				ctx, cancel := context.WithTimeout(context.Background(), 25*time.Second)
				defer cancel()
				_, err := ipfs.Content(ctx, newEvent.Data)
				if err == nil {
					fmt.Println("this shit may have actually worked")
				} else {
					fmt.Println("Error ipfs.Content command: " + err.Error())
				}
			}()

		}
	}()

	//Monitor for interrupts and terminate cleanly
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt)
	defer signal.Stop(sigc)
	<-sigc
}
