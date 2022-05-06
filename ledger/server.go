package ledger

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/mux"
)

const smartContractAddress = "app/contracts"

type ErrorInfo struct {
	Msg string
}

type LedgerNode struct {
	srv *http.Server

	addr            string
	writeTimeout    time.Duration
	readTimeout     time.Duration
	idleTimeout     time.Duration
	gracefulTimeout time.Duration

	ethRPCUrl string
	ethClient *EthClient

	logger log.Logger
}

// NewLedgerNode creates new instance of the ledger node.
func NewLedgerNode(addr string,
	writeTimeout,
	readTimeout, idleTimeout, gracefulTimeout time.Duration,
	ethRPCUrl string,
	logger log.Logger) *LedgerNode {
	return &LedgerNode{
		addr:            addr,
		writeTimeout:    writeTimeout,
		readTimeout:     readTimeout,
		idleTimeout:     idleTimeout,
		gracefulTimeout: gracefulTimeout,

		ethRPCUrl: ethRPCUrl,

		logger: logger,
	}
}

// ListenAndServe initialize the rounter and starts the http server.
func (l *LedgerNode) ListenAndServe() {
	r := mux.NewRouter()
	r.HandleFunc("/initialize", l.initialize).Methods("POST")
	r.HandleFunc("/song/add", l.addSong).Methods("POST")
	r.HandleFunc("/song/cids", l.getAllCids).Methods("GET")
	r.HandleFunc("/song/{cid}", l.getSong).Methods("GET")

	r.HandleFunc("/token/buy", l.buyToken).Methods("POST")
	r.HandleFunc("/token/transfer", l.tranfer).Methods("POST")

	l.srv = &http.Server{
		Handler: r,

		Addr:         l.addr,
		WriteTimeout: l.writeTimeout,
		ReadTimeout:  l.readTimeout,
	}

	go func() {
		l.logger.Println("starting ledger node.")
		if err := l.start(); err != nil {
			l.logger.Panic("failed to start ledger node.", err)
		}

		l.logger.Println("ledger node is ready to serve request.")
		if err := l.srv.ListenAndServe(); err != nil {
			l.logger.Panic("failed to start ledger node.", err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), l.gracefulTimeout)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	l.srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	l.logger.Println("shutting down")
	os.Exit(0)
}

func (l *LedgerNode) start() error {
	conn, err := ethclient.Dial(l.ethRPCUrl)
	if err != nil {
		l.logger.Printf("failed to create eth client for url: %v", l.ethRPCUrl)
		return err
	}

	l.ethClient = NewEthClient(conn, l.logger)
	return nil
}

type PublisherDetails struct {
	AccountAddress string `json:"address"`
	GasLimit       uint64 `json:"gas_limit"`
	GasPrice       int64  `json:"gas_price"`
	Value          int64  `json:"value" default:"0"`
}

type AddSongRequest struct {
	CID string `json:"cid"`

	PublisherDetails PublisherDetails `json:"publisher_details"`

	Artists     []string `json:"artists"`
	SongName    string   `json:"song_name"`
	Description string   `json:"description"`
}

type InitializeRequest struct {
	Address string `json:"address"`

	GasLimit uint64 `json:"gas_limit"`
	GasPrice int64  `json:"gas_price"`
}

// initialize the system. It will deploy token and publish song smart contracts.
// Addresses of deployed smart contracts will be stored locally.
func (l *LedgerNode) initialize(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		l.logger.Println("failed to parse body.")
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}

	var req InitializeRequest
	if err = json.Unmarshal(body, &req); err != nil {
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}

	if err = l.ethClient.initialize(req.Address, req.GasLimit, req.GasPrice); err != nil {
		http.Error(w, "failed to deploy smart contracts", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

func (l *LedgerNode) addSong(w http.ResponseWriter, r *http.Request) {
	var req AddSongRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&ErrorInfo{
			Msg: "Failed to decode request body.",
		})
		return
	}

	l.ethClient.addSong(&req.PublisherDetails, &song{
		cid:         req.CID,
		artists:     req.Artists,
		name:        req.SongName,
		description: req.Description,
	})

	w.WriteHeader(http.StatusCreated)
}

type GetSongDetailsResponse struct {
	Cid string

	Artists     []string
	Name        string
	Description string
}

func (l *LedgerNode) getSong(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	c, ok := vars["cid"]
	if !ok {
		l.logger.Print("CID is not present in the get song request.")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&ErrorInfo{
			Msg: "CID is not present in the get request.",
		})
		return
	}

	song, err := l.ethClient.getSong(c)
	if err != nil {
		l.logger.Printf("failed to fetch song details. %v", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&ErrorInfo{
			Msg: "Failed to fetch song details",
		})
		return
	}

	resp := &GetSongDetailsResponse{
		Cid:         song.cid,
		Name:        song.name,
		Description: song.description,
		Artists:     song.artists,
	}

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(resp)
}

type GetAllCIDsResponse struct {
	CIDS []string `json:"cids"`
}

func (l *LedgerNode) getAllCids(w http.ResponseWriter, r *http.Request) {
	cids, err := l.ethClient.getAllCids()
	if err != nil {
		l.logger.Printf("failed to fetch all cids. %v", err)
		http.Error(w, "failed to fetch all cids.", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(GetAllCIDsResponse{
		CIDS: cids,
	})
}

type BuyTokenRequest struct {
	PublisherDetails PublisherDetails `json:"publisher_details"`
}

type BuyTokenResponse struct {
	TotalBalance int64 `json:"total_balance"`
}

func (l *LedgerNode) buyToken(w http.ResponseWriter, r *http.Request) {
	var req BuyTokenRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&ErrorInfo{
			Msg: "Failed to decode request body.",
		})
		return
	}

	balance, err := l.ethClient.buyToken(&req.PublisherDetails)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&ErrorInfo{
			Msg: "Failed to decode request body.",
		})
		return
	}

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(&BuyTokenResponse{
		TotalBalance: balance.Int64(),
	})
}

func (l *LedgerNode) tranfer(w http.ResponseWriter, r *http.Request) {

}
