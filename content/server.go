package content

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/ipfs/go-cid"
)

const tmpUploadFilesDirName = "tmp-upload-files"

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

	isTemporary bool

	logger log.Logger
}

// NewCntNode creates new instance of the content node.
func NewCntNode(addr string,
	writeTimeout,
	readTimeout, idleTimeout, gracefulTimeout time.Duration,
	temparoryNode bool,
	logger log.Logger) *CntNode {
	return &CntNode{
		addr:            addr,
		writeTimeout:    writeTimeout,
		readTimeout:     readTimeout,
		idleTimeout:     idleTimeout,
		gracefulTimeout: gracefulTimeout,

		isTemporary: temparoryNode,

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

	go func() {
		cn.logger.Println("starting content node.")
		// Start IPFS service.
		cn.ipfs = NewIpfsService(context.Background(), &cn.logger, cn.isTemporary)
		if err := cn.ipfs.Start(); err != nil {
			cn.logger.Panic(err)
		}

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
	CID string
}

// uploadHandler reads the multipart file from the body and creates temparory file in local filesytem
// Local file is then uploaded to the ipfs network. Ultimately, local file is deleted.
func (cn *CntNode) uploadHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)

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
	tempFile, err := ioutil.TempFile(cn.getOrCreateTmpDir(tmpUploadFilesDirName), "*-"+header.Filename)
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
		CID: cid.String(),
	})

	defer func() {
		if err = os.Remove(tempFile.Name()); err != nil {
			cn.logger.Print("Failed to remove tmp file.")
		}
	}()

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

// downloadHandler parses and validate CID. If CID is valid then file is downloaded and stored locally before
// sending it in response. Ultimately, this local file is deleted.
func (cn *CntNode) downloadHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	c, ok := vars["cid"]
	if !ok {
		cn.logger.Print("CID is not present in the download request.")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&ErrorInfo{
			Msg: "CID is not present in the download request.",
		})
		return
	}

	cid, err := cid.Decode(c)
	if err != nil {
		cn.logger.Print("Failed to decode the CID.")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&ErrorInfo{
			Msg: "Failed to decode the CID.",
		})

		return
	}

	downloadedFilePath, err := cn.ipfs.DownloadFile(cid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&ErrorInfo{
			Msg: "Failed to download file from IPFS network.",
		})

		return
	}

	w.Header().Set("Content-Disposition", "attachment; filename="+strconv.Quote(cid.String()))
	w.Header().Set("Content-Type", "application/octet-stream")
	http.ServeFile(w, r, downloadedFilePath.String())

	// Delete the temporary file.
	defer func() {
		if err := os.Remove(downloadedFilePath.String()); err != nil {
			cn.logger.Printf("Failed to remvoe tempory downloaded file. file path: %v", downloadedFilePath.String())
		}
	}()

	cn.logger.Printf("Successfully downloaded the file having cid %v", cid)
}
