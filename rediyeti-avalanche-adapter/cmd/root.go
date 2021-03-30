package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"gitlab.red.dev/redi-yeti/rediyeti-avalanche-adapter/internal/adapter"

	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "rediyeti-avalanche-adapter",
	Short: "The chainlink node adapter",
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the server",
	Run: func(cmd *cobra.Command, args []string) {
		a := getAdapter()
		a.Serve()
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(serveCmd)
}

func getAdapter() *adapter.Adapter {
	cfg := adapter.Config{
		Port:    viper.GetInt("port"),
		ApiHost: viper.GetString("api_host"),
		ApiPort: viper.GetInt("api_port"),
	}
	a := adapter.NewAdapter(cfg)
	return a
}

func initConfig() {
	viper.SetDefault("port", 8080)
	viper.SetDefault("api_host", "http://localhost")
	viper.SetDefault("api_port", 9650)

	viper.AddConfigPath(".")
	viper.SetConfigName("config")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Failed to load config with error: %s", err)
	}
	fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
}
