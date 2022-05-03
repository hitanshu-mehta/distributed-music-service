package ledger

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

type LedgerNode struct {
	srv *http.Server

	addr            string
	writeTimeout    time.Duration
	readTimeout     time.Duration
	idleTimeout     time.Duration
	gracefulTimeout time.Duration

	logger log.Logger
}

// NewLedgerNode creates new instance of the ledger node.
func NewLedgerNode(addr string,
	writeTimeout,
	readTimeout, idleTimeout, gracefulTimeout time.Duration,
	logger log.Logger) *LedgerNode {
	return &LedgerNode{
		addr:            addr,
		writeTimeout:    writeTimeout,
		readTimeout:     readTimeout,
		idleTimeout:     idleTimeout,
		gracefulTimeout: gracefulTimeout,

		logger: logger,
	}
}

// ListenAndServe initialize the rounter and starts the http server.
func (l *LedgerNode) ListenAndServe() {
	r := mux.NewRouter()
	r.HandleFunc("/contruct", l.construct).Methods("GET")

	l.srv = &http.Server{
		Handler: r,

		Addr:         l.addr,
		WriteTimeout: l.writeTimeout,
		ReadTimeout:  l.readTimeout,
	}

	go func() {
		l.logger.Println("starting content node.")

		l.logger.Println("content node is ready to serve request.")
		if err := l.srv.ListenAndServe(); err != nil {
			l.logger.Println(err)
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

func (l *LedgerNode) construct(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}
