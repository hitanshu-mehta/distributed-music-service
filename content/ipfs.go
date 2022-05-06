package content

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	cid "github.com/ipfs/go-cid"
	config "github.com/ipfs/go-ipfs-config"
	files "github.com/ipfs/go-ipfs-files"
	"github.com/ipfs/go-ipfs/core"
	"github.com/ipfs/go-ipfs/core/coreapi"
	"github.com/ipfs/go-ipfs/core/node/libp2p"
	"github.com/ipfs/go-ipfs/plugin/loader"
	"github.com/ipfs/go-ipfs/repo/fsrepo"
	icore "github.com/ipfs/interface-go-ipfs-core"
	"github.com/ipfs/interface-go-ipfs-core/path"
)

const tmpDownloadFilesDirName = "app/download"

// IpfsService interacts with ipfs node.
type IpfsService struct {
	ipfs icore.CoreAPI

	// Create the temporary IPFS node. Data directory of this node will be cleaned up after the session ends.
	isTemporary bool

	ctx context.Context

	logger *log.Logger
}

// NewIpfsService creates the new instance of ipfs service.
func NewIpfsService(ctx context.Context, logger *log.Logger, isTemporary bool) *IpfsService {
	return &IpfsService{
		ctx:         ctx,
		logger:      logger,
		isTemporary: isTemporary,
	}
}

// Start setup the IPFS repo and start the ipfs node.
func (ipfsService *IpfsService) Start() error {
	var err error

	if ipfsService.isTemporary {
		ipfsService.ipfs, err = ipfsService.spawnEphemeral(ipfsService.ctx)
		if err != nil {
			return err
		}
	} else {
		ipfsService.ipfs, err = ipfsService.spawnDefault(ipfsService.ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

// UploadFile upload the file to ipfs node and returns the CID.
func (ipfsService *IpfsService) UploadFile(inputFilePath string) (*cid.Cid, error) {
	someFile, err := ipfsService.getUnixfsNode("./" + inputFilePath)
	if err != nil {
		ipfsService.logger.Print("Could not get the file", err)
		return nil, err
	}

	cidFile, err := ipfsService.ipfs.Unixfs().Add(ipfsService.ctx, someFile)
	if err != nil {
		ipfsService.logger.Print("Failed to upload the file to IPFS.", err)
		return nil, err
	}

	ipfsService.logger.Printf("Successfully uploaded the file to IPFS. cid: %v", cidFile.Cid())

	CID := cidFile.Cid()
	return &CID, nil
}

// DownloadFile donwload the file having given cid, stores the file in local filesystem and
// returns the path to the downloaded file.
func (ipfsService *IpfsService) DownloadFile(cid cid.Cid) (path.Path, error) {
	cidFile := path.IpfsPath(cid)

	outputBasePath := ipfsService.getOrCreateTmpDir(tmpDownloadFilesDirName)
	outputPathFile := path.Join(path.New(outputBasePath), cidFile.Cid().String())

	rootNodeFile, err := ipfsService.ipfs.Unixfs().Get(ipfsService.ctx, cidFile)
	if err != nil {
		ipfsService.logger.Printf("Failed to download file having CID: %v", cid.String())
		return nil, err
	}

	err = files.WriteTo(rootNodeFile, outputPathFile.String())
	if err != nil {
		ipfsService.logger.Printf("Could not write out the fetched CID: %v", cid.String())
		return nil, err
	}

	return outputPathFile, nil
}

// getOrCreateTmpDir returns the path of temparory directory to store the upload files.
// if directory does not exists then new one is created.
func (ipfsService *IpfsService) getOrCreateTmpDir(name string) string {
	if err := os.Mkdir(name, os.ModePerm); err != nil {
		ipfsService.logger.Print("Temparory directory already exits.")
	}

	return name
}

func (ipfsService *IpfsService) getUnixfsNode(path string) (files.Node, error) {
	st, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	f, err := files.NewSerialFile(path, false, st)
	if err != nil {
		return nil, err
	}

	return f, nil
}

// Spawns a node on the default repo location, if the repo exists
func (ipfsService *IpfsService) spawnDefault(ctx context.Context) (icore.CoreAPI, error) {
	defaultPath, err := config.PathRoot()
	if err != nil {
		// shouldn't be possible
		return nil, err
	}

	if err := ipfsService.setupPlugins(defaultPath); err != nil {
		return nil, err
	}

	return ipfsService.createNode(ctx, defaultPath)
}

// Spawns a node to be used just for this run (i.e. creates a tmp repo)
func (ipfsService *IpfsService) spawnEphemeral(ctx context.Context) (icore.CoreAPI, error) {
	if err := ipfsService.setupPlugins(""); err != nil {
		return nil, err
	}

	// Create a Temporary Repo
	repoPath, err := ipfsService.createTempRepo()
	if err != nil {
		return nil, fmt.Errorf("failed to create temp repo: %s", err)
	}

	// Spawning an ephemeral IPFS node
	return ipfsService.createNode(ctx, repoPath)
}

func (ipfsService *IpfsService) setupPlugins(externalPluginsPath string) error {
	// Load any external plugins if available on externalPluginsPath
	plugins, err := loader.NewPluginLoader(filepath.Join(externalPluginsPath, "plugins"))
	if err != nil {
		return fmt.Errorf("error loading plugins: %s", err)
	}

	// Load preloaded and external plugins
	if err := plugins.Initialize(); err != nil {
		return fmt.Errorf("error initializing plugins: %s", err)
	}

	if err := plugins.Inject(); err != nil {
		return fmt.Errorf("error initializing plugins: %s", err)
	}

	ipfsService.logger.Println("IPFS plugins loaded.")

	return nil
}

func (ipfsService *IpfsService) createTempRepo() (string, error) {
	repoPath, err := ioutil.TempDir("", "ipfs-shell")
	if err != nil {
		return "", fmt.Errorf("failed to get temp dir: %s", err)
	}

	// Create a config with default options and a 2048 bit key
	cfg, err := config.Init(ioutil.Discard, 2048)
	if err != nil {
		return "", err
	}

	// Create the repo with the config
	err = fsrepo.Init(repoPath, cfg)
	if err != nil {
		return "", fmt.Errorf("failed to init ephemeral node: %s", err)
	}

	return repoPath, nil
}

// Creates an IPFS node and returns its coreAPI
func (ipfsService *IpfsService) createNode(ctx context.Context, repoPath string) (icore.CoreAPI, error) {
	// Open the repo
	repo, err := fsrepo.Open(repoPath)
	if err != nil {
		return nil, err
	}

	// Construct the node
	nodeOptions := &core.BuildCfg{
		Online:  true,
		Routing: libp2p.DHTOption, // This option sets the node to be a full DHT node (both fetching and storing DHT Records)
		// Routing: libp2p.DHTClientOption, // This option sets the node to be a client DHT node (only fetching records)
		Repo: repo,
	}

	node, err := core.NewNode(ctx, nodeOptions)
	if err != nil {
		return nil, err
	}

	// Attach the Core API to the constructed node
	return coreapi.NewCoreAPI(node)
}
