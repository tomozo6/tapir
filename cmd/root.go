package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tapir",
	Short: "Tapir can check the status of CloudWatchAlarm and can enable and disable alarm actions.",
	Long:  "Tapir can check the status of CloudWatchAlarm and can enable and disable alarm actions.",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.tapir.yaml)")
	rootCmd.PersistentFlags().StringP("prefix", "p", "", "Alam Name Preix")
	rootCmd.PersistentFlags().StringP("region", "r", "us-east-1", "Target AWS Region")
	rootCmd.PersistentFlags().StringToStringP("tags", "t", nil, "for filter. ex) project=test,env=dev")

	viper.BindPFlag("prefix", rootCmd.PersistentFlags().Lookup("prefix"))
	viper.BindPFlag("region", rootCmd.PersistentFlags().Lookup("region"))
	viper.BindPFlag("tags", rootCmd.PersistentFlags().Lookup("tags"))
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".tapir" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".tapir")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
