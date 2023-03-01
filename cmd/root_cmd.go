package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string
var azureCLIAuthOnly bool

//var authProvider internal.AuthProvider

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:           "prp-cli",
	Short:         "prp CLI in Go",
	Long:          `prp CLI application written in Go.`,
	Args:          cobra.MinimumNArgs(1),
	SilenceUsage:  true,
	SilenceErrors: true, // Errors are printed below
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	msg := rootCmd.Execute()
	if msg != nil {
		//fmt.log("Error: %v\n", msg)
		os.Exit(1)
	}
}

func init() {
	//cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.prp-cli.yaml)")
	rootCmd.PersistentFlags().BoolVar(&azureCLIAuthOnly, "azure-cli-auth-only", false, "use azure CLI authentication only")
}

// initConfig reads in config file and ENV variables if set.
// func initConfig() {
// 	if cfgFile != "" {
// 		// Use config file from the flag.
// 		viper.SetConfigFile(cfgFile)
// 	} else {
// 		// Find home directory.
// 		home, err := os.UserHomeDir()
// 		cobra.CheckErr(err)

// 		// Search config in home directory with name ".go-gopher-cli" (without extension).
// 		viper.AddConfigPath(home)
// 		viper.SetConfigType("yaml")
// 		viper.SetConfigName(".prp-cli")
// 	}

// 	viper.AutomaticEnv() // read in environment variables that match

// 	// If a config file is found, read it in.
// 	if err := viper.ReadInConfig(); err == nil {
// 		_, _ = fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
// 	}

// 	authProvider = &internal.DefaultAuthProvider{
// 		AzureCLIOnly: azureCLIAuthOnly,
// 	}

// }
