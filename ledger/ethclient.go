package ledger

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/hitanshu-mehta/distributed-music-service/ledger/contracts/publishsongapi"
	"github.com/hitanshu-mehta/distributed-music-service/ledger/contracts/tokenapi"
	"github.com/hitanshu-mehta/distributed-music-service/ledger/contracts/tokencrowdsaleapi"
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
	conn, err := publishsongapi.NewPublishsongapi(common.HexToAddress(smtData.PublishSongContractAddr), c.client)
	if err != nil {
		c.logger.Print("failed to create new api.", err)
		return errors.New("failed to add song")
	}

	auth, err := c.getAccountAuth(publisher.AccountAddress, publisher.GasLimit, publisher.GasPrice, 0)
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
	TokenContractAddr       string
	CrowdSaleContractAddr   string
}

func (c *EthClient) initialize(addr string, gasLimit uint64, gasPrice int64) error {
	auth, err := c.getAccountAuth(addr, gasLimit, gasPrice, 0)
	if err != nil {
		c.logger.Println("failed to bind address.")
		return err
	}

	songContractAddr, _, _, err := publishsongapi.DeployPublishsongapi(auth, c.client)
	if err != nil {
		c.logger.Printf("failed to deploy publish song contract. %v", err)
		return err
	}

	auth, err = c.getAccountAuth(addr, gasLimit, gasPrice, 0)
	if err != nil {
		c.logger.Println("failed to bind address.")
		return err
	}
	c.logger.Printf("publish contract successfully deployed at %v", songContractAddr.Hex())

	tokenContractAddr, _, _, err := tokenapi.DeployTokenapi(auth, c.client)
	if err != nil {
		c.logger.Printf("failed to deploy token contract. %v", err)
		return err
	}
	c.logger.Printf("token contract successfully deployed at %v", tokenContractAddr.Hex())

	auth, err = c.getAccountAuth(addr, gasLimit, gasPrice, 0)
	if err != nil {
		c.logger.Println("failed to bind address.")
		return err
	}

	crowdSaleContractAddr, _, _, err := tokencrowdsaleapi.DeployTokencrowdsaleapi(auth, c.client, big.NewInt(1), auth.From, tokenContractAddr)
	if err != nil {
		c.logger.Printf("failed to deploy token crowdsale contract. %v", err)
		return err
	}
	c.logger.Printf("token crowd sale contract successfully deployed at %v", crowdSaleContractAddr.Hex())

	if err = c.giveOwnershipToCrowdSale(addr, tokenContractAddr, crowdSaleContractAddr, gasLimit, gasPrice); err != nil {
		c.logger.Printf("failed to transfer minting ownership to crowdsale. %v", err)
		return err
	}

	file, err := os.Create(smartContractAddress)
	if err != nil {
		c.logger.Println("failed to create file.")
		return err
	}
	defer file.Close()

	data := smartContractData{
		PublishSongContractAddr: songContractAddr.String(),
		TokenContractAddr:       tokenContractAddr.String(),
		CrowdSaleContractAddr:   crowdSaleContractAddr.String(),
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
	conn, err := publishsongapi.NewPublishsongapi(common.HexToAddress(smtData.PublishSongContractAddr), c.client)
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
	conn, err := publishsongapi.NewPublishsongapi(common.HexToAddress(smtData.PublishSongContractAddr), c.client)
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

func (c *EthClient) buyToken(publisher *PublisherDetails) (*big.Int, error) {
	failedOpErr := errors.New("failed to buy tokens")

	contractData, err := c.readSmartContractFileData()
	if err != nil {
		c.logger.Print("failed to read smart contract file data.", err)
		return nil, failedOpErr
	}

	csconn, err := tokencrowdsaleapi.NewTokencrowdsaleapi(common.HexToAddress(contractData.CrowdSaleContractAddr), c.client)
	if err != nil {
		c.logger.Print("failed to create new token contract api.", err)
		return nil, failedOpErr
	}

	auth, err := c.getAccountAuth(publisher.AccountAddress, publisher.GasLimit, publisher.GasPrice, publisher.Value)
	if err != nil {
		c.logger.Printf("failed to authorize transactor. %v", err)
		return nil, failedOpErr
	}

	_, err = csconn.BuyTokens(auth, common.HexToAddress(publisher.AccountAddress))
	if err != nil {
		c.logger.Printf("failed to buy token. %v", err)
		return nil, failedOpErr
	}

	tAddr, err := csconn.Token(&bind.CallOpts{})
	if err != nil {
		return nil, failedOpErr
	}

	tconn, err := tokenapi.NewTokenapi(tAddr, c.client)
	if err != nil {
		c.logger.Printf("failed to create new token contract api. %v", err)
		return nil, failedOpErr
	}

	amt, err := tconn.BalanceOf(&bind.CallOpts{}, common.HexToAddress(publisher.AccountAddress))
	if err != nil {
		c.logger.Printf("failed to fetch balace. %v", err)
		return nil, failedOpErr
	}

	c.logger.Print("successfully bought the token.")
	return amt, nil
}

func (c *EthClient) getAccountAuth(accountAddress string, gasLimit uint64, gasPrice, value int64) (*bind.TransactOpts, error) {
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
	chainID, err := c.client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(value)       // in wei
	auth.GasLimit = uint64(gasLimit)     // in wei
	auth.GasPrice = big.NewInt(gasPrice) // in wei

	return auth, nil
}

func (c *EthClient) giveOwnershipToCrowdSale(addr string, tokenAddr, crowdSaleAddr common.Address, gasLimit uint64, gasPrice int64) error {
	tconn, err := tokenapi.NewTokenapi(tokenAddr, c.client)
	if err != nil {
		c.logger.Print("failed to create new token contract api.", err)
		return err
	}

	auth, err := c.getAccountAuth(addr, gasLimit, gasPrice, 0)
	if err != nil {
		c.logger.Printf("failed to authorize transactor. %v", err)
		return err
	}

	_, err = tconn.AddMinter(auth, crowdSaleAddr)
	if err != nil {
		return err
	}

	return nil
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
