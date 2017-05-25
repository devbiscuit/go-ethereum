package main

import (
	"fmt"
	"os"

	"gopkg.in/urfave/cli.v1"

	ethUtils "github.com/ethereum/go-ethereum/cmd/utils"

	"github.com/tendermint/ethermint/cmd/utils"
	"github.com/tendermint/ethermint/version"
)

const (
	// Client identifier to advertise over the network
	clientIdentifier = "ethermint"
)

var (
	// The app that holds all commands and flags.
	app = ethUtils.NewApp(version.Version, "the ethermint command line interface")
	// flags that configure the go-ethereum node
	nodeFlags = []cli.Flag{
		ethUtils.DataDirFlag,
		ethUtils.KeyStoreDirFlag,
		ethUtils.NoUSBFlag,
		ethUtils.NetworkIdFlag,
		ethUtils.TestnetFlag,
		ethUtils.IdentityFlag,
		// Performance tuning
		ethUtils.CacheFlag,
		ethUtils.TrieCacheGenFlag,
		// Account settings
		ethUtils.UnlockedAccountFlag,
		ethUtils.PasswordFileFlag,
		ethUtils.VMEnableDebugFlag,
		// Logging and debug settings
		ethUtils.EthStatsURLFlag,
		ethUtils.MetricsEnabledFlag,
		ethUtils.FakePoWFlag,
		ethUtils.NoCompactionFlag,
		// Network settings
		ethUtils.MaxPeersFlag,
		ethUtils.MaxPendingPeersFlag,
		ethUtils.ListenPortFlag,
		ethUtils.BootnodesFlag,
		ethUtils.NodeKeyFileFlag,
		ethUtils.NodeKeyHexFlag,
		ethUtils.NATFlag,
		ethUtils.NoDiscoverFlag,
		ethUtils.DiscoveryV5Flag,
		ethUtils.NetrestrictFlag,
		ethUtils.WhisperEnabledFlag,
		ethUtils.JSpathFlag,
		// Gas price oracle settings
		ethUtils.GpoBlocksFlag,
		ethUtils.GpoPercentileFlag,
	}

	rpcFlags = []cli.Flag{
		ethUtils.RPCEnabledFlag,
		ethUtils.RPCListenAddrFlag,
		ethUtils.RPCPortFlag,
		ethUtils.RPCCORSDomainFlag,
		ethUtils.RPCApiFlag,
		ethUtils.IPCDisabledFlag,
		ethUtils.WSEnabledFlag,
		ethUtils.WSListenAddrFlag,
		ethUtils.WSPortFlag,
		ethUtils.WSApiFlag,
		ethUtils.WSAllowedOriginsFlag,
		ethUtils.ExecFlag,
		ethUtils.PreloadJSFlag,
	}

	// flags that configure the ABCI app
	ethermintFlags = []cli.Flag{
		utils.TendermintAddrFlag,
		utils.ABCIAddrFlag,
		utils.ABCIProtocolFlag,
	}

	debugFlags = []cli.Flag{
		utils.VerbosityFlag,
		utils.DebugFlag,
	}
)

func init() {
	app.Action = ethermintCmd
	app.HideVersion = true
	app.Commands = []cli.Command{
		{
			Action:      initCmd,
			Name:        "init",
			Usage:       "init genesis.json",
			Description: "Initialize the files",
		},
		{
			Action:      versionCmd,
			Name:        "version",
			Usage:       "",
			Description: "Print the version",
		},
	}

	app.Flags = append(app.Flags, nodeFlags...)
	app.Flags = append(app.Flags, rpcFlags...)
	app.Flags = append(app.Flags, ethermintFlags...)
	app.Flags = append(app.Flags, debugFlags...)

	app.Before = func(ctx *cli.Context) error {
		if err := utils.Setup(ctx); err != nil {
			return err
		}

		ethUtils.SetupNetwork(ctx)

		return nil
	}
}

func versionCmd(ctx *cli.Context) error {
	fmt.Println(version.Version)
	return nil
}

func main() {
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
