package ledger

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/mux"
)

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
	r.HandleFunc("/song/{cid}", l.getSong).Methods("GET")
	r.HandleFunc("/song/cids", l.getAllCids).Methods("GET")

	l.srv = &http.Server{
		Handler: r,

		Addr:         l.addr,
		WriteTimeout: l.writeTimeout,
		ReadTimeout:  l.readTimeout,
	}

	go func() {
		l.logger.Println("starting content node.")
		if err := l.start(); err != nil {
			l.logger.Panic("failed to start ledger node.", err)
		}

		l.logger.Println("content node is ready to serve request.")
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
	AccontAddress string
	GasLimit      uint64
	GasPrice      int64
}

type AddSongRequest struct {
	CID string

	PublisherDetails PublisherDetails

	Artists     []string
	SongName    string
	Description string
}

// initialize the system. It will deploy token and publish song smart contracts.
// Addresses of deployed smart contracts will be stored locally.
func (l *LedgerNode) initialize(w http.ResponseWriter, r *http.Request) {

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

}

func (l *LedgerNode) getSong(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))

	// api.DeployApi()
}

func (l *LedgerNode) getAllCids(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))

}
