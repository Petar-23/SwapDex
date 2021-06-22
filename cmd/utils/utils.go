package utils

import (
	"github.com/tomochain/tomochain/eth"
	"github.com/tomochain/tomochain/eth/downloader"
	"github.com/tomochain/tomochain/ethstats"
	"github.com/tomochain/tomochain/les"
	"github.com/tomochain/tomochain/node"
	"github.com/tomochain/tomochain/sdxx"
	"github.com/tomochain/tomochain/sdxxlending"
	whisper "github.com/tomochain/tomochain/whisper/whisperv6"
)

// RegisterEthService adds an Ethereum client to the stack.
func RegisterEthService(stack *node.Node, cfg *eth.Config) {
	var err error
	if cfg.SyncMode == downloader.LightSync {
		err = stack.Register(func(ctx *node.ServiceContext) (node.Service, error) {
			return les.New(ctx, cfg)
		})
	} else {
		err = stack.Register(func(ctx *node.ServiceContext) (node.Service, error) {
			var sdxXServ *sdxx.SdxX
			ctx.Service(&sdxXServ)
			var lendingServ *sdxxlending.Lending
			ctx.Service(&lendingServ)
			fullNode, err := eth.New(ctx, cfg, sdxXServ, lendingServ)
			if fullNode != nil && cfg.LightServ > 0 {
				ls, _ := les.NewLesServer(fullNode, cfg)
				fullNode.AddLesServer(ls)
			}
			return fullNode, err
		})
	}
	if err != nil {
		Fatalf("Failed to register the Ethereum service: %v", err)
	}
}

// RegisterShhService configures Whisper and adds it to the given node.
func RegisterShhService(stack *node.Node, cfg *whisper.Config) {
	if err := stack.Register(func(n *node.ServiceContext) (node.Service, error) {
		return whisper.New(cfg), nil
	}); err != nil {
		Fatalf("Failed to register the Whisper service: %v", err)
	}
}

// RegisterEthStatsService configures the Ethereum Stats daemon and adds it to
// th egiven node.
func RegisterEthStatsService(stack *node.Node, url string) {
	if err := stack.Register(func(ctx *node.ServiceContext) (node.Service, error) {
		// Retrieve both eth and les services
		var ethServ *eth.Ethereum
		ctx.Service(&ethServ)

		var lesServ *les.LightEthereum
		ctx.Service(&lesServ)

		return ethstats.New(url, ethServ, lesServ)
	}); err != nil {
		Fatalf("Failed to register the Ethereum Stats service: %v", err)
	}
}

func RegisterSdxXService(stack *node.Node, cfg *sdxx.Config) {
	sdxX := sdxx.New(cfg)
	if err := stack.Register(func(n *node.ServiceContext) (node.Service, error) {
		return sdxX, nil
	}); err != nil {
		Fatalf("Failed to register the SdxX service: %v", err)
	}

	// register sdxxlending service
	if err := stack.Register(func(n *node.ServiceContext) (node.Service, error) {
		return sdxxlending.New(sdxX), nil
	}); err != nil {
		Fatalf("Failed to register the SdxXLending service: %v", err)
	}
}
