package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	serverCmd = &cobra.Command{
		Use:   "server",
		Short: "Start integrated server",
		Run:   serverAction,
		Args:  cobra.NoArgs,
	}

	defaultMetricsAddr = "0.0.0.0:8081"
	defaultServerAddr  = "0.0.0.0:8080"
	defaultServerCert  = ""
	defaultServerKey   = ""
)

func init() {
	rootCmd.AddCommand(serverCmd)

	//region Metrics
	serverCmd.PersistentFlags().String("metrics-addr", defaultMetricsAddr, "Address to bind the metrics")
	viper.SetDefault("metrics.addr", defaultMetricsAddr)
	_ = viper.BindPFlag("metrics.addr", serverCmd.PersistentFlags().Lookup("metrics-addr"))

	serverCmd.PersistentFlags().String("metrics-token", "", "Token to make metrics secure")
	viper.SetDefault("metrics.token", "")
	_ = viper.BindPFlag("metrics.token", serverCmd.PersistentFlags().Lookup("metrics-token"))
	//endregion

	//region Server
	serverCmd.PersistentFlags().String("server-addr", defaultServerAddr, "Address to bind the server")
	viper.SetDefault("server.addr", defaultServerAddr)
	_ = viper.BindPFlag("server.addr", serverCmd.PersistentFlags().Lookup("server-addr"))

	serverCmd.PersistentFlags().String("server-cert", defaultServerCert, "Path to cert for SSL encryption")
	viper.SetDefault("server.cert", defaultServerCert)
	_ = viper.BindPFlag("server.cert", serverCmd.PersistentFlags().Lookup("server-cert"))

	serverCmd.PersistentFlags().String("server-key", defaultServerKey, "Path to key for SSL encryption")
	viper.SetDefault("server.key", defaultServerKey)
	_ = viper.BindPFlag("server.key", serverCmd.PersistentFlags().Lookup("server-key"))
	//endregion
}

func serverAction(_ *cobra.Command, _ []string) {

}
