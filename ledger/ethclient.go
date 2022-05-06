package ledger

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/hitanshu-mehta/distributed-music-service/ledger/contracts/api"
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

func (c *EthClient) addSong(publisher *PublisherDetails, song *song) error {
	smtData, err := c.readSmartContractFileData()
	if err != nil {
		c.logger.Print("failed to read smart contract file data.", err)
		return errors.New("failed to add song")
	}

	//creating api object to intract with smart contract function
	conn, err := api.NewApi(common.HexToAddress(smtData.PublishSongContractAddr), c.client)
	if err != nil {
		c.logger.Print("failed to create new api.", err)
		return errors.New("failed to add song")
	}

	auth, err := c.getAccountAuth(publisher.AccontAddress, publisher.GasLimit, publisher.GasPrice)
	if err != nil {
		c.logger.Printf("failed to authorize transactor. %v", err)
		return errors.New("failed to add song")
	}

	txn, err := conn.AddSong(auth, song.cid, song.name, song.artists, song.description)
	if err != nil {
		c.logger.Print("failed to publish song.", err)
		return errors.New("failed to add song")
	}

	c.logger.Print("successfully added new song.", txn)
	return nil
}

type smartContractData struct {
	PublishSongContractAddr string
}

func (c *EthClient) initialize(addr string, gasLimit uint64, gasPrice int64) error {
	auth, err := c.getAccountAuth(addr, gasLimit, gasPrice)
	if err != nil {
		c.logger.Println("failed to bind address.")
		return err
	}

	contractAddr, _, _, err := api.DeployApi(auth, c.client)
	if err != nil {
		c.logger.Println("failed to deploy contract.")
		return err
	}

	c.logger.Printf("contract successfully deployed at %v", contractAddr.Hex())

	file, err := os.Create(smartContractAddress)
	if err != nil {
		c.logger.Println("failed to create file.")
		return err
	}
	defer file.Close()

	data := smartContractData{
		PublishSongContractAddr: contractAddr.String(),
	}

	marshalledData, err := json.Marshal(data)
	if err != nil {
		c.logger.Println("failed to marshall data.", err)
		return err
	}

	file.Write(marshalledData)
	return nil
}

func (c *EthClient) getSong(cid string) (*song, error) {
	smtData, err := c.readSmartContractFileData()
	if err != nil {
		c.logger.Print("failed to read smart contract file data.", err)
		return nil, errors.New("failed to get song")
	}

	//creating api object to intract with smart contract function
	conn, err := api.NewApi(common.HexToAddress(smtData.PublishSongContractAddr), c.client)
	if err != nil {
		c.logger.Print("failed to create new api.", err)
		return nil, errors.New("failed to get song")
	}

	publishedSong, err := conn.GetSong(&bind.CallOpts{}, cid)
	if err != nil {
		c.logger.Print("failed to get song.", err)
		return nil, errors.New("failed to get song")
	}

	c.logger.Print("successfully fetched song.")
	return &song{
		cid:         publishedSong.Cid,
		artists:     publishedSong.Artists,
		name:        publishedSong.Name,
		description: publishedSong.Description,
	}, nil
}

func (c *EthClient) getAllCids() ([]string, error) {
	smtData, err := c.readSmartContractFileData()
	if err != nil {
		c.logger.Print("failed to read smart contract file data.", err)
		return nil, errors.New("failed to get song cids")
	}

	//creating api object to intract with smart contract function
	conn, err := api.NewApi(common.HexToAddress(smtData.PublishSongContractAddr), c.client)
	if err != nil {
		c.logger.Print("failed to create new api.", err)
		return nil, errors.New("failed to get song cids")
	}

	cids, err := conn.GetAllCid(&bind.CallOpts{})
	if err != nil {
		c.logger.Print("failed to get song cids", err)
		return nil, errors.New("failed to get song cids")
	}

	c.logger.Print("successfully fetched all cids")
	return cids, nil

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

func (c *EthClient) readSmartContractFileData() (*smartContractData, error) {
	data, err := os.ReadFile(smartContractAddress)
	if err != nil {
		return nil, err
	}

	var contractData smartContractData
	if err = json.Unmarshal(data, &contractData); err != nil {
		return nil, err
	}

	return &contractData, nil
}
