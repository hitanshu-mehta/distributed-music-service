package ledger

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// EthClient is the wrapper around rpc client provided by geth.
type EthClient struct {
	client *ethclient.Client

	logger log.Logger
}

// NewEthClient creates the new instance of eth client.
func NewEthClient(
	client *ethclient.Client,
	logger log.Logger,
) *EthClient {
	return &EthClient{
		client: client,
		logger: logger,
	}
}

// Song contains details of the song.
type song struct {
	cid string

	artists     []string
	name        string
	description string
}

func (c *EthClient) addSong(publisher *PublisherDetails, song *song) {

}

func (c *EthClient) getAccountAuth(accountAddress string, gasLimit uint64, gasPrice int64) (*bind.TransactOpts, error) {

	privateKey, err := crypto.HexToECDSA(accountAddress)
	if err != nil {
		return nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("invalid key")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	//fetch the last use nonce of account
	nonce, err := c.client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, err
	}
	fmt.Println("nounce=", nonce)
	chainID, err := c.client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)           // in wei
	auth.GasLimit = uint64(gasLimit)     // in wei
	auth.GasPrice = big.NewInt(gasPrice) // in wei

	return auth, nil
}
