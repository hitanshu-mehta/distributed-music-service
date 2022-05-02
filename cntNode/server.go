package content_node

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

const tmpFilesDirName = "tmp-files"

type ErrorInfo struct {
	Msg string
}

type CntNode struct {
	srv *http.Server

	ipfs *IpfsService

	addr            string
	writeTimeout    time.Duration
	readTimeout     time.Duration
	idleTimeout     time.Duration
	gracefulTimeout time.Duration

	logger log.Logger
}

// NewCntNode creates new instance of the content node.
func NewCntNode(addr string, writeTimeout,
	readTimeout, idleTimeout, gracefulTimeout time.Duration,
	logger log.Logger) *CntNode {

	return &CntNode{
		addr:            addr,
		writeTimeout:    writeTimeout,
		readTimeout:     readTimeout,
		idleTimeout:     idleTimeout,
		gracefulTimeout: gracefulTimeout,

		ipfs: NewIpfsService(context.Background(), &logger),

		logger: logger,
	}
}

// ListenAndServe initialize the rounter and starts the http server.
func (cn *CntNode) ListenAndServe() {
	r := mux.NewRouter()
	r.HandleFunc("/upload", cn.uploadHandler).Methods("POST")
	r.HandleFunc("/download/{cid}", cn.downloadHandler).Methods("GET")

	cn.srv = &http.Server{
		Handler: r,

		Addr:         cn.addr,
		WriteTimeout: cn.writeTimeout,
		ReadTimeout:  cn.readTimeout,
	}

	cn.ipfs.Start()

	go func() {
		cn.logger.Println("content node is ready to serve request.")
		if err := cn.srv.ListenAndServe(); err != nil {
			cn.logger.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), cn.gracefulTimeout)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	cn.srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	cn.logger.Println("shutting down")
	os.Exit(0)
}

type UploadFileResponse struct {
	Cid CID
}

func (cn *CntNode) uploadHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)

	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, header, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&ErrorInfo{
			Msg: "Failed to parse the file.",
		})

		return
	}
	defer file.Close()

	// Create a temporary file within our temp-files directory that follows
	// a particular naming pattern
	tempFile, err := ioutil.TempFile(cn.getOrCreateTmpDir(tmpFilesDirName), "*-"+header.Filename)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&ErrorInfo{
			Msg: "Failed to upload file.",
		})
		return
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	tempFile.Write(fileBytes)

	cid, err := cn.ipfs.UploadFile(tempFile.Name())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&ErrorInfo{
			Msg: "Failed to upload file to IPFS node.",
		})

		return
	}

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(&UploadFileResponse{
		Cid: cid,
	})

	if err = os.Remove(tempFile.Name()); err != nil {
		cn.logger.Print("Failed to remove tmp file.")
	}

	cn.logger.Print("Uploaded file to IPFS node successfully.")
}

// getOrCreateTmpDir returns the path of temparory directory to store the upload files.
// if directory does not exists then new one is created.
func (cn *CntNode) getOrCreateTmpDir(name string) string {
	if err := os.Mkdir(name, os.ModePerm); err != nil {
		cn.logger.Print("Temparory directory already exits.")
	}

	return name
}

func (cn *CntNode) downloadHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Downloading file")
}
