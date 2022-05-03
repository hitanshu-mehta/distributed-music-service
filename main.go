package main

import (
	"log"
	"time"

	"github.com/alecthomas/kong"

	"github.com/hitanshu-mehta/distributed-music-service/content"
	"github.com/hitanshu-mehta/distributed-music-service/ledger"
)

type Context struct {
	Debug bool
}

type ContentCmd struct {
	HttpAddr string `name:"http-addr" help:"Http address to connect with content node." default:"127.0.0.1:2000"`

	WriteTimeout    time.Duration `help:"Write timeout duration of http server." default:"15s" type:"time.Duration"`
	ReadTimeout     time.Duration `help:"Read timeout duration of http server." default:"15s" type:"time.Duration"`
	IdleTimeout     time.Duration `help:"Idle timeout duration of http server." default:"15s" type:"time.Duration"`
	GracefulTimeout time.Duration `help:"Graceful shutdown timeout duration of http server." default:"15s" type:"time.Duration"`

	TemporaryNode bool `help:"Create the temporary IPFS node. Data directory of this node will be cleaned up after the session ends." default:"false"`
}

func (contentCmd *ContentCmd) Run(ctx *Context) error {
	cntNode := content.NewCntNode(
		contentCmd.HttpAddr,
		contentCmd.WriteTimeout,
		contentCmd.ReadTimeout,
		contentCmd.IdleTimeout,
		contentCmd.GracefulTimeout,
		contentCmd.TemporaryNode,
		*log.Default())
	cntNode.ListenAndServe()
	return nil
}

type LedgerCmd struct {
	HttpAddr string `name:"http-addr" help:"Http address to connect with content node." default:"127.0.0.1:3000"`

	WriteTimeout    time.Duration `help:"Write timeout duration of http server." default:"15s" type:"time.Duration"`
	ReadTimeout     time.Duration `help:"Read timeout duration of http server." default:"15s" type:"time.Duration"`
	IdleTimeout     time.Duration `help:"Idle timeout duration of http server." default:"15s" type:"time.Duration"`
	GracefulTimeout time.Duration `help:"Graceful shutdown timeout duration of http server." default:"15s" type:"time.Duration"`
}

func (ledgerCmd *LedgerCmd) Run(ctx *Context) error {
	ledgerNode := ledger.NewLedgerNode(
		ledgerCmd.HttpAddr,
		ledgerCmd.WriteTimeout,
		ledgerCmd.ReadTimeout,
		ledgerCmd.IdleTimeout,
		ledgerCmd.GracefulTimeout,
		*log.Default(),
	)
	ledgerNode.ListenAndServe()
	return nil
}

var cli struct {
	Debug       bool       `help:"Enable debug mode."`
	ContentNode ContentCmd `cmd:"" name:"content" help:"Content node interacts with IPFS network."`
	LedgerNode  LedgerCmd  `cmd:"" name:"ledger" help:"Legder node interacts with smart contracts and exposes their functions via REST API."`
}

func main() {
	ctx := kong.Parse(&cli, kong.Name("music"),
		kong.Description("A distributed music service."),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
			Summary: true,
		}))
	err := ctx.Run(&Context{Debug: cli.Debug})
	ctx.FatalIfErrorf(err)
}
