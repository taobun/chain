package yoda

import (
	"fmt"
	"os"
	"path"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	band "github.com/bandprotocol/chain/v2/app"
)

const (
	flagValidator        = "validator"
	flagLogLevel         = "log-level"
	flagExecutor         = "executor"
	flagBroadcastTimeout = "broadcast-timeout"
	flagRPCPollInterval  = "rpc-poll-interval"
	flagMaxTry           = "max-try"
	flagMaxReport        = "max-report"
)

// Config data structure for yoda daemon.
type Config struct {
	ChainID           string `mapstructure:"chain-id"`            // ChainID of the target chain
	NodeURI           string `mapstructure:"node"`                // Remote RPC URI of BandChain node to connect to
	Validator         string `mapstructure:"validator"`           // The validator address that I'm responsible for
	GasPrices         string `mapstructure:"gas-prices"`          // Gas prices of the transaction
	LogLevel          string `mapstructure:"log-level"`           // Log level of the logger
	Executor          string `mapstructure:"executor"`            // Executor name and URL (example: "Executor name:URL")
	BroadcastTimeout  string `mapstructure:"broadcast-timeout"`   // The time that Yoda will wait for tx commit
	RPCPollInterval   string `mapstructure:"rpc-poll-interval"`   // The duration of rpc poll interval
	MaxTry            uint64 `mapstructure:"max-try"`             // The maximum number of tries to submit a report transaction
	MaxReport         uint64 `mapstructure:"max-report"`          // The maximum number of reports in one transaction
	MetricsListenAddr string `mapstructure:"metrics-listen-addr"` // Address to listen on for prometheus metrics
}

// Global instances.
var (
	cfg Config
	kb  keyring.Keyring
)

func initConfig(cmd *cobra.Command) error {
	home, err := cmd.PersistentFlags().GetString(flags.FlagHome)
	if err != nil {
		return err
	}
	viper.SetConfigFile(path.Join(home, "config.yaml"))
	_ = viper.ReadInConfig() // If we fail to read config file, we'll just rely on cmd flags.
	if err := viper.Unmarshal(&cfg); err != nil {
		return err
	}
	return nil
}

func Main() {
	appConfig := sdk.GetConfig()
	band.SetBech32AddressPrefixesAndBip44CoinTypeAndSeal(appConfig)

	ctx := &Context{}
	rootCmd := &cobra.Command{
		Use:   "yoda",
		Short: "BandChain oracle daemon to subscribe and response to oracle requests",
	}

	rootCmd.AddCommand(
		configCmd(),
		keysCmd(ctx),
		runCmd(ctx),
		version.NewVersionCommand(),
	)
	rootCmd.PersistentPreRunE = func(_ *cobra.Command, _ []string) error {
		home, err := rootCmd.PersistentFlags().GetString(flags.FlagHome)
		if err != nil {
			return err
		}
		if err := os.MkdirAll(home, os.ModePerm); err != nil {
			return err
		}
		kb, err = keyring.New("band", "test", home, nil)
		if err != nil {
			return err
		}
		return initConfig(rootCmd)
	}
	rootCmd.PersistentFlags().String(flags.FlagHome, os.ExpandEnv("$HOME/.yoda"), "home directory")
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
