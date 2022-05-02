package main

import (
	"log"
	"time"

	"github.com/alecthomas/kong"

	cntNode "github.com/hitanshu-mehta/distributed-music-service/cntNode"
)

type Context struct {
	Debug bool
}

type CntNodeCmd struct {
	HttpAddr string `name:"http-addr" help:"Http address to connect with content node." default:"127.0.0.1:2000"`

	WriteTimeout    time.Duration `help:"Write timeout duration of http server." default:"15s" type:"time.Duration"`
	ReadTimeout     time.Duration `help:"Read timeout duration of http server." default:"15s" type:"time.Duration"`
	IdleTimeout     time.Duration `help:"Idle timeout duration of http server." default:"15s" type:"time.Duration"`
	GracefulTimeout time.Duration `help:"Graceful shutdown timeout duration of http server." default:"60s" type:"time.Duration"`
}

func (cntNodeCmd *CntNodeCmd) Run(ctx *Context) error {
	cntNode := cntNode.NewCntNode(cntNodeCmd.HttpAddr, cntNodeCmd.WriteTimeout, cntNodeCmd.ReadTimeout, cntNodeCmd.IdleTimeout, cntNodeCmd.GracefulTimeout, *log.Default())
	cntNode.ListenAndServe()
	return nil
}

var cli struct {
	Debug       bool       `help:"Enable debug mode."`
	ContentNode CntNodeCmd `cmd:"" help:"Content node interacts with IPFS network."`
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
